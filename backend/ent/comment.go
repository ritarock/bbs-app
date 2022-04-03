// Code generated by entc, DO NOT EDIT.

package ent

import (
	"bbs-app/backend/ent/comment"
	"bbs-app/backend/ent/topic"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Comment is the model entity for the Comment schema.
type Comment struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Body holds the value of the "body" field.
	Body string `json:"body,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// TopicID holds the value of the "topic_id" field.
	TopicID int `json:"topic_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CommentQuery when eager-loading is set.
	Edges CommentEdges `json:"edges"`
}

// CommentEdges holds the relations/edges for other nodes in the graph.
type CommentEdges struct {
	// Topic holds the value of the topic edge.
	Topic *Topic `json:"topic,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// TopicOrErr returns the Topic value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CommentEdges) TopicOrErr() (*Topic, error) {
	if e.loadedTypes[0] {
		if e.Topic == nil {
			// The edge topic was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: topic.Label}
		}
		return e.Topic, nil
	}
	return nil, &NotLoadedError{edge: "topic"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Comment) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case comment.FieldID, comment.FieldTopicID:
			values[i] = new(sql.NullInt64)
		case comment.FieldBody:
			values[i] = new(sql.NullString)
		case comment.FieldCreatedAt, comment.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Comment", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Comment fields.
func (c *Comment) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case comment.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case comment.FieldBody:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field body", values[i])
			} else if value.Valid {
				c.Body = value.String
			}
		case comment.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case comment.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		case comment.FieldTopicID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field topic_id", values[i])
			} else if value.Valid {
				c.TopicID = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryTopic queries the "topic" edge of the Comment entity.
func (c *Comment) QueryTopic() *TopicQuery {
	return (&CommentClient{config: c.config}).QueryTopic(c)
}

// Update returns a builder for updating this Comment.
// Note that you need to call Comment.Unwrap() before calling this method if this Comment
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Comment) Update() *CommentUpdateOne {
	return (&CommentClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Comment entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Comment) Unwrap() *Comment {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Comment is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Comment) String() string {
	var builder strings.Builder
	builder.WriteString("Comment(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", body=")
	builder.WriteString(c.Body)
	builder.WriteString(", created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", topic_id=")
	builder.WriteString(fmt.Sprintf("%v", c.TopicID))
	builder.WriteByte(')')
	return builder.String()
}

// Comments is a parsable slice of Comment.
type Comments []*Comment

func (c Comments) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}