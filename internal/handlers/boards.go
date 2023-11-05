package handlers

import (
	"net/http"

	resp "github.com/inuoshios/little-jira/internal/helpers"
	"github.com/inuoshios/little-jira/internal/models"
	"github.com/inuoshios/little-jira/internal/services"
)

func (h *Handlers) CreateBoard(w http.ResponseWriter, r *http.Request) {
	var boards = &models.CreateBoard{}
	// var boardColumns []models.BoardColumns

	if err := resp.ReadJSON(w, r, boards); err != nil {
		resp.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	result, err := services.CreateBoard(boards)
	if err != nil {
		resp.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	var boardColumnIds []string

	for _, col := range boards.BoardColumns {
		// if err := resp.ReadJSON(w, r, &col); err != nil {
		// 	resp.ErrorJSON(w, err, http.StatusBadRequest)
		// 	return
		// }

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
