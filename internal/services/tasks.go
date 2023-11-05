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

	query := `insert into tasks (title, description, board_column_id, board_id)
	value s($1, $2, $3, $) returning id;
	`
	if err := db.QueryRowContext(ctx, query, tasks.Title, tasks.Description, tasks.BoardColumnID, tasks.BoardID).Scan(&tasks.ID); err != nil {
		return "", fmt.Errorf("error creating tasks: %w", err)
	}

	return tasks.ID, nil
}
