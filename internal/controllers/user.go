package controllers

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"time"

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

func (h *Handlers) SignIn(w http.ResponseWriter, r *http.Request) {
	var user models.Login
	if err := resp.ReadJSON(w, r, &user); err != nil {
		resp.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	result, err := services.GetUser(&user)
	if err != nil {
		if err == sql.ErrNoRows {
			resp.ErrorJSON(w, utils.ErrSqlNoRowsUser, http.StatusNotFound)
			return
		}
		resp.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	if err := utils.ComparePasswords(result.Password, user.Password); err != nil {
		resp.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	accessToken, err := utils.GenerateToken(result.ID, result.Username, result.Email, time.Duration(time.Hour*24))
	if err != nil {
		resp.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	result.Password = ""

	resp.WriteJSON(w, http.StatusOK, map[string]any{
		"access_token": accessToken,
		"message":      "success",
		"payload":      result,
	})

	log.Printf("%s logged in successfully", result.Email)
}

func (h *Handlers) GetUsers(w http.ResponseWriter, r *http.Request) {
	result, err := services.GetUsers()
	if err != nil {
		resp.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	resp.WriteJSON(w, http.StatusOK, result)
}
