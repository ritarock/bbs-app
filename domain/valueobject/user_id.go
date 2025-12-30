package valueobject

type UserID struct {
	value int
}

func NewUserID(id int) UserID {
	return UserID{value: id}
}

func (id UserID) Int() int {
	return id.value
}
