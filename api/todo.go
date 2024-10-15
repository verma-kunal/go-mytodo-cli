package api

import "time"

// a slice of type Todo
type Todos struct {
	Todos []Todo `json:"todos"`
}

type Todo struct {
	Id         int       `json:"id"`
	Owner      string    `json:"owner"`
	Title      string    `json:"title"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	IsComplete bool      `json:"isComplete"`
}
