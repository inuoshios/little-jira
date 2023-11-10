package handlers

import (
	"net/http"

	resp "github.com/inuoshios/little-jira/internal/helpers"
	"github.com/inuoshios/little-jira/internal/models"
	"github.com/inuoshios/little-jira/internal/services"
	"github.com/inuoshios/little-jira/internal/utils"
)

func (h *Handlers) CreateBoard(w http.ResponseWriter, r *http.Request) {
	var boards = &models.CreateBoard{}

	if err := resp.ReadJSON(w, r, boards); err != nil {
		resp.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	userId, ok := r.Context().Value(utils.AuthPayloadUserID).(string)
	if !ok {
		resp.ErrorJSON(w, utils.ErrAuthorizationError, http.StatusBadRequest)
		return
	}

	if boards.UserID != userId {
		resp.ErrorJSON(w, utils.ErrAuthorizationError, http.StatusUnauthorized)
		return
	}

	result, err := services.CreateBoard(boards)
	if err != nil {
		resp.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	var boardColumnIds []string

	for _, col := range boards.BoardColumns {
		col.BoardId = result

		boardColumResponse, err := services.CreateBoardColumn(&col)
		if err != nil {
			resp.ErrorJSON(w, err, http.StatusBadRequest)
			return
		}

		boardColumnIds = append(boardColumnIds, boardColumResponse)
	}

	resp.WriteJSON(w, http.StatusCreated, map[string]any{
		"status":          "success",
		"board_id":        result,
		"board_column_id": boardColumnIds,
	})
}
