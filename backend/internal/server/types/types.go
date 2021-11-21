package types

type Topic struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

type Comment struct {
	Id      string `json:"id"`
	TopicId string `json:"topic_id"`
	Body    string `json:"body"`
}
