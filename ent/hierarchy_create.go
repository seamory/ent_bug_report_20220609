// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/bug/ent/hierarchy"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// HierarchyCreate is the builder for creating a Hierarchy entity.
type HierarchyCreate struct {
	config
	mutation *HierarchyMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (hc *HierarchyCreate) SetName(s string) *HierarchyCreate {
	hc.mutation.SetName(s)
	return hc
}

// SetDescribe sets the "describe" field.
func (hc *HierarchyCreate) SetDescribe(s string) *HierarchyCreate {
	hc.mutation.SetDescribe(s)
	return hc
}

// SetSort sets the "sort" field.
func (hc *HierarchyCreate) SetSort(i int) *HierarchyCreate {
	hc.mutation.SetSort(i)
	return hc
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (hc *HierarchyCreate) SetNillableSort(i *int) *HierarchyCreate {
	if i != nil {
		hc.SetSort(*i)
	}
	return hc
}

// AddChildIDs adds the "child" edge to the Hierarchy entity by IDs.
func (hc *HierarchyCreate) AddChildIDs(ids ...int) *HierarchyCreate {
	hc.mutation.AddChildIDs(ids...)
	return hc
}

// AddChild adds the "child" edges to the Hierarchy entity.
func (hc *HierarchyCreate) AddChild(h ...*Hierarchy) *HierarchyCreate {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return hc.AddChildIDs(ids...)
}

// AddParentIDs adds the "parent" edge to the Hierarchy entity by IDs.
func (hc *HierarchyCreate) AddParentIDs(ids ...int) *HierarchyCreate {
	hc.mutation.AddParentIDs(ids...)
	return hc
}

// AddParent adds the "parent" edges to the Hierarchy entity.
func (hc *HierarchyCreate) AddParent(h ...*Hierarchy) *HierarchyCreate {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return hc.AddParentIDs(ids...)
}

// Mutation returns the HierarchyMutation object of the builder.
func (hc *HierarchyCreate) Mutation() *HierarchyMutation {
	return hc.mutation
}

// Save creates the Hierarchy in the database.
func (hc *HierarchyCreate) Save(ctx context.Context) (*Hierarchy, error) {
	var (
		err  error
		node *Hierarchy
	)
	hc.defaults()
	if len(hc.hooks) == 0 {
		if err = hc.check(); err != nil {
			return nil, err
		}
		node, err = hc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HierarchyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = hc.check(); err != nil {
				return nil, err
			}
			hc.mutation = mutation
			if node, err = hc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(hc.hooks) - 1; i >= 0; i-- {
			if hc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = hc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, hc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (hc *HierarchyCreate) SaveX(ctx context.Context) *Hierarchy {
	v, err := hc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hc *HierarchyCreate) Exec(ctx context.Context) error {
	_, err := hc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hc *HierarchyCreate) ExecX(ctx context.Context) {
	if err := hc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (hc *HierarchyCreate) defaults() {
	if _, ok := hc.mutation.Sort(); !ok {
		v := hierarchy.DefaultSort
		hc.mutation.SetSort(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (hc *HierarchyCreate) check() error {
	if _, ok := hc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Hierarchy.name"`)}
	}
	if _, ok := hc.mutation.Describe(); !ok {
		return &ValidationError{Name: "describe", err: errors.New(`ent: missing required field "Hierarchy.describe"`)}
	}
	if _, ok := hc.mutation.Sort(); !ok {
		return &ValidationError{Name: "sort", err: errors.New(`ent: missing required field "Hierarchy.sort"`)}
	}
	return nil
}

func (hc *HierarchyCreate) sqlSave(ctx context.Context) (*Hierarchy, error) {
	_node, _spec := hc.createSpec()
	if err := sqlgraph.CreateNode(ctx, hc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (hc *HierarchyCreate) createSpec() (*Hierarchy, *sqlgraph.CreateSpec) {
	var (
		_node = &Hierarchy{config: hc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: hierarchy.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: hierarchy.FieldID,
			},
		}
	)
	if value, ok := hc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: hierarchy.FieldName,
		})
		_node.Name = value
	}
	if value, ok := hc.mutation.Describe(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: hierarchy.FieldDescribe,
		})
		_node.Describe = value
	}
	if value, ok := hc.mutation.Sort(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: hierarchy.FieldSort,
		})
		_node.Sort = value
	}
	if nodes := hc.mutation.ChildIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   hierarchy.ChildTable,
			Columns: hierarchy.ChildPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: hierarchy.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := hc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   hierarchy.ParentTable,
			Columns: hierarchy.ParentPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: hierarchy.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// HierarchyCreateBulk is the builder for creating many Hierarchy entities in bulk.
type HierarchyCreateBulk struct {
	config
	builders []*HierarchyCreate
}

// Save creates the Hierarchy entities in the database.
func (hcb *HierarchyCreateBulk) Save(ctx context.Context) ([]*Hierarchy, error) {
	specs := make([]*sqlgraph.CreateSpec, len(hcb.builders))
	nodes := make([]*Hierarchy, len(hcb.builders))
	mutators := make([]Mutator, len(hcb.builders))
	for i := range hcb.builders {
		func(i int, root context.Context) {
			builder := hcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*HierarchyMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, hcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, hcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, hcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (hcb *HierarchyCreateBulk) SaveX(ctx context.Context) []*Hierarchy {
	v, err := hcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hcb *HierarchyCreateBulk) Exec(ctx context.Context) error {
	_, err := hcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hcb *HierarchyCreateBulk) ExecX(ctx context.Context) {
	if err := hcb.Exec(ctx); err != nil {
		panic(err)
	}
}
