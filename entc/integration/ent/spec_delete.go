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
	"github.com/facebookincubator/ent/entc/integration/ent/predicate"
	"github.com/facebookincubator/ent/entc/integration/ent/spec"
	"github.com/facebookincubator/ent/schema/field"
)

// SpecDelete is the builder for deleting a Spec entity.
type SpecDelete struct {
	config
	hooks      []ent.Hook
	mutation   *SpecMutation
	predicates []predicate.Spec
}

// Where adds a new predicate to the delete builder.
func (sd *SpecDelete) Where(ps ...predicate.Spec) *SpecDelete {
	sd.predicates = append(sd.predicates, ps...)
	return sd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sd *SpecDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(sd.hooks) == 0 {
		affected, err = sd.sqlExec(ctx)
	} else {
		var mut ent.Mutator = ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			mutation, ok := m.(*SpecMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sd.mutation = mutation
			affected, err = sd.sqlExec(ctx)
			return affected, err
		})
		for _, hook := range sd.hooks {
			mut = hook(mut)
		}
		if _, err := mut.Mutate(ctx, sd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (sd *SpecDelete) ExecX(ctx context.Context) int {
	n, err := sd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sd *SpecDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: spec.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: spec.FieldID,
			},
		},
	}
	if ps := sd.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, sd.driver, _spec)
}

// SpecDeleteOne is the builder for deleting a single Spec entity.
type SpecDeleteOne struct {
	sd *SpecDelete
}

// Exec executes the deletion query.
func (sdo *SpecDeleteOne) Exec(ctx context.Context) error {
	n, err := sdo.sd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{spec.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sdo *SpecDeleteOne) ExecX(ctx context.Context) {
	sdo.sd.ExecX(ctx)
}
