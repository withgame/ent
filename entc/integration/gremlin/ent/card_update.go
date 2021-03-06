// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/gremlin"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/__"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/g"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/p"
	"github.com/facebookincubator/ent/entc/integration/gremlin/ent/card"
	"github.com/facebookincubator/ent/entc/integration/gremlin/ent/predicate"
	"github.com/facebookincubator/ent/entc/integration/gremlin/ent/spec"
	"github.com/facebookincubator/ent/entc/integration/gremlin/ent/user"
)

// CardUpdate is the builder for updating Card entities.
type CardUpdate struct {
	config

	update_time *time.Time

	name         *string
	clearname    bool
	owner        map[string]struct{}
	spec         map[string]struct{}
	clearedOwner bool
	removedSpec  map[string]struct{}
	predicates   []predicate.Card
}

// Where adds a new predicate for the builder.
func (cu *CardUpdate) Where(ps ...predicate.Card) *CardUpdate {
	cu.predicates = append(cu.predicates, ps...)
	return cu
}

// SetName sets the name field.
func (cu *CardUpdate) SetName(s string) *CardUpdate {
	cu.name = &s
	return cu
}

// SetNillableName sets the name field if the given value is not nil.
func (cu *CardUpdate) SetNillableName(s *string) *CardUpdate {
	if s != nil {
		cu.SetName(*s)
	}
	return cu
}

// ClearName clears the value of name.
func (cu *CardUpdate) ClearName() *CardUpdate {
	cu.name = nil
	cu.clearname = true
	return cu
}

// SetOwnerID sets the owner edge to User by id.
func (cu *CardUpdate) SetOwnerID(id string) *CardUpdate {
	if cu.owner == nil {
		cu.owner = make(map[string]struct{})
	}
	cu.owner[id] = struct{}{}
	return cu
}

// SetNillableOwnerID sets the owner edge to User by id if the given value is not nil.
func (cu *CardUpdate) SetNillableOwnerID(id *string) *CardUpdate {
	if id != nil {
		cu = cu.SetOwnerID(*id)
	}
	return cu
}

// SetOwner sets the owner edge to User.
func (cu *CardUpdate) SetOwner(u *User) *CardUpdate {
	return cu.SetOwnerID(u.ID)
}

// AddSpecIDs adds the spec edge to Spec by ids.
func (cu *CardUpdate) AddSpecIDs(ids ...string) *CardUpdate {
	if cu.spec == nil {
		cu.spec = make(map[string]struct{})
	}
	for i := range ids {
		cu.spec[ids[i]] = struct{}{}
	}
	return cu
}

// AddSpec adds the spec edges to Spec.
func (cu *CardUpdate) AddSpec(s ...*Spec) *CardUpdate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cu.AddSpecIDs(ids...)
}

// ClearOwner clears the owner edge to User.
func (cu *CardUpdate) ClearOwner() *CardUpdate {
	cu.clearedOwner = true
	return cu
}

// RemoveSpecIDs removes the spec edge to Spec by ids.
func (cu *CardUpdate) RemoveSpecIDs(ids ...string) *CardUpdate {
	if cu.removedSpec == nil {
		cu.removedSpec = make(map[string]struct{})
	}
	for i := range ids {
		cu.removedSpec[ids[i]] = struct{}{}
	}
	return cu
}

// RemoveSpec removes spec edges to Spec.
func (cu *CardUpdate) RemoveSpec(s ...*Spec) *CardUpdate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cu.RemoveSpecIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (cu *CardUpdate) Save(ctx context.Context) (int, error) {
	if cu.update_time == nil {
		v := card.UpdateDefaultUpdateTime()
		cu.update_time = &v
	}
	if cu.name != nil {
		if err := card.NameValidator(*cu.name); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"name\": %v", err)
		}
	}
	if len(cu.owner) > 1 {
		return 0, errors.New("ent: multiple assignments on a unique edge \"owner\"")
	}
	return cu.gremlinSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CardUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CardUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CardUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *CardUpdate) gremlinSave(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := cu.gremlin().Query()
	if err := cu.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	if err, ok := isConstantError(res); ok {
		return 0, err
	}
	return res.ReadInt()
}

func (cu *CardUpdate) gremlin() *dsl.Traversal {
	type constraint struct {
		pred *dsl.Traversal // constraint predicate.
		test *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 1)
	v := g.V().HasLabel(card.Label)
	for _, p := range cu.predicates {
		p(v)
	}
	var (
		rv = v.Clone()
		_  = rv

		trs []*dsl.Traversal
	)
	if value := cu.update_time; value != nil {
		v.Property(dsl.Single, card.FieldUpdateTime, *value)
	}
	if value := cu.name; value != nil {
		v.Property(dsl.Single, card.FieldName, *value)
	}
	var properties []interface{}
	if cu.clearname {
		properties = append(properties, card.FieldName)
	}
	if len(properties) > 0 {
		v.SideEffect(__.Properties(properties...).Drop())
	}
	if cu.clearedOwner {
		tr := rv.Clone().InE(user.CardLabel).Drop().Iterate()
		trs = append(trs, tr)
	}
	for id := range cu.owner {
		v.AddE(user.CardLabel).From(g.V(id)).InV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(user.CardLabel).OutV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(card.Label, user.CardLabel, id)),
		})
	}
	for id := range cu.removedSpec {
		tr := rv.Clone().InE(spec.CardLabel).Where(__.OtherV().HasID(id)).Drop().Iterate()
		trs = append(trs, tr)
	}
	for id := range cu.spec {
		v.AddE(spec.CardLabel).From(g.V(id)).InV()
	}
	v.Count()
	if len(constraints) > 0 {
		constraints = append(constraints, &constraint{
			pred: rv.Count(),
			test: __.Is(p.GT(1)).Constant(&ConstraintError{msg: "update traversal contains more than one vertex"}),
		})
		v = constraints[0].pred.Coalesce(constraints[0].test, v)
		for _, cr := range constraints[1:] {
			v = cr.pred.Coalesce(cr.test, v)
		}
	}
	trs = append(trs, v)
	return dsl.Join(trs...)
}

// CardUpdateOne is the builder for updating a single Card entity.
type CardUpdateOne struct {
	config
	id string

	update_time *time.Time

	name         *string
	clearname    bool
	owner        map[string]struct{}
	spec         map[string]struct{}
	clearedOwner bool
	removedSpec  map[string]struct{}
}

// SetName sets the name field.
func (cuo *CardUpdateOne) SetName(s string) *CardUpdateOne {
	cuo.name = &s
	return cuo
}

// SetNillableName sets the name field if the given value is not nil.
func (cuo *CardUpdateOne) SetNillableName(s *string) *CardUpdateOne {
	if s != nil {
		cuo.SetName(*s)
	}
	return cuo
}

// ClearName clears the value of name.
func (cuo *CardUpdateOne) ClearName() *CardUpdateOne {
	cuo.name = nil
	cuo.clearname = true
	return cuo
}

// SetOwnerID sets the owner edge to User by id.
func (cuo *CardUpdateOne) SetOwnerID(id string) *CardUpdateOne {
	if cuo.owner == nil {
		cuo.owner = make(map[string]struct{})
	}
	cuo.owner[id] = struct{}{}
	return cuo
}

// SetNillableOwnerID sets the owner edge to User by id if the given value is not nil.
func (cuo *CardUpdateOne) SetNillableOwnerID(id *string) *CardUpdateOne {
	if id != nil {
		cuo = cuo.SetOwnerID(*id)
	}
	return cuo
}

// SetOwner sets the owner edge to User.
func (cuo *CardUpdateOne) SetOwner(u *User) *CardUpdateOne {
	return cuo.SetOwnerID(u.ID)
}

// AddSpecIDs adds the spec edge to Spec by ids.
func (cuo *CardUpdateOne) AddSpecIDs(ids ...string) *CardUpdateOne {
	if cuo.spec == nil {
		cuo.spec = make(map[string]struct{})
	}
	for i := range ids {
		cuo.spec[ids[i]] = struct{}{}
	}
	return cuo
}

// AddSpec adds the spec edges to Spec.
func (cuo *CardUpdateOne) AddSpec(s ...*Spec) *CardUpdateOne {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cuo.AddSpecIDs(ids...)
}

// ClearOwner clears the owner edge to User.
func (cuo *CardUpdateOne) ClearOwner() *CardUpdateOne {
	cuo.clearedOwner = true
	return cuo
}

// RemoveSpecIDs removes the spec edge to Spec by ids.
func (cuo *CardUpdateOne) RemoveSpecIDs(ids ...string) *CardUpdateOne {
	if cuo.removedSpec == nil {
		cuo.removedSpec = make(map[string]struct{})
	}
	for i := range ids {
		cuo.removedSpec[ids[i]] = struct{}{}
	}
	return cuo
}

// RemoveSpec removes spec edges to Spec.
func (cuo *CardUpdateOne) RemoveSpec(s ...*Spec) *CardUpdateOne {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cuo.RemoveSpecIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (cuo *CardUpdateOne) Save(ctx context.Context) (*Card, error) {
	if cuo.update_time == nil {
		v := card.UpdateDefaultUpdateTime()
		cuo.update_time = &v
	}
	if cuo.name != nil {
		if err := card.NameValidator(*cuo.name); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"name\": %v", err)
		}
	}
	if len(cuo.owner) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"owner\"")
	}
	return cuo.gremlinSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CardUpdateOne) SaveX(ctx context.Context) *Card {
	c, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return c
}

// Exec executes the query on the entity.
func (cuo *CardUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CardUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *CardUpdateOne) gremlinSave(ctx context.Context) (*Card, error) {
	res := &gremlin.Response{}
	query, bindings := cuo.gremlin(cuo.id).Query()
	if err := cuo.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	c := &Card{config: cuo.config}
	if err := c.FromResponse(res); err != nil {
		return nil, err
	}
	return c, nil
}

func (cuo *CardUpdateOne) gremlin(id string) *dsl.Traversal {
	type constraint struct {
		pred *dsl.Traversal // constraint predicate.
		test *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 1)
	v := g.V(id)
	var (
		rv = v.Clone()
		_  = rv

		trs []*dsl.Traversal
	)
	if value := cuo.update_time; value != nil {
		v.Property(dsl.Single, card.FieldUpdateTime, *value)
	}
	if value := cuo.name; value != nil {
		v.Property(dsl.Single, card.FieldName, *value)
	}
	var properties []interface{}
	if cuo.clearname {
		properties = append(properties, card.FieldName)
	}
	if len(properties) > 0 {
		v.SideEffect(__.Properties(properties...).Drop())
	}
	if cuo.clearedOwner {
		tr := rv.Clone().InE(user.CardLabel).Drop().Iterate()
		trs = append(trs, tr)
	}
	for id := range cuo.owner {
		v.AddE(user.CardLabel).From(g.V(id)).InV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(user.CardLabel).OutV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(card.Label, user.CardLabel, id)),
		})
	}
	for id := range cuo.removedSpec {
		tr := rv.Clone().InE(spec.CardLabel).Where(__.OtherV().HasID(id)).Drop().Iterate()
		trs = append(trs, tr)
	}
	for id := range cuo.spec {
		v.AddE(spec.CardLabel).From(g.V(id)).InV()
	}
	v.ValueMap(true)
	if len(constraints) > 0 {
		v = constraints[0].pred.Coalesce(constraints[0].test, v)
		for _, cr := range constraints[1:] {
			v = cr.pred.Coalesce(cr.test, v)
		}
	}
	trs = append(trs, v)
	return dsl.Join(trs...)
}
