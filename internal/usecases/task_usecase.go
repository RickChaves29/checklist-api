package usecases

import (
	"github.com/RickChaves29/checklist-api/internal/dtos"
	"github.com/RickChaves29/checklist-api/internal/entities"
	"github.com/RickChaves29/checklist-api/internal/interfaces"
	"github.com/RickChaves29/checklist-api/internal/repositories"
)

type TaskUsecase struct {
	TaskRepository interfaces.ITaskRepository
}

func NewCreateTaskUsecase(taskRepository repositories.Repository) *TaskUsecase {
	return &TaskUsecase{
		TaskRepository: &taskRepository,
	}
}
func (uc *TaskUsecase) Save(data dtos.CreateTaskDTO) error {
	task, err := entities.NewTask(data.Title, data.Description)

	if err != nil {
		return err
	}
	err = uc.TaskRepository.Save(task)
	if err != nil {
		return err
	}
	return nil
}

func (uc *TaskUsecase) FindAll() ([]dtos.TaskDTO, error) {
	tasks, err := uc.TaskRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (uc *TaskUsecase) FindById(id int) (dtos.TaskDTO, error) {
	task, err := uc.TaskRepository.FindById(id)
	if err != nil {
		return task, err
	}
	return task, nil
}

func (uc *TaskUsecase) Update(id int, done bool) error {
	err := uc.TaskRepository.Update(id, done)
	if err != nil {
		return err
	}
	return nil
}

func (uc *TaskUsecase) Delete(id int) error {
	err := uc.TaskRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
