package api_test

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.infratographer.com/x/gidx"

	"go.infratographer.com/example-api/internal/ent/generated"
)

func TestQuery_todo(t *testing.T) {
	ctx := context.Background()
	td1 := (&TodoBuilder{}).MustNew(ctx)
	td2 := (&TodoBuilder{
		Description: gofakeit.HackerPhrase(),
	}).MustNew(ctx)

	testCases := []struct {
		name           string
		queryID        gidx.PrefixedID
		hasDescription bool
		expextedTodo   *generated.Todo
		errorMsg       string
	}{
		{
			name:         "happy path - td1",
			queryID:      td1.ID,
			expextedTodo: td1,
		},
		{
			name:           "happy path - td2",
			queryID:        td2.ID,
			hasDescription: true,
			expextedTodo:   td2,
		},
		{
			name:     "invalid-id",
			queryID:  gidx.MustNewID("testing"),
			errorMsg: "todo not found",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := graphTestClient().GetTodo(ctx, tc.queryID)

			if tc.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.Todo)
			assert.EqualValues(t, tc.expextedTodo.ID, resp.Todo.ID)
			if tc.hasDescription {
				assert.Equal(t, tc.expextedTodo.Description, *resp.Todo.Description)
			}
		})
	}
}
