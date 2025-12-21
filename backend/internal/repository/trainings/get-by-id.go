package trainings

import (
	"context"
	"time"

	t "github.com/GrudTrigger/training_tracker/backend/gen/trainings"
)

func (r *Repository) GetByID(ctx context.Context, id *t.GetByIDPayload) (*t.TrainingAll, error) {
	var res t.TrainingAll

	row := r.conn.QueryRow(ctx, "SELECT trainings.id, trainings.title, trainings.date, trainings.duration, trainings.created_at, exercises.id,exercises.title, exercises.muscle_group, exercise_sets.id, exercise_sets.reps, exercise_sets.weight FROM trainings JOIN training_exercises te on trainings.id = te.training_id JOIN exercises on te.exercise_id = exercises.id JOIN exercise_sets on te.id = exercise_sets.training_exercise_id WHERE trainings.id=$1", id)

	var (
		trainingID         string
		trainingTitle      string
		trainingDate       time.Time
		trainingDuration   int
		trainingCreatedAt  *time.Time
		exerciseID         string
		exerciseTitle      string
		muscleGroup        int32
		trainingExerciseID string
		setID              string
		reps               int
		weight             float64
	)

	err := row.Scan(
		&trainingID,
		&trainingTitle,
		&trainingDate,
		&trainingDuration,
		&trainingCreatedAt,
		&exerciseID,
		&exerciseTitle,
		&muscleGroup,
		&trainingExerciseID,
		&setID,
		&reps,
		&weight,
	)
	if err != nil {
		return nil, err
	}

	if res.ID == "" {
		c := trainingCreatedAt.Format(time.DateTime)
		res = t.TrainingAll{
			ID:        trainingID,
			Title:     trainingTitle,
			Date:      trainingDate.Format(time.DateOnly),
			Duration:  trainingDuration,
			CreatedAt: &c,
		}
	}
	var ex *t.ExercisesWithTraining

	for _, e := range res.Exercises {
		if e.ID == exerciseID {
			ex = e
			break
		}
	}

	if ex == nil {
		ex = &t.ExercisesWithTraining{
			ID:          exerciseID,
			Title:       exerciseTitle,
			MuscleGroup: muscleGroup,
			Sets:        make([]*t.ExerciseSet, 0),
		}
		res.Exercises = append(res.Exercises, ex)
	}
	ex.Sets = append(ex.Sets, &t.ExerciseSet{
		ID:     &setID,
		Reps:   reps,
		Weight: &weight,
	})

	return &res, nil
}
