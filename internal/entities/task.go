package entities

import "errors"

type Task struct {
	Title       string
	Description string
}

func NewTask(title, description string) (*Task, error) {
	task := &Task{
		Title:       title,
		Description: description,
	}
	err := task.IsValide()
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (t *Task) IsValide() error {
	if t.Title == "" || len(t.Title) > 150 {
		return errors.New("invalid title")
	}

	if t.Description == "" {
		return errors.New("invalid description")
	}
	return nil
}
