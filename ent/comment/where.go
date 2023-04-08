// Code generated by ent, DO NOT EDIT.

package comment

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/ritarock/bbs-app/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Comment {
	return predicate.Comment(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Comment {
	return predicate.Comment(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Comment {
	return predicate.Comment(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Comment {
	return predicate.Comment(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Comment {
	return predicate.Comment(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Comment {
	return predicate.Comment(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Comment {
	return predicate.Comment(sql.FieldLTE(FieldID, id))
}

// Content applies equality check predicate on the "content" field. It's identical to ContentEQ.
func Content(v string) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldContent, v))
}

// CommentedAt applies equality check predicate on the "commented_at" field. It's identical to CommentedAtEQ.
func CommentedAt(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldCommentedAt, v))
}

// PostID applies equality check predicate on the "post_id" field. It's identical to PostIDEQ.
func PostID(v int) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldPostID, v))
}

// ContentEQ applies the EQ predicate on the "content" field.
func ContentEQ(v string) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldContent, v))
}

// ContentNEQ applies the NEQ predicate on the "content" field.
func ContentNEQ(v string) predicate.Comment {
	return predicate.Comment(sql.FieldNEQ(FieldContent, v))
}

// ContentIn applies the In predicate on the "content" field.
func ContentIn(vs ...string) predicate.Comment {
	return predicate.Comment(sql.FieldIn(FieldContent, vs...))
}

// ContentNotIn applies the NotIn predicate on the "content" field.
func ContentNotIn(vs ...string) predicate.Comment {
	return predicate.Comment(sql.FieldNotIn(FieldContent, vs...))
}

// ContentGT applies the GT predicate on the "content" field.
func ContentGT(v string) predicate.Comment {
	return predicate.Comment(sql.FieldGT(FieldContent, v))
}

// ContentGTE applies the GTE predicate on the "content" field.
func ContentGTE(v string) predicate.Comment {
	return predicate.Comment(sql.FieldGTE(FieldContent, v))
}

// ContentLT applies the LT predicate on the "content" field.
func ContentLT(v string) predicate.Comment {
	return predicate.Comment(sql.FieldLT(FieldContent, v))
}

// ContentLTE applies the LTE predicate on the "content" field.
func ContentLTE(v string) predicate.Comment {
	return predicate.Comment(sql.FieldLTE(FieldContent, v))
}

// ContentContains applies the Contains predicate on the "content" field.
func ContentContains(v string) predicate.Comment {
	return predicate.Comment(sql.FieldContains(FieldContent, v))
}

// ContentHasPrefix applies the HasPrefix predicate on the "content" field.
func ContentHasPrefix(v string) predicate.Comment {
	return predicate.Comment(sql.FieldHasPrefix(FieldContent, v))
}

// ContentHasSuffix applies the HasSuffix predicate on the "content" field.
func ContentHasSuffix(v string) predicate.Comment {
	return predicate.Comment(sql.FieldHasSuffix(FieldContent, v))
}

// ContentEqualFold applies the EqualFold predicate on the "content" field.
func ContentEqualFold(v string) predicate.Comment {
	return predicate.Comment(sql.FieldEqualFold(FieldContent, v))
}

// ContentContainsFold applies the ContainsFold predicate on the "content" field.
func ContentContainsFold(v string) predicate.Comment {
	return predicate.Comment(sql.FieldContainsFold(FieldContent, v))
}

// CommentedAtEQ applies the EQ predicate on the "commented_at" field.
func CommentedAtEQ(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldCommentedAt, v))
}

// CommentedAtNEQ applies the NEQ predicate on the "commented_at" field.
func CommentedAtNEQ(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldNEQ(FieldCommentedAt, v))
}

// CommentedAtIn applies the In predicate on the "commented_at" field.
func CommentedAtIn(vs ...time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldIn(FieldCommentedAt, vs...))
}

// CommentedAtNotIn applies the NotIn predicate on the "commented_at" field.
func CommentedAtNotIn(vs ...time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldNotIn(FieldCommentedAt, vs...))
}

// CommentedAtGT applies the GT predicate on the "commented_at" field.
func CommentedAtGT(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldGT(FieldCommentedAt, v))
}

// CommentedAtGTE applies the GTE predicate on the "commented_at" field.
func CommentedAtGTE(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldGTE(FieldCommentedAt, v))
}

// CommentedAtLT applies the LT predicate on the "commented_at" field.
func CommentedAtLT(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldLT(FieldCommentedAt, v))
}

// CommentedAtLTE applies the LTE predicate on the "commented_at" field.
func CommentedAtLTE(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldLTE(FieldCommentedAt, v))
}

// PostIDEQ applies the EQ predicate on the "post_id" field.
func PostIDEQ(v int) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldPostID, v))
}

// PostIDNEQ applies the NEQ predicate on the "post_id" field.
func PostIDNEQ(v int) predicate.Comment {
	return predicate.Comment(sql.FieldNEQ(FieldPostID, v))
}

// PostIDIn applies the In predicate on the "post_id" field.
func PostIDIn(vs ...int) predicate.Comment {
	return predicate.Comment(sql.FieldIn(FieldPostID, vs...))
}

// PostIDNotIn applies the NotIn predicate on the "post_id" field.
func PostIDNotIn(vs ...int) predicate.Comment {
	return predicate.Comment(sql.FieldNotIn(FieldPostID, vs...))
}

// PostIDIsNil applies the IsNil predicate on the "post_id" field.
func PostIDIsNil() predicate.Comment {
	return predicate.Comment(sql.FieldIsNull(FieldPostID))
}

// PostIDNotNil applies the NotNil predicate on the "post_id" field.
func PostIDNotNil() predicate.Comment {
	return predicate.Comment(sql.FieldNotNull(FieldPostID))
}

// HasPost applies the HasEdge predicate on the "post" edge.
func HasPost() predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PostTable, PostColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPostWith applies the HasEdge predicate on the "post" edge with a given conditions (other predicates).
func HasPostWith(preds ...predicate.Post) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PostInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PostTable, PostColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Comment) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Comment) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Comment) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		p(s.Not())
	})
}
