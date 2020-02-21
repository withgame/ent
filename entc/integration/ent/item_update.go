// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/entc/integration/ent/item"
	"github.com/facebookincubator/ent/entc/integration/ent/predicate"
	"github.com/facebookincubator/ent/schema/field"
)

// ItemUpdate is the builder for updating Item entities.
type ItemUpdate struct {
	config
	hooks      []ent.Hook
	mutation   *ItemMutation
	predicates []predicate.Item
}

// Where adds a new predicate for the builder.
func (iu *ItemUpdate) Where(ps ...predicate.Item) *ItemUpdate {
	iu.predicates = append(iu.predicates, ps...)
	return iu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (iu *ItemUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(iu.hooks) == 0 {
		affected, err = iu.sqlSave(ctx)
	} else {
		var mut ent.Mutator = ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			mutation, ok := m.(*ItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			iu.mutation = mutation
			affected, err = iu.sqlSave(ctx)
			return affected, err
		})
		for _, hook := range iu.hooks {
			mut = hook(mut)
		}
		if _, err := mut.Mutate(ctx, iu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (iu *ItemUpdate) SaveX(ctx context.Context) int {
	affected, err := iu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iu *ItemUpdate) Exec(ctx context.Context) error {
	_, err := iu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iu *ItemUpdate) ExecX(ctx context.Context) {
	if err := iu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (iu *ItemUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   item.Table,
			Columns: item.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: item.FieldID,
			},
		},
	}
	if ps := iu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if n, err = sqlgraph.UpdateNodes(ctx, iu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{item.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ItemUpdateOne is the builder for updating a single Item entity.
type ItemUpdateOne struct {
	config
	hooks    []ent.Hook
	mutation *ItemMutation
}

// Save executes the query and returns the updated entity.
func (iuo *ItemUpdateOne) Save(ctx context.Context) (*Item, error) {
	var (
		err  error
		node *Item
	)
	if len(iuo.hooks) == 0 {
		node, err = iuo.sqlSave(ctx)
	} else {
		var mut ent.Mutator = ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			mutation, ok := m.(*ItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			iuo.mutation = mutation
			node, err = iuo.sqlSave(ctx)
			return node, err
		})
		for _, hook := range iuo.hooks {
			mut = hook(mut)
		}
		if _, err := mut.Mutate(ctx, iuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (iuo *ItemUpdateOne) SaveX(ctx context.Context) *Item {
	i, err := iuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return i
}

// Exec executes the query on the entity.
func (iuo *ItemUpdateOne) Exec(ctx context.Context) error {
	_, err := iuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iuo *ItemUpdateOne) ExecX(ctx context.Context) {
	if err := iuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (iuo *ItemUpdateOne) sqlSave(ctx context.Context) (i *Item, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   item.Table,
			Columns: item.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: item.FieldID,
			},
		},
	}
	id, ok := iuo.mutation.ID()
	if !ok {
		return nil, fmt.Errorf("missing Item.ID for update")
	}
	_spec.Node.ID.Value = id
	i = &Item{config: iuo.config}
	_spec.Assign = i.assignValues
	_spec.ScanValues = i.scanValues()
	if err = sqlgraph.UpdateNode(ctx, iuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{item.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return i, nil
}
