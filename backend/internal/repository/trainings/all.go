package trainings

import (
	"context"
	"time"

	t "github.com/GrudTrigger/training_tracker/backend/gen/trainings"
)

// TODO: при интеграции initData и добавления таблицы users, получать только тренировки которые создал пользователь, чтобы не видить чужие тренировки
func (r *Repository) All(ctx context.Context, data *t.AllPayload) ([]*t.TrainingAll, error) {
	rows, err := r.conn.Query(ctx, `
			SELECT
			t.id                AS training_id,
			t.title             AS training_title,
			t.date              AS training_date,
			t.duration          AS training_duration,
			t.note              AS training_note,
			t.created_at        AS training_created_at,

			e.id                AS exercise_id,
			e.title             AS exercise_title,
			e.muscle_group,

			te.id               AS training_exercise_id,

			es.id               AS set_id,
			es.reps,
			es.weight
			FROM trainings t
			JOIN training_exercises te ON te.training_id = t.id
			JOIN exercises e           ON e.id = te.exercise_id
			JOIN exercise_sets es ON es.training_exercise_id = te.id
			ORDER BY t.date DESC, e.title LIMIT $1 OFFSET $2;
	`, data.Limit, data.Offset)
	if err != nil {
		return nil, err
	}

	trainings := map[string]*t.TrainingAll{}
	for rows.Next() {
		var (
			trainingID         string
			trainingTitle      string
			trainingDate       time.Time
			trainingDuration   int
			trainingNote       *string
			trainingCreatedAt  *time.Time
			exerciseID         string
			exerciseTitle      string
			muscleGroup        int32
			trainingExerciseID string
			setID              string
			reps               int
			weight             float64
		)

		err := rows.Scan(
			&trainingID,
			&trainingTitle,
			&trainingDate,
			&trainingDuration,
			&trainingNote,
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
		tr, ok := trainings[trainingID]
		if !ok {
			c := trainingCreatedAt.Format(time.DateTime)
			tr = &t.TrainingAll{
				ID:        trainingID,
				Title:     trainingTitle,
				Date:      trainingDate.Format(time.DateOnly),
				Duration:  trainingDuration,
				Note:      trainingNote,
				CreatedAt: &c,
				Exercises: make([]*t.ExercisesWithTraining, 0),
			}
			trainings[trainingID] = tr
		}
		var ex *t.ExercisesWithTraining
		for _, e := range tr.Exercises {
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
			tr.Exercises = append(tr.Exercises, ex)
		}
		ex.Sets = append(ex.Sets, &t.ExerciseSet{
			ID:     &setID,
			Reps:   reps,
			Weight: &weight,
		})
		if err := rows.Err(); err != nil {
			return nil, err
		}
	}
	res := make([]*t.TrainingAll, 0, len(trainings))
	for _, tr := range trainings {
		res = append(res, tr)
	}
	return res, nil
}
