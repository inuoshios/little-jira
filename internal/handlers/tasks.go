package handlers

import (
	resp "github.com/inuoshios/little-jira/internal/helpers"
	"github.com/inuoshios/little-jira/internal/models"
	"github.com/inuoshios/little-jira/internal/services"
	"net/http"
)

func CreateTasks(w http.ResponseWriter, r *http.Request) {
	var task models.Task

	if err := resp.ReadJSON(w, r, &task); err != nil {
		_ = resp.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	result, err := services.CreateTask(&task)
	if err != nil {
		_ = resp.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	var subTasksIDs []string

	for _, subTask := range task.SubTasks {
		subTaskID, err := services.CreateSubTasks(&subTask)
		if err != nil {
			if err != nil {
				_ = resp.ErrorJSON(w, err, http.StatusBadRequest)
				return
			}

			subTasksIDs = append(subTasksIDs, subTaskID)
		}
	}

	_ = resp.WriteJSON(w, http.StatusCreated, map[string]any{
		"status":      "success",
		"task_id":     result,
		"subtask_ids": subTasksIDs,
	})
}
