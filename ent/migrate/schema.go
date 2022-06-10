// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// HierarchiesColumns holds the columns for the "hierarchies" table.
	HierarchiesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "describe", Type: field.TypeString},
		{Name: "sort", Type: field.TypeInt, Default: 0},
	}
	// HierarchiesTable holds the schema information for the "hierarchies" table.
	HierarchiesTable = &schema.Table{
		Name:       "hierarchies",
		Columns:    HierarchiesColumns,
		PrimaryKey: []*schema.Column{HierarchiesColumns[0]},
	}
	// HierarchyRelationColumns holds the columns for the "hierarchy_relation" table.
	HierarchyRelationColumns = []*schema.Column{
		{Name: "parent_id", Type: field.TypeInt},
		{Name: "child_id", Type: field.TypeInt},
	}
	// HierarchyRelationTable holds the schema information for the "hierarchy_relation" table.
	HierarchyRelationTable = &schema.Table{
		Name:       "hierarchy_relation",
		Columns:    HierarchyRelationColumns,
		PrimaryKey: []*schema.Column{HierarchyRelationColumns[0], HierarchyRelationColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "hierarchy_relation_parent_id",
				Columns:    []*schema.Column{HierarchyRelationColumns[0]},
				RefColumns: []*schema.Column{HierarchiesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "hierarchy_relation_child_id",
				Columns:    []*schema.Column{HierarchyRelationColumns[1]},
				RefColumns: []*schema.Column{HierarchiesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		HierarchiesTable,
		HierarchyRelationTable,
	}
)

func init() {
	HierarchyRelationTable.ForeignKeys[0].RefTable = HierarchiesTable
	HierarchyRelationTable.ForeignKeys[1].RefTable = HierarchiesTable
}
