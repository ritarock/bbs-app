package valueobject

type CommentID struct {
	value int
}

func NewCommentID(id int) CommentID {
	return CommentID{value: id}
}

func (id CommentID) Int() int {
	return id.value
}
