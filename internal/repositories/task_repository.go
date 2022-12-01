package repositories

import (
	"database/sql"

	"github.com/RickChaves29/checklist-api/internal/dtos"
	"github.com/RickChaves29/checklist-api/internal/entities"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Save(task *entities.Task) error {
	insert, err := r.db.Prepare("INSERT INTO tasks (title, description) VALUES ($1, $2)")
	defer insert.Close()
	if err != nil {
		return err
	}
	_, err = insert.Exec(task.Title, task.Description)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) FindAll() ([]dtos.TaskDTO, error) {
	var tasks []dtos.TaskDTO
	rows, err := r.db.Query("SELECT * FROM tasks")
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var task dtos.TaskDTO
		rows.Scan(&task.ID, &task.Title, &task.Description, &task.Done)
		tasks = append(tasks, task)

	}
	return tasks, nil
}

func (r *Repository) FindById(id int) (dtos.TaskDTO, error) {
	var task dtos.TaskDTO
	err := r.db.QueryRow("SELECT * FROM tasks WHERE id = $1", id).Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Done,
	)

	if err != nil {
		return task, err
	}
	return task, nil
}

func (r *Repository) Update(id int, done bool) error {
	update, err := r.db.Prepare("UPDATE tasks SET done = $1 WHERE id = $2")
	defer update.Close()
	if err != nil {
		return err
	}
	_, err = update.Exec(done, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM tasks WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
