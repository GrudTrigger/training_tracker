package exercise

import (
	"github.com/GrudTrigger/trainin_tracker/graph/model"
	"github.com/GrudTrigger/trainin_tracker/internal/approach"
	"github.com/GrudTrigger/trainin_tracker/pkg/storage"
)

type IExerciseRepository interface {
	Create(input []*model.CreateExercise, trainingId string) ([]*model.Exercise, error)
	//GetAll(input *model.SearchExercise) ([]*model.Exercise, error)
	//GetById(id string) (*model.Exercise, error)
	//DeleteById(id string) (string, error)
}

type Repository struct {
	*storage.DbPostgres
	repoApproach approach.IApproachRepository
}

func NewRepository(db *storage.DbPostgres, repoApproach approach.IApproachRepository) IExerciseRepository {
	return &Repository{db, repoApproach}
}

func (r *Repository) Create(input []*model.CreateExercise, trainingId string) ([]*model.Exercise, error) {
	var list []*model.Exercise
	for _, v := range input {
		var e model.Exercise
		e = model.Exercise{
			Approaches:   []*model.Approach{},
			ExerciseList: &model.ExerciseList{},
		}
		query := "INSERT INTO exercise(training_id, exercise_list_id) VALUES ($1, $2) RETURNING id, training_id"
		err := r.QueryRow(query, trainingId, v.ExerciseListID).Scan(&e.ID, &e.TrainingID)
		if err != nil {
			return nil, err
		}

		querySearchExerciseList := "SELECT * FROM exercise_list WHERE id = $1"

		var exList model.ExerciseList
		err = r.QueryRow(querySearchExerciseList, v.ExerciseListID).Scan(&exList.ID, &exList.Title, &exList.CategoryMuscle, &exList.CreatedAt)
		if err != nil {
			return nil, err
		}

		e.ExerciseList = &exList

		approaches, err := r.repoApproach.Create(v.Approaches, e.ID)
		e.Approaches = approaches
		list = append(list, &e)
	}
	return list, nil
}

//
//func (r *Repository) GetAll(input *model.SearchExercise) ([]*model.Exercise, error) {
//	var exercise []*model.Exercise
//	query, args := QueryGetAll(input)
//	rows, err := r.Query(query, args...)
//	if errors.Is(err, sql.ErrNoRows) {
//		return nil, errors.New("подход не найден")
//	}
//	if err != nil {
//		return nil, err
//	}
//	for rows.Next() {
//		var ex model.Exercise
//		err = rows.Scan(&ex.ID, &ex.TrainingID, &ex.Title, &ex.MuscleGroup, &ex.ApproachCount, &ex.Weight, &ex.CreatedAt)
//		if err != nil {
//			return nil, err
//		}
//		exercise = append(exercise, &ex)
//	}
//	return exercise, nil
//}
//
//func (r *Repository) GetById(id string) (*model.Exercise, error) {
//	var e model.Exercise
//	query := "SELECT * from exercise WHERE id = $1"
//	err := r.QueryRow(query, id).Scan(&e.ID, &e.TrainingID, &e.Title, &e.MuscleGroup, &e.ApproachCount, &e.Weight, &e.CreatedAt)
//	if err != nil {
//		return nil, err
//	}
//	return &e, nil
//}
//
//func(r *Repository) DeleteById(id string) (string,error) {
//	query := "DELETE from exercise WHERE id = $1"
//	_, err := r.Exec(query, id)
//	if err != nil {
//		return "", res.ErrAccessDenied
//	}
//	return res.SuccessResponse, nil
//}
