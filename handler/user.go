package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ekokurniawann/startup/helper"
	"github.com/ekokurniawann/startup/user"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func (h *userHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var input user.RegisterUserInput

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := map[string]interface{}{"errors": errors}

		respone := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		respondJSON(w, http.StatusUnprocessableEntity, respone)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		respone := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
		respondJSON(w, http.StatusBadRequest, respone)
		return
	}

	formatter := user.FormatUser(newUser, "token")

	respone := helper.APIResponse("Account has been registered", http.StatusOK, "succes", formatter)
	respondJSON(w, http.StatusOK, respone)
}
