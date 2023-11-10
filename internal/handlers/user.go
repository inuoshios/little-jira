package handlers

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	entranslations "github.com/go-playground/validator/v10/translations/en"
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
	// v creates an instance of to validate
	var v = validator.New()
	enTranslator := en.New()
	uni = ut.New(enTranslator, enTranslator)
	trans, _ := uni.GetTranslator("en")
	_ = entranslations.RegisterDefaultTranslations(v, trans)

	var user = &models.User{}
	if err := resp.ReadJSON(w, r, user); err != nil {
		_ = resp.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	err := v.Struct(user)
	if err != nil {
		var validationError validator.ValidationErrors
		if errors.As(err, &validationError) {
			for _, e := range validationError {
				errorMessage := e.Translate(trans)
				_ = resp.ErrorJSON(w, errors.New(errorMessage), http.StatusInternalServerError)
				return
			}
		}
	}

	hashPassword, err := utils.Hash(user.Password)
	if err != nil {
		_ = resp.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
	user.Password = hashPassword

	result, err := services.CreateUser(user)
	if err != nil {
		_ = resp.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	_ = resp.WriteJSON(w, http.StatusCreated, map[string]string{
		"status":  "success",
		"user_id": result,
	})
}

func (h *Handlers) SignIn(w http.ResponseWriter, r *http.Request) {
	var user models.Login
	if err := resp.ReadJSON(w, r, &user); err != nil {
		_ = resp.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	result, err := services.GetUser(&user)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			_ = resp.ErrorJSON(w, utils.ErrSqlNoRowsUser, http.StatusNotFound)
			return
		}
		_ = resp.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	if err := utils.ComparePasswords(result.Password, user.Password); err != nil {
		_ = resp.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	accessToken, err := utils.GenerateToken(result.ID, result.Username, result.Email, time.Hour*24)
	if err != nil {
		_ = resp.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	result.Password = ""

	_ = resp.WriteJSON(w, http.StatusOK, map[string]any{
		"access_token": accessToken,
		"message":      "success",
		"payload":      result,
	})

	log.Printf("%s logged in successfully", result.Email)
}

func (h *Handlers) GetUsers(w http.ResponseWriter, _ *http.Request) {
	result, err := services.GetUsers()
	if err != nil {
		_ = resp.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	_ = resp.WriteJSON(w, http.StatusOK, result)
}
