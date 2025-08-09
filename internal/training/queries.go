package training

import (
	"fmt"
	"strings"

	"github.com/GrudTrigger/trainin_tracker/graph/model"
)

func QueryGetAll(input model.SearchTrainings) (string, []interface{}) {
	var (
		queryBuilder strings.Builder
		args         []interface{}
		conditions   []string
	)
	queryBuilder.WriteString("SELECT training.*, exercise.id, exercise.training_id, exercise_list.id, exercise_list.title, exercise_list.category_muscle, exercise_list.created_at, approach.id, approach.exercise_id, approach.repetition, approach.weight FROM training JOIN exercise ON training.id = exercise.training_id JOIN exercise_list ON exercise.exercise_list_id = exercise_list.id JOIN approach ON exercise.id = approach.exercise_id")

	if input.Name != nil {
		args = append(args, "%"+*input.Name+"%")
		conditions = append(conditions, fmt.Sprintf("name ILIKE $%d", len(args)))
	}

	if input.Type != nil {
		args = append(args, input.Type)
		conditions = append(conditions, fmt.Sprintf("type = $%d", len(args)))
	}

	if len(conditions) > 0 {
		queryBuilder.WriteString(" WHERE ")
		queryBuilder.WriteString(strings.Join(conditions, " AND "))
	}
	args = append(args, input.Limit, input.Offset)
	queryBuilder.WriteString(fmt.Sprintf(" LIMIT $%d OFFSET $%d", len(args)-1, len(args)))
	return queryBuilder.String(), args
}
