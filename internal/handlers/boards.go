package handlers

import (
	"database/sql"
	"errors"
	"net/http"

	resp "github.com/inuoshios/little-jira/internal/helpers"
	"github.com/inuoshios/little-jira/internal/models"
	"github.com/inuoshios/little-jira/internal/services"
	"github.com/inuoshios/little-jira/internal/utils"
)

func (h *Handlers) CreateBoard(w http.ResponseWriter, r *http.Request) {
	var boards = &models.CreateBoard{}

	if err := resp.ReadJSON(w, r, boards); err != nil {
		_ = resp.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	userId, ok := r.Context().Value(utils.AuthPayloadUserID).(string)
	if !ok {
		_ = resp.ErrorJSON(w, utils.ErrAuthorizationError, http.StatusBadRequest)
		return
	}

	if boards.UserID != userId {
		_ = resp.ErrorJSON(w, utils.ErrAuthorizationError, http.StatusUnauthorized)
		return
	}

	result, err := services.CreateBoard(boards)
	if err != nil {
		_ = resp.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	var boardColumnIds []string

	for _, col := range boards.BoardColumns {
		col.BoardId = result

		boardColumResponse, err := services.CreateBoardColumn(&col)
		if err != nil {
			_ = resp.ErrorJSON(w, err, http.StatusBadRequest)
			return
		}

		boardColumnIds = append(boardColumnIds, boardColumResponse)
	}

	_ = resp.WriteJSON(w, http.StatusCreated, map[string]any{
		"status":          "success",
		"board_id":        result,
		"board_column_id": boardColumnIds,
	})
}

// CreateBoardColumn make it create an array of board columns
func (h *Handlers) CreateBoardColumn(w http.ResponseWriter, r *http.Request) {
	var boards models.CreateBoard

	if err := resp.ReadJSON(w, r, &boards); err != nil {
		_ = resp.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	_, err := services.GetBoard(boards.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			_ = resp.ErrorJSON(w, errors.New("cannot find board"), http.StatusNotFound)
			return
		}
		_ = resp.ErrorJSON(w, errors.New("an error occurred while trying to get board"), http.StatusBadRequest)
		return
	}

	var boardColumnIds []string

	for _, boardColumn := range boards.BoardColumns {
		boardColumnId, err := services.CreateBoardColumn(&boardColumn)
		if err != nil {
			_ = resp.ErrorJSON(w, err, http.StatusBadRequest)
			return
		}

		boardColumnIds = append(boardColumnIds, boardColumnId)
	}

	_ = resp.WriteJSON(w, http.StatusCreated, map[string]any{
		"status":          "success",
		"board_column_id": boardColumnIds,
	})
}
