// Code generated by entc, DO NOT EDIT.

package ent

import (
	"bbs-app/backend/ent/comment"
	"bbs-app/backend/ent/topic"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TopicCreate is the builder for creating a Topic entity.
type TopicCreate struct {
	config
	mutation *TopicMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (tc *TopicCreate) SetName(s string) *TopicCreate {
	tc.mutation.SetName(s)
	return tc
}

// SetDetail sets the "detail" field.
func (tc *TopicCreate) SetDetail(s string) *TopicCreate {
	tc.mutation.SetDetail(s)
	return tc
}

// SetNillableDetail sets the "detail" field if the given value is not nil.
func (tc *TopicCreate) SetNillableDetail(s *string) *TopicCreate {
	if s != nil {
		tc.SetDetail(*s)
	}
	return tc
}

// SetCreatedAt sets the "created_at" field.
func (tc *TopicCreate) SetCreatedAt(t time.Time) *TopicCreate {
	tc.mutation.SetCreatedAt(t)
	return tc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tc *TopicCreate) SetNillableCreatedAt(t *time.Time) *TopicCreate {
	if t != nil {
		tc.SetCreatedAt(*t)
	}
	return tc
}

// SetUpdatedAt sets the "updated_at" field.
func (tc *TopicCreate) SetUpdatedAt(t time.Time) *TopicCreate {
	tc.mutation.SetUpdatedAt(t)
	return tc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tc *TopicCreate) SetNillableUpdatedAt(t *time.Time) *TopicCreate {
	if t != nil {
		tc.SetUpdatedAt(*t)
	}
	return tc
}

// AddCommentIDs adds the "comments" edge to the Comment entity by IDs.
func (tc *TopicCreate) AddCommentIDs(ids ...int) *TopicCreate {
	tc.mutation.AddCommentIDs(ids...)
	return tc
}

// AddComments adds the "comments" edges to the Comment entity.
func (tc *TopicCreate) AddComments(c ...*Comment) *TopicCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return tc.AddCommentIDs(ids...)
}

// Mutation returns the TopicMutation object of the builder.
func (tc *TopicCreate) Mutation() *TopicMutation {
	return tc.mutation
}

// Save creates the Topic in the database.
func (tc *TopicCreate) Save(ctx context.Context) (*Topic, error) {
	var (
		err  error
		node *Topic
	)
	tc.defaults()
	if len(tc.hooks) == 0 {
		if err = tc.check(); err != nil {
			return nil, err
		}
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TopicMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tc.check(); err != nil {
				return nil, err
			}
			tc.mutation = mutation
			if node, err = tc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			if tc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TopicCreate) SaveX(ctx context.Context) *Topic {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TopicCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TopicCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TopicCreate) defaults() {
	if _, ok := tc.mutation.Detail(); !ok {
		v := topic.DefaultDetail
		tc.mutation.SetDetail(v)
	}
	if _, ok := tc.mutation.CreatedAt(); !ok {
		v := topic.DefaultCreatedAt
		tc.mutation.SetCreatedAt(v)
	}
	if _, ok := tc.mutation.UpdatedAt(); !ok {
		v := topic.DefaultUpdatedAt
		tc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TopicCreate) check() error {
	if _, ok := tc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Topic.name"`)}
	}
	if _, ok := tc.mutation.Detail(); !ok {
		return &ValidationError{Name: "detail", err: errors.New(`ent: missing required field "Topic.detail"`)}
	}
	if _, ok := tc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Topic.created_at"`)}
	}
	if _, ok := tc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Topic.updated_at"`)}
	}
	return nil
}

func (tc *TopicCreate) sqlSave(ctx context.Context) (*Topic, error) {
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (tc *TopicCreate) createSpec() (*Topic, *sqlgraph.CreateSpec) {
	var (
		_node = &Topic{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: topic.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: topic.FieldID,
			},
		}
	)
	if value, ok := tc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: topic.FieldName,
		})
		_node.Name = value
	}
	if value, ok := tc.mutation.Detail(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: topic.FieldDetail,
		})
		_node.Detail = value
	}
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: topic.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := tc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: topic.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := tc.mutation.CommentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   topic.CommentsTable,
			Columns: []string{topic.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: comment.FieldID,
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

// TopicCreateBulk is the builder for creating many Topic entities in bulk.
type TopicCreateBulk struct {
	config
	builders []*TopicCreate
}

// Save creates the Topic entities in the database.
func (tcb *TopicCreateBulk) Save(ctx context.Context) ([]*Topic, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Topic, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TopicMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TopicCreateBulk) SaveX(ctx context.Context) []*Topic {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TopicCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TopicCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
