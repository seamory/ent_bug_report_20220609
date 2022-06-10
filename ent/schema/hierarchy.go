package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Hierarchy holds the schema definition for the Hierarchy entity.
type Hierarchy struct {
	ent.Schema
}

func (Hierarchy) Mixin() []ent.Mixin {
	return []ent.Mixin{}
}

// Fields of the Hierarchy.
func (Hierarchy) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique().Comment("hierarchy name"),
		field.String("describe").Comment("hierarchy describe"),
		field.Int("sort").Default(0).Comment("hierarchy sort"),
	}
}

// Edges of the Hierarchy.
func (Hierarchy) Edges() []ent.Edge {
	return []ent.Edge{
		// todo Uncomment following, and run generate, will throw error will multiple storage-keys.
		//edge.To("child", Hierarchy.Type).
		//	StorageKey(edge.Table("hierarchy_relation"), edge.Columns("parent_id", "child_id")).
		//	From("parent"),
		// todo Uncomment following, and run TestBugSQLite in bug_test.go.
		edge.To("child", Hierarchy.Type).StorageKey(edge.Table("hierarchy_relation"), edge.Columns("parent_id", "child_id")),
		edge.From("parent", Hierarchy.Type).Ref("child"),
	}
}
