package approach

import (
	"github.com/GrudTrigger/trainin_tracker/graph/model"
	"github.com/GrudTrigger/trainin_tracker/pkg/storage"
)

type IApproachRepository interface {
	Create(input []*model.CreateApproach, exerciseId string) ([]*model.Approach, error)
}

type Repository struct {
	*storage.DbPostgres
}

func NewRepository(db *storage.DbPostgres) IApproachRepository {
	return &Repository{db}
}

func (r *Repository) Create(input []*model.CreateApproach, exerciseId string) ([]*model.Approach, error) {
	var list []*model.Approach
	//TODO:посмотреть, как будет работать и переписать на горутины, создавать в каждой горутине отдельно
	for _, v := range input {
		var a model.Approach
		query := "INSERT INTO approach(exercise_id, repetition, weight) VALUES($1, $2, $3) RETURNING id, exercise_id, repetition, weight"
		err := r.QueryRow(query, exerciseId, v.Repetition, v.Weight).Scan(&a.ID, &a.ExerciseID, &a.Repetition, &a.Weight)
		if err != nil {
			return nil, err
		}
		list = append(list, &a)
	}
	return list, nil
}
