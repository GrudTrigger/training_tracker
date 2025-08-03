package exlist

import (
	"fmt"
	"github.com/GrudTrigger/trainin_tracker/graph/model"
	"strings"
)

func getQueryFindAll(input *model.GetExerciseList) (string, []interface{}) {
	var (
		queryBuilder strings.Builder
		args         []interface{}
		conditions   []string
	)

	queryBuilder.WriteString("SELECT * FROM exercise_list")

	if input.Title != nil {
		args = append(args, "%"+*input.Title+"%")
		conditions = append(conditions, fmt.Sprintf("title ILIKE %d", len(args)))
	}

	if input.CategoryMuscle != nil {
		args = append(args, input.CategoryMuscle)
		conditions = append(conditions, fmt.Sprintf("category_muscle = %d", len(args)))
	}

	if len(conditions) > 0 {
		queryBuilder.WriteString(" WHERE ")
		queryBuilder.WriteString(strings.Join(conditions, " AND "))
	}
	args = append(args, input.Limit, input.Offset)
	queryBuilder.WriteString(fmt.Sprintf(" ORDER BY category_muscle LIMIT %d OFFSET %d", len(args)-1, len(args)))
	return queryBuilder.String(), args
}
