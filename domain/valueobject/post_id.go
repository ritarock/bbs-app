package valueobject

type PostID struct {
	value int
}

func NewPostID(id int) PostID {
	return PostID{value: id}
}

func (id PostID) Int() int {
	return id.value
}
