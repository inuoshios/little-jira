package models

type (
	Task struct {
		ID            string `json:"id"`
		Title         string `json:"task_title"`
		Description   string `json:"task_description"`
		BoardColumnID string `json:"task_board_column_id"`
		BoardID       string `json:"task_board_id"`
	}

	SubTask struct {
		Name      string `json:"subtask_name"`
		Completed bool   `json:"subtask_completed"`
		TaskID    string `json:"subtask_task_id"`
	}
)
