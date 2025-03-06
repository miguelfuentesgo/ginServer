package domain

type Todo struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type CreateTodoRequest struct {
	Description string `json:"description"`
}
