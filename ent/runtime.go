// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/ritarock/bbs-app/ent/comment"
	"github.com/ritarock/bbs-app/ent/schema"
	"github.com/ritarock/bbs-app/ent/topic"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	commentFields := schema.Comment{}.Fields()
	_ = commentFields
	// commentDescBody is the schema descriptor for body field.
	commentDescBody := commentFields[0].Descriptor()
	// comment.DefaultBody holds the default value on creation for the body field.
	comment.DefaultBody = commentDescBody.Default.(string)
	// commentDescCreatedAt is the schema descriptor for created_at field.
	commentDescCreatedAt := commentFields[1].Descriptor()
	// comment.DefaultCreatedAt holds the default value on creation for the created_at field.
	comment.DefaultCreatedAt = commentDescCreatedAt.Default.(time.Time)
	// commentDescUpdatedAt is the schema descriptor for updated_at field.
	commentDescUpdatedAt := commentFields[2].Descriptor()
	// comment.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	comment.DefaultUpdatedAt = commentDescUpdatedAt.Default.(time.Time)
	topicFields := schema.Topic{}.Fields()
	_ = topicFields
	// topicDescCreatedAt is the schema descriptor for created_at field.
	topicDescCreatedAt := topicFields[2].Descriptor()
	// topic.DefaultCreatedAt holds the default value on creation for the created_at field.
	topic.DefaultCreatedAt = topicDescCreatedAt.Default.(time.Time)
	// topicDescUpdatedAt is the schema descriptor for updated_at field.
	topicDescUpdatedAt := topicFields[3].Descriptor()
	// topic.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	topic.DefaultUpdatedAt = topicDescUpdatedAt.Default.(time.Time)
}
