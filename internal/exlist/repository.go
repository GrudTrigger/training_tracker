package exlist

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/GrudTrigger/trainin_tracker/graph/model"
	"github.com/GrudTrigger/trainin_tracker/pkg/res"
	"github.com/GrudTrigger/trainin_tracker/pkg/storage"
)

type IRepository interface {
	Create(input *model.CreateExerciseForList) (*model.ExerciseList, error)
	FindAll(input *model.GetExerciseList) ([]*model.ExerciseList, error)
	Update(input *model.UpdateExerciseForList) (*model.ExerciseList, error)
	DeleteById(id string) (string, error)
	Statistics() (*model.ExerciseListStatistic, error)
}

type Repository struct {
	*storage.DbPostgres
}

func NewRepository(db *storage.DbPostgres) IRepository {
	return &Repository{db}
}

func (r *Repository) Create(input *model.CreateExerciseForList) (*model.ExerciseList, error) {
	var e model.ExerciseList
	query := "INSERT INTO exercise_list(title, category_muscle)  VALUES($1, $2) RETURNING id, title, category_muscle, created_at"
	err := r.QueryRow(query, input.Title, input.CategoryMuscle).Scan(e.ID, e.Title, e.CategoryMuscle, e.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &e, nil
}
func (r *Repository) FindAll(input *model.GetExerciseList) ([]*model.ExerciseList, error) {
	var eList []*model.ExerciseList

	query, args := getQueryFindAll(input)
	rows, err := r.Query(query, args...)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, errors.New("упражнения не найдены")
	}

	for rows.Next() {
		var e model.ExerciseList
		err := rows.Scan(&e.ID, &e.Title, &e.CategoryMuscle, &e.CreatedAt)
		if err != nil {
			return nil, err
		}
		eList = append(eList, &e)
	}
	return eList, nil
}

func (r *Repository) Update(input *model.UpdateExerciseForList) (*model.ExerciseList, error) {
	var existedExercise model.ExerciseList
	searchQuery := "SELECT id, title, category_muscle, created_at FROM exercise_list WHERE id = $1"
	err := r.QueryRow(searchQuery, input.ID).Scan(&existedExercise.ID, &existedExercise.Title, &existedExercise.CategoryMuscle, &existedExercise.CreatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errors.New("упражнение не найдено, id =" + input.ID)
	}

	if existedExercise.Title != input.Title {
		existedExercise.Title = input.Title
	}

	if existedExercise.CategoryMuscle != input.CategoryMuscle {
		existedExercise.CategoryMuscle = input.CategoryMuscle
	}
	updateQuery := "UPDATE exercise_list SET title = $1, category_muscle = $2 WHERE id = $3"
	_, err = r.Exec(updateQuery, existedExercise.Title, existedExercise.CategoryMuscle, existedExercise.ID)
	if err != nil {
		return nil, err
	}
	return &existedExercise, nil
}

func (r *Repository) DeleteById(id string) (string, error) {
	query := "DELETE  from exercise_list WHERE id = $1"
	_, err := r.Exec(query, id)
	if err != nil {
		return "", err
	}
	return res.SuccessResponse, nil
}

func (r *Repository) Statistics() (*model.ExerciseListStatistic, error) {
	var result model.ExerciseListStatistic
	//TODO: попробывать сделать под один буфферизированный канал, с label:string, value: {}interface, чтобы не плодить каналы под каждую горутину, а читать из одного канала все данные и объеденить
	chForAll := make(chan int32)
	chForCat := make(chan int32)
	chForStat := make(chan []*model.MuscleGroupCount)
	chForErr := make(chan error, 3)
	ctx, cancel := context.WithCancel(context.Background())
	// AllExercise
	go func() {
		var countAllExercise int32
		query := "SELECT COUNT(id) FROM exercise_list"
		err := r.QueryRow(query).Scan(&countAllExercise)
		if err != nil {
			chForErr <- fmt.Errorf("all exercise: %w", err)
			cancel()
		}
		chForAll <- countAllExercise
	}()

	// AllCategory
	go func() {
		var countCategory int32
		query := "SELECT COUNT(DISTINCT category_muscle) FROM exercise_list"
		err := r.QueryRow(query).Scan(&countCategory)
		if err != nil {
			chForErr <- fmt.Errorf("all category: %w", err)
			cancel()
		}
		chForCat <- countCategory
	}()

	// StatisticCategory
	go func() {
		// StatisticCategory []*MuscleGroupCount
		var mgc []*model.MuscleGroupCount
		query := "SELECT category_muscle, COUNT(*) as count FROM exercise_list GROUP BY category_muscle"
		rows, err := r.Query(query)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				chForErr <- fmt.Errorf("statistic category query: %w", err)
				cancel()
			}
		}
		for rows.Next() {
			var countCat model.MuscleGroupCount
			err := rows.Scan(&countCat.CategoryMuscle, &countCat.Count)
			if err != nil {
				chForErr <- fmt.Errorf("statistic category scan: %w", err)
				cancel()
			}
			mgc = append(mgc, &countCat)
		}
		chForStat <- mgc
	}()

	if len(chForErr) > 0 {
		return nil, <-chForErr
	}
	count := 0
	for {
		select {
		case v := <-chForAll:
			count++
			result.AllExercise = v
			if count == 2 {
				return &result, nil
			}
		case v := <- chForCat:
			count++
			result.AllCategory = v
			if count == 2 {
				return &result, nil
			}
		case v := <- chForStat:
			count++
			result.StatisticCategory = v
			if count == 2 {
				return &result, nil
			}
		case <- ctx.Done():
			return nil, <-chForErr
		}
	}
	// result.AllExercise = <-chForAll
	// result.AllCategory = <-chForCat
	// result.StatisticCategory = <-chForStat
}
