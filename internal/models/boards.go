package models

type (
	CreateBoard struct {
		ID         string `json:"id"`
		BoardTitle string `json:"board_title"`
	}

	BoardColums struct {
		ID              string `json:"board_column_id"`
		BoardColumnName string `json:"board_column_name"`
		BoardId         string `json:"board_id"`
	}
)
