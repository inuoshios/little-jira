package controllers

import (
	"errors"
	"net/http"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	resp "github.com/inuoshios/little-jira/internal/helpers"
	"github.com/inuoshios/little-jira/internal/models"
	"github.com/inuoshios/little-jira/internal/services"
	"github.com/inuoshios/little-jira/internal/utils"
)

type Handlers struct{}

var uni *ut.UniversalTranslator

func NewHandler() *Handlers {
	return &Handlers{}
}

func (h *Handlers) CreateUser(w http.ResponseWriter, r *http.Request) {
	// v creates an instance of the validate
	var v = validator.New()
	en := en.New()
	uni = ut.New(en, en)
	trans, _ := uni.GetTranslator("en")
	en_translations.RegisterDefaultTranslations(v, trans)

	var user = &models.User{}
	if err := resp.ReadJSON(w, r, user); err != nil {
		resp.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	err := v.Struct(user)
	if err != nil {
		if validationError, ok := err.(validator.ValidationErrors); ok {
			for _, e := range validationError {
				errorMessage := e.Translate(trans)
				resp.ErrorJSON(w, errors.New(errorMessage), http.StatusInternalServerError)
				return
			}
		} else {
			resp.ErrorJSON(w, err, http.StatusInternalServerError)
			return
		}
	}

	hashPassword, err := utils.Hash(user.Password)
	if err != nil {
		resp.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
	user.Password = hashPassword

	result, err := services.CreateUser(user)
	if err != nil {
		resp.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	resp.WriteJSON(w, http.StatusCreated, map[string]string{
		"status":  "success",
		"user_id": result,
	})
}

func (h *Handlers) GetUsers(w http.ResponseWriter, r *http.Request) {
	result, err := services.GetUsers()
	if err != nil {
		resp.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	resp.WriteJSON(w, http.StatusOK, result)
}
