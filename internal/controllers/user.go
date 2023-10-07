package controllers

import (
	"net/http"

	resp "github.com/inuoshios/little-jira/internal/helpers"
	"github.com/inuoshios/little-jira/internal/services"
)

type Handlers struct{}

func NewHandler() *Handlers {
	return &Handlers{}
}

func (h *Handlers) GetUsers(w http.ResponseWriter, r *http.Request) {
	result, err := services.GetUsers()
	if err != nil {
		resp.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	resp.WriteJSON(w, http.StatusOK, result)
}
