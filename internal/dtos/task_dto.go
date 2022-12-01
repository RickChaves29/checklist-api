package dtos

type CreateTaskDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type TaskDTO struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type UpdateTaskDTO struct {
	Done bool `json:"done"`
}
