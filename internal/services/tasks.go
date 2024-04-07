package services

import (
	"context"
	"fmt"
	"time"

	"github.com/inuoshios/little-jira/internal/models"
)

func CreateTask(tasks *models.Task) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `insert into task (title, description, board_column_id, board_id)
	values ($1, $2, $3, $4) returning id;
	`
	if err := db.QueryRowContext(ctx, query, tasks.Title, tasks.Description, tasks.BoardColumnID, tasks.BoardID).Scan(&tasks.ID); err != nil {
		return "", fmt.Errorf("error creating tasks: %w", err)
	}

	return tasks.ID, nil
}

func CreateSubTasks(subTasks *models.SubTask) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `insert into sub_tasks (subtask_name, completed, task_id)
	values ($1, $2, $3)
	`

	if err := db.QueryRowContext(ctx, query, subTasks.Name, subTasks.Completed, subTasks.TaskID).Scan(&subTasks.ID); err != nil {
		return "", fmt.Errorf("error creating sub tasks: %w", err)
	}

	return subTasks.ID, nil
}

func GetTasks() ([]*models.Task, error) {
	// ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// defer cancel()

	// // query :=  ``
	return nil, fmt.Errorf("")
}
