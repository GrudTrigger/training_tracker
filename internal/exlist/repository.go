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
	DeleteById(id string) (string, error)
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

	query := "SELECT * from exercise_list ORDER BY category_muscle LIMIT $1 OFFSET $2"
	rows, err := r.Query(query, input.Limit, input.Offset)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, errors.New("упражнения не найдены")
	}

	if rows.Next() {
		var e model.ExerciseList
		err := rows.Scan(&e.ID, &e.Title, &e.CategoryMuscle, &e.CreatedAt)
		if err != nil {
			return nil, err
		}
		eList = append(eList, &e)
	}
	return eList, nil
}

func (r *Repository) DeleteById(id string) (string, error) {
	query := "DELETE  from exercise_list WHERE id = $1"
	_, err := r.Exec(query, id)
	if err != nil {
		return "", err
	}
	return res.SuccessResponse, nil
}
