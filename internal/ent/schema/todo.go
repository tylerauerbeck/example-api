package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"go.infratographer.com/x/entx"
	"go.infratographer.com/x/gidx"
)

// Todo holds the schema definition for the Todo entity.
type Todo struct {
	ent.Schema
}

// Mixin of the Todo type
func (Todo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		entx.NewTimestampMixin(),
	}
}

// Fields of the Todo.
func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(gidx.PrefixedID("")).
			Unique().
			Immutable().
			Comment("The ID of the Todo.").
			Annotations(
				entgql.OrderField("ID"),
			).
			DefaultFunc(func() gidx.PrefixedID { return gidx.MustNewID(TodoPrefix) }),
		field.Text("name").
			NotEmpty().
			Comment("The name of the todo.").
			Annotations(
				entgql.OrderField("NAME"),
			),
		field.Text("description").
			Optional().
			Comment("The description of the todo.").
			Annotations(
				entgql.OrderField("DESCRIPTION"),
			),
		field.String("tenant_id").
			GoType(gidx.PrefixedID("")).
			Immutable().
			Comment("The ID for the tenant for this todo.").
			Annotations(
				entgql.QueryField(),
				entgql.Type("ID"),
				entgql.Skip(entgql.SkipWhereInput, entgql.SkipMutationUpdateInput, entgql.SkipType),
				entgql.OrderField("TENANT"),
			),
	}
}

// Edges of the Todo.
func (Todo) Edges() []ent.Edge {
	return nil
}

// Indexes of the Todo.
func (Todo) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("tenant_id"),
	}
}

// Annotations of the Todo
func (Todo) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entx.GraphKeyDirective("id"),
		schema.Comment("Represents an todo todo on the graph."),
		entgql.RelayConnection(),
		entgql.Mutations(
			entgql.MutationCreate().Description("Create a new todo todo."),
			entgql.MutationUpdate().Description("Update an existing todo todo."),
		),
	}
}
