package exercise

import (
	"fmt"
	"github.com/GrudTrigger/trainin_tracker/graph/model"
	"strings"
)

func QueryGetAll(input *model.SearchExercise) (string, []interface{}) {
	var (
		queryBuilder strings.Builder
		args         []interface{}
		conditions   []string
	)
	queryBuilder.WriteString("SELECT * from exercise")

	if input.Title != nil {
		args = append(args, "%"+*input.Title+"%")
		conditions = append(conditions, fmt.Sprintf("title ILIKE $%d", len(args)))
	}

	if input.MuscleGroup != nil {
		args = append(args, *input.MuscleGroup)
		conditions = append(conditions, fmt.Sprintf("muscle_group = $%d", len(args)))
	}

	if input.ApproachCount != nil {
		args = append(args, *input.ApproachCount)
		conditions = append(conditions, fmt.Sprintf("approach_count = $%d", len(args)))
	}

	if input.Weight != nil {
		args = append(args, *input.Weight)
		conditions = append(conditions, fmt.Sprintf("weight = $%d", len(args)))
	}

	if len(conditions) > 0 {
		queryBuilder.WriteString(" WHERE ")
		queryBuilder.WriteString(strings.Join(conditions, " AND "))
	}

	args = append(args, input.Limit, input.Offset)
	queryBuilder.WriteString(fmt.Sprintf("LIMIT $%d OFFSET $%d", len(args)-1, len(args)))
	return queryBuilder.String(), args
}
