package training

import (
	"github.com/GrudTrigger/trainin_tracker/graph/model"
	"github.com/stretchr/testify/require"
	"testing"
)

type testData struct {
	title    string
	input    model.SearchTrainings
	wantSql  string
	wantArgs []interface{}
}

func TestQueryGetAll(t *testing.T) {
	typeTest := int32(1)
	test := []testData{
		{
			title:    "empty search params",
			input:    model.SearchTrainings{Limit: 20, Offset: 0},
			wantSql:  "SELECT * FROM training LIMIT $1 OFFSET $2",
			wantArgs: []interface{}{int32(20), int32(0)},
		},
		{
			title:    "search params with name",
			input:    model.SearchTrainings{Name: ptr("qwerty"), Limit: 20, Offset: 0},
			wantSql:  "SELECT * FROM training WHERE name ILIKE $1 LIMIT $2 OFFSET $3",
			wantArgs: []interface{}{"%qwerty%", int32(20), int32(0)},
		},
		{
			title:    "search params with type",
			input:    model.SearchTrainings{Type: ptr(int32(1)), Limit: 20, Offset: 0},
			wantSql:  "SELECT * FROM training WHERE type = $1 LIMIT $2 OFFSET $3",
			wantArgs: []interface{}{&typeTest, int32(20), int32(0)},
		}, {
			title:    "search params with name and type",
			input:    model.SearchTrainings{Name: ptr("qwerty"), Type: ptr(int32(1)), Limit: 20, Offset: 0},
			wantSql:  "SELECT * FROM training WHERE name ILIKE $1 AND type = $2 LIMIT $3 OFFSET $4",
			wantArgs: []interface{}{"%qwerty%", &typeTest, int32(20), int32(0)},
		},
	}
	for _, query := range test {
		t.Run(query.title, func(t *testing.T) {
			gotSql, gotArgs := QueryGetAll(query.input)
			require.Equal(t, query.wantSql, gotSql)
			require.Equal(t, query.wantArgs, gotArgs)
		})
	}
}

func ptr[T any](v T) *T {
	return &v
}
