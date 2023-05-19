package api_test

import (
	"context"

	"github.com/brianvoe/gofakeit/v6"
	"go.infratographer.com/x/gidx"

	ent "go.infratographer.com/example-api/internal/ent/generated"
)

type TodoBuilder struct {
	Name        string
	Description string
	TenantID    gidx.PrefixedID
}

func (t *TodoBuilder) MustNew(ctx context.Context) *ent.Todo {
	todoCreate := EntClient.Todo.Create()

	if t.Name == "" {
		t.Name = gofakeit.JobTitle()
	}

	todoCreate.SetName(t.Name)

	if t.Description != "" {
		todoCreate.SetDescription(t.Description)
	}

	if t.TenantID == "" {
		t.TenantID = gidx.MustNewID(tenantPrefix)
	}

	todoCreate.SetTenantID(t.TenantID)

	return todoCreate.SaveX(ctx)
}
