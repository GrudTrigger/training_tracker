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
	queryBuilder.WriteString("SELECT * FROM training")

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
