package models

type Todo struct {
	UserId    int         `json:"user_id" bson:"user_id"`
	ID        interface{} `json:"id,omitempty" bson:"_id,omitempty"`
	Title     string      `json:"title" bson:"title"`
	Completed bool        `json:"completed" bson:"completed"`
}

type TodoUpdate struct {
	ModifiedCount int64 `json:"modified_count"`
	Result        Todo
}

type TodoDelete struct {
	DeletedCount int64 `json:"deleted_count"`
}
