package models

// a slice of type Todo
type Todos struct {
	Todos []Todo `json:"todos"`
}

type Todo struct {
	Id     int    `json:"id"`
	Owner  string `json:"owner"`
	Title  *string `json:"title"` // changeable as pointer
	Status *string `json:"status"` // changeable as pointer
}
