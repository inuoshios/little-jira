package models

type (
	CreateBoard struct {
		ID           string         `json:"id"`
		UserID       string         `json:"user_id"`
		BoardTitle   string         `json:"board_title"`
		BoardColumns []BoardColumns `json:"board_columns"`
	}

	BoardColumns struct {
		ID              string `json:"board_column_id"`
		BoardColumnName string `json:"board_column_name"`
		BoardId         string `json:"board_id"`
	}
)
