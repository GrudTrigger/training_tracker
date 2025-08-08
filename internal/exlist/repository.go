package exlist

import (
	"database/sql"
	"errors"

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
	err := r.QueryRow(query, input.Title, input.CategoryMuscle).Scan(&e.ID, &e.Title, &e.CategoryMuscle, &e.CreatedAt)
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
	result := &model.ExerciseListStatistic{}
	countCh := 3
	chResult := make(chan statResult, countCh)

	// AllExercise
	go func() {
		var count int32
		err := r.QueryRow("SELECT COUNT(id) FROM exercise_list").Scan(&count)
		chResult <- statResult{label: "AllExercise", value: count, err: err}
	}()

	// AllCategory
	go func() {
		var count int32
		err := r.QueryRow("SELECT COUNT(DISTINCT category_muscle) FROM exercise_list").Scan(&count)
		chResult <- statResult{label: "AllCategory", value: count, err: err}
	}()

	// StatisticCategory
	go func() {
		// StatisticCategory []*MuscleGroupCount
		rows, err := r.Query("SELECT category_muscle, COUNT(*) as count FROM exercise_list GROUP BY category_muscle")
		if err != nil {
			chResult <- statResult{label: "StatisticCategory", err: err}
			return
		}
		defer rows.Close()

		var list []*model.MuscleGroupCount
		for rows.Next() {
			var c model.MuscleGroupCount
			err := rows.Scan(&c.CategoryMuscle, &c.Count)
			if err != nil {
				chResult <- statResult{label: "StatisticCategory", err: err}
				return
			}
			list = append(list, &c)
		}
		chResult <- statResult{label: "StatisticCategory", value: list}
	}()

	for i := 0; i < countCh; i++ {
		select {
		case r := <-chResult:
			if r.err != nil {
				return nil, r.err
			}
			switch r.label {
			case "AllExercise":
				result.AllExercise = r.value.(int32)
			case "AllCategory":
				result.AllCategory = r.value.(int32)
			case "StatisticCategory":
				result.StatisticCategory = r.value.([]*model.MuscleGroupCount)
			}
		}
	}
	return result, nil
}
