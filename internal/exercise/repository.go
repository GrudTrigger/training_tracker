package exercise

import (
	"github.com/GrudTrigger/trainin_tracker/pkg/storage"
)

type IRepository interface {
	//Create(input *model.CreateExercise) (*model.Exercise, error)
	//GetAll(input *model.SearchExercise) ([]*model.Exercise, error)
	//GetById(id string) (*model.Exercise, error)
	//DeleteById(id string) (string, error)
}

type Repository struct {
	*storage.DbPostgres
}

func NewRepository(db *storage.DbPostgres) IRepository {
	return &Repository{db}
}

//func (r *Repository) Create(input *model.CreateExercise) (*model.Exercise, error) {
//	var e model.Exercise
//
//	query := `INSERT INTO exercise(title, muscle_group, approach_count, weight, training_id) VALUES ($1, $2, $3, $4, $5)
//				RETURNING id, title, muscle_group, approach_count, weight, created_at`
//	err := r.QueryRow(query, input.Title, input.MuscleGroup, input.ApproachCount, input.Weight, input.TrainingID).
//		Scan(e.ID, e.Title, e.MuscleGroup, e.ApproachCount, e.Weight, e.CreatedAt)
//	if err != nil {
//		return nil, err
//	}
//	return &e, nil
//}
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
