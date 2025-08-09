package training

import (
	"database/sql"
	"errors"
	"sync"

	"github.com/GrudTrigger/trainin_tracker/graph/model"
	"github.com/GrudTrigger/trainin_tracker/internal/exercise"
	"github.com/GrudTrigger/trainin_tracker/pkg/res"
	"github.com/GrudTrigger/trainin_tracker/pkg/storage"
	"github.com/GrudTrigger/trainin_tracker/pkg/utils"
)

type IRepository interface {
	Create(input InputWithUser) (*model.Training, error)
	GetAll(input model.SearchTrainings) ([]*model.Training, error)
	GetById(id string) (*model.Training, error)
	GetMyTrainings(userId string) ([]*model.Training, error)
	DeleteById(id string) (string, error)
}

type Repository struct {
	*storage.DbPostgres
	repoExercise exercise.IExerciseRepository
}

func NewTrainingRepository(db *storage.DbPostgres, repoExercise exercise.IExerciseRepository) IRepository {
	return &Repository{db, repoExercise}
}

func (r *Repository) Create(input InputWithUser) (*model.Training, error) {
	var training model.Training
	training = model.Training{
		Exercises: []*model.Exercise{},
	}
	query := `INSERT INTO training(user_id, name, duration, date, notes) 
									VALUES($1, $2, $3, $4, $5)
									RETURNING id, user_id, name,duration, date, notes, created_at`
	err := r.QueryRow(query, input.UserId, input.Title, input.Duration, input.Date, input.Notes).
		Scan(&training.ID, &training.UserID, &training.Title, &training.Duration, &training.Date, &training.Notes, &training.CreatedAt)
	if err != nil {
		return nil, err
	}

	e, err := r.repoExercise.Create(input.Exercises, training.ID)
	if err != nil {
		return nil, err
	}
	training.Exercises = e
	return &training, nil
}

func (r *Repository) GetAll(input model.SearchTrainings) ([]*model.Training, error) {

	query, args := QueryGetAll(input)
	rows, err := r.Query(query, args...)
	if err != nil {
		return nil, err
	}

	trainingsMap := make(map[string]*model.Training)
	for rows.Next() {
		t := utils.NewTraining()
		e := model.Exercise{}
		el := model.ExerciseList{}
		ap := model.Approach{}

		err = rows.Scan(
			&t.ID,
			&t.UserID,
			&t.Title,
			&t.Duration,
			&t.Date,
			&t.Notes,
			&t.CreatedAt,
			&e.ID,
			&e.TrainingID,
			&el.ID,
			&el.Title,
			&el.CategoryMuscle,
			&el.CreatedAt,
			&ap.ID,
			&ap.ExerciseID,
			&ap.Repetition,
			&ap.Weight,
		)
		if err != nil {
			return nil, err
		}

		// 1. Получаем или создаём тренировку
		tr, ok := trainingsMap[t.ID]
		if !ok {
			ch := make(chan ChanCounter, 3)
			wg := sync.WaitGroup{}

			tr = t
			tr.Exercises = []*model.Exercise{}
			trainingsMap[t.ID] = tr

			wg.Add(1)
			go func() {
				defer wg.Done()
				var appCount int32
				var weightCount int32
				var exerciseCount int32

				q := "SELECT COUNT(approach.id) AS count_approach, SUM(approach.weight) AS sum_weight, COUNT(DISTINCT exercise.id) AS exercises_count  FROM training JOIN exercise ON exercise.training_id = training.id JOIN approach ON approach.exercise_id = exercise.id GROUP BY training.id"
				err = r.QueryRow(q).Scan(&appCount, &weightCount, &exerciseCount)
				ch <- ChanCounter{label: "approach_count", value: appCount, err: err}
				ch <- ChanCounter{label: "total_weight", value: weightCount, err: err}
				ch <- ChanCounter{label: "exercise_count", value: exerciseCount, err: err}
			}()
			go func() {
				wg.Wait()
				close(ch)
			}()

			for v := range ch {
				if v.err != nil {
					return nil, v.err
				}
				if v.label == "approach_count" {
					tr.ApproachCount = v.value.(int32)
				}
				if v.label == "total_weight" {
					tr.TotalWeight = v.value.(int32)
				}
				if v.label == "exercise_count" {
					tr.ExercisesCount = v.value.(int32)
				}
			}
		}

		// 2. Если есть упражнение
		var ex *model.Exercise
		if e.ID != "" {
			// ищем в существующих упражнениях
			found := false
			for _, existingEx := range tr.Exercises {
				if existingEx.ID == e.ID {
					ex = existingEx
					found = true
					break
				}
			}

			// если нет — добавляем
			if !found {
				e.ExerciseList = &el
				e.Approaches = []*model.Approach{}
				tr.Exercises = append(tr.Exercises, &e)
				ex = &e
			}
		}

		// 3. Если есть подход
		if ap.ID != "" && ex != nil {
			found := false
			for _, existingAp := range ex.Approaches {
				if existingAp.ID == ap.ID {
					found = true
					break
				}
			}
			if !found {
				ex.Approaches = append(ex.Approaches, &ap)
			}
		}
	}

	// Конвертируем map в slice
	trainings := make([]*model.Training, 0, len(trainingsMap))
	for _, tr := range trainingsMap {
		trainings = append(trainings, tr)
	}
	return trainings, nil
}

func (r *Repository) GetById(id string) (*model.Training, error) {
	training := utils.NewTraining()

	query := `SELECT training.*, users.id, users.email, users.login, users.password, users.role, users.created_at, exercise.* from training JOIN users ON training.user_id = users.id JOIN exercise ON training.id = exercise.training_id WHERE training.id = $1`
	row := r.QueryRow(query, id)
	if errors.Is(row.Err(), sql.ErrNoRows) {
		return nil, errors.New("тренировка не найдена")
	}

	err := row.Scan(
		&training.ID,
		&training.UserID,
		&training.Title,
		&training.Duration,
		&training.Date,
		&training.Notes,
		&training.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return training, nil
}

func (r *Repository) GetMyTrainings(userId string) ([]*model.Training, error) {

	var trainings []*model.Training

	query := "SELECT * FROM training WHERE user_id = $1"

	rows, err := r.Query(query, userId)

	if !errors.Is(err, sql.ErrNoRows) {
		return nil, errors.New("тренировки не найдены")
	}

	for rows.Next() {
		var training model.Training
		err = rows.Scan(&training.ID, &training.UserID, &training.Title, &training.Duration, &training.Date, &training.Notes, &training.CreatedAt)
		if err != nil {
			return nil, err
		}
		trainings = append(trainings, &training)
	}

	return trainings, nil
}

func (r *Repository) DeleteById(id string) (string, error) {
	query := "DELETE FROM training WHERE id = $1"
	_, err := r.Exec(query, id)
	if err != nil {
		return "", err
	}
	return res.SuccessResponse, nil
}
