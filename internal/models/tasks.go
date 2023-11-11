package models

type (
	Task struct {
		ID            string    `json:"id"`
		Title         string    `json:"task_title"`
		Description   string    `json:"task_description"`
		BoardColumnID string    `json:"task_board_column_id"`
		BoardID       string    `json:"task_board_id"`
		SubTasks      []SubTask `json:"sub_tasks"`
	}

	SubTask struct {
		ID        string `json:"subtask_id"`
		Name      string `json:"subtask_name"`
		Completed bool   `json:"subtask_completed"`
		TaskID    string `json:"subtask_task_id"`
	}
)
