package api_test

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.infratographer.com/x/gidx"

	"go.infratographer.com/example-api/internal/ent/generated"
	"go.infratographer.com/example-api/internal/testclient"
)

func TestQuery_todo(t *testing.T) {
	client := graphTestClient()
	ctx := context.Background()
	tenantID := gidx.MustNewID("testtnt")
	td1 := (&TodoBuilder{
		TenantID: tenantID,
	}).MustNew(ctx)
	td2 := (&TodoBuilder{
		Description: gofakeit.HackerPhrase(),
		TenantID:    tenantID,
	}).MustNew(ctx)

	testCases := []struct {
		name           string
		queryID        gidx.PrefixedID
		hasDescription bool
		expextedTodo   *generated.Todo
		errorMsg       string
	}{
		{
			name:         "happy path td1",
			queryID:      td1.ID,
			expextedTodo: td1,
		},
		{
			name:           "happy path td2",
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
		t.Run("Create "+tc.name, func(t *testing.T) {
			resp, err := client.GetTodo(ctx, tc.queryID)

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

func Test_HappyPath(t *testing.T) {
	client := graphTestClient()
	ctx := context.Background()
	tenantID := gidx.MustNewID("testtnt")

	t.Run("Create + List + Update + Delete", func(t *testing.T) {
		td, err := client.TodoCreate(ctx, testclient.CreateTodoInput{
			Name:        gofakeit.JobTitle(),
			Description: nil,
			TenantID:    tenantID,
		})
		require.NoError(t, err)
		require.NotNil(t, td)

		list, err := client.ListTodos(ctx, tenantID, nil)
		require.NoError(t, err)
		require.NotNil(t, list)
		assert.Len(t, list.Entities[0].Todo.Edges, 1)

		tdUpdate, err := client.TodoUpdate(ctx, td.TodoCreate.Todo.ID, testclient.UpdateTodoInput{
			Name:        newString(gofakeit.JobTitle()),
			Description: newString(gofakeit.HackerPhrase()),
		})

		require.NoError(t, err)
		require.NotNil(t, tdUpdate)

		assert.NotEqual(t, td.TodoCreate.Todo.Name, tdUpdate.TodoUpdate.Todo.Name)
		assert.NotEqual(t, td.TodoCreate.Todo.Description, tdUpdate.TodoUpdate.Todo.Description)

		deleteID, err := client.TodoDelete(ctx, td.TodoCreate.Todo.ID)
		require.NoError(t, err)
		require.NotNil(t, deleteID)
		assert.EqualValues(t, td.TodoCreate.Todo.ID, deleteID.TodoDelete.DeletedID)
	})
}
