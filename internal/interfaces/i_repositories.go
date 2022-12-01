package interfaces

import (
	"github.com/RickChaves29/checklist-api/internal/dtos"
	"github.com/RickChaves29/checklist-api/internal/entities"
)

type ITaskRepository interface {
	Save(task *entities.Task) error
	FindById(id int) (dtos.TaskDTO, error)
	FindAll() ([]dtos.TaskDTO, error)
	Update(id int, done bool) error
	Delete(id int) error
}
