// Code generated by ent, DO NOT EDIT.

package post

import (
	"time"
)

const (
	// Label holds the string label denoting the post type in the database.
	Label = "post"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldPostedAt holds the string denoting the posted_at field in the database.
	FieldPostedAt = "posted_at"
	// EdgeComments holds the string denoting the comments edge name in mutations.
	EdgeComments = "comments"
	// Table holds the table name of the post in the database.
	Table = "posts"
	// CommentsTable is the table that holds the comments relation/edge.
	CommentsTable = "comments"
	// CommentsInverseTable is the table name for the Comment entity.
	// It exists in this package in order to avoid circular dependency with the "comment" package.
	CommentsInverseTable = "comments"
	// CommentsColumn is the table column denoting the comments relation/edge.
	CommentsColumn = "post_id"
)

// Columns holds all SQL columns for post fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldContent,
	FieldPostedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultPostedAt holds the default value on creation for the "posted_at" field.
	DefaultPostedAt time.Time
)