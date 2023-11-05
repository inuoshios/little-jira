package services

import (
	"context"
	"fmt"
	"time"

	"github.com/inuoshios/little-jira/internal/models"
)

func CreateBoard(board *models.CreateBoard) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `insert into boards (board_title) values ($1) returning id;`

	if err := db.QueryRowContext(ctx, query, board.BoardTitle).Scan(&board.ID); err != nil {
		return "", fmt.Errorf("error creating boards: %w", err)
	}

	return board.ID, nil
}

func CreateBoardColumn(boardColumn *models.BoardColumns) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `insert into board_columns
	 (board_column_name, board_id) values ($1, $2) returning id
	`

	if err := db.QueryRowContext(ctx, query, boardColumn.BoardColumnName, boardColumn.BoardId).Scan(&boardColumn.ID); err != nil {
		return "", fmt.Errorf("error creating board column: %w", err)
	}

	return boardColumn.ID, nil
}

func GetBoard(payload *models.CreateBoard) (*models.CreateBoard, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id board_title from boards where id = $1`
	var boards models.CreateBoard

	if err := db.QueryRowContext(ctx, query, &payload.ID).Scan(boards.ID, boards.BoardTitle); err != nil {
		return nil, err
	}

	return &boards, nil
}

func GetBoardsWithBoardColums() {}
