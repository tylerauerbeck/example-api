// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphapi

import (
	"go.infratographer.com/example-api/internal/ent/generated"
	"go.infratographer.com/x/gidx"
)

// Return response for createExample mutation
type ExampleCreatePayload struct {
	// Created example
	Example *generated.Example `json:"example"`
}

// Return response for deleteExample mutation
type ExampleDeletePayload struct {
	// Deleted example
	DeletedID gidx.PrefixedID `json:"deletedID"`
}

// Return response for updateExample mutation
type ExampleUpdatePayload struct {
	// Updated example
	Example *generated.Example `json:"example"`
}

type Tenant struct {
	ID       gidx.PrefixedID              `json:"id"`
	Examples *generated.ExampleConnection `json:"examples"`
}

func (Tenant) IsEntity() {}
