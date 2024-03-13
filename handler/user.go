package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/ekokurniawann/startup/auth"
	"github.com/ekokurniawann/startup/helper"
	"github.com/ekokurniawann/startup/user"
)

type contextKey string

const (
	UserContextKey contextKey = "currentUser"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
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

	token, err := h.authService.GenerateToken(newUser.ID)
	if err != nil {
		respone := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
		respondJSON(w, http.StatusBadRequest, respone)
		return
	}

	formatter := user.FormatUser(newUser, token)

	respone := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)
	respondJSON(w, http.StatusOK, respone)
}

func (h *userHandler) Login(w http.ResponseWriter, r *http.Request) {
	var input user.LoginInput

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := map[string]interface{}{"errors": errors}
		respone := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		respondJSON(w, http.StatusUnprocessableEntity, respone)
		return
	}

	loginUser, err := h.userService.Login(input)
	if err != nil {
		errorMessage := map[string]interface{}{"errors": err.Error()}
		respone := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		respondJSON(w, http.StatusUnprocessableEntity, respone)
		return
	}

	token, err := h.authService.GenerateToken(loginUser.ID)
	if err != nil {
		errorMessage := map[string]interface{}{"errors": err.Error()}
		respone := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		respondJSON(w, http.StatusUnprocessableEntity, respone)
		return
	}

	formatter := user.FormatUser(loginUser, token)

	respone := helper.APIResponse("Successfully logged in", http.StatusOK, "success", formatter)
	respondJSON(w, http.StatusOK, respone)
}

func (h *userHandler) CheckEmailAvailability(w http.ResponseWriter, r *http.Request) {
	var input user.CheckEmailInput

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		errorMessage := map[string]interface{}{"errors": err.Error()}
		response := helper.APIResponse("Email checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		respondJSON(w, http.StatusUnprocessableEntity, response)
		return
	}

	isEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errorMessage := map[string]interface{}{"errors": err.Error()}
		response := helper.APIResponse("Email checking failed", http.StatusInternalServerError, "error", errorMessage)
		respondJSON(w, http.StatusInternalServerError, response)
		return
	}

	var metaMessage string
	if isEmailAvailable {
		metaMessage = "Email is available"
	} else {
		metaMessage = "Email has been registered"
	}

	data := map[string]interface{}{
		"is_available": isEmailAvailable,
	}

	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
	respondJSON(w, http.StatusOK, response)
}

func (h *userHandler) UploadAvatar(w http.ResponseWriter, r *http.Request) {

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		data := map[string]interface{}{
			"is_uploaded": false,
		}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)
		respondJSON(w, http.StatusBadRequest, response)
		return
	}

	file, handler, err := r.FormFile("avatar")
	if err != nil {
		data := map[string]interface{}{
			"is_uploaded": false,
		}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)

		respondJSON(w, http.StatusBadRequest, response)
		return
	}
	defer file.Close()

	currentUser, ok := r.Context().Value(UserContextKey).(user.User)

	if !ok {
		response := helper.APIResponse("Failed to get current user", http.StatusInternalServerError, "error", nil)
		respondJSON(w, http.StatusInternalServerError, response)
		return
	}

	userID := currentUser.ID
	path := fmt.Sprintf("images/%d-%s", userID, handler.Filename)

	dst, err := os.Create(path)
	if err != nil {
		data := map[string]interface{}{
			"is_uploaded": false,
		}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)

		respondJSON(w, http.StatusBadRequest, response)
		return
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		data := map[string]interface{}{
			"is_uploaded": false,
		}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)

		respondJSON(w, http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.SaveAvatar(userID, path)
	if err != nil {
		data := map[string]interface{}{
			"is_uploaded": false,
		}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)

		respondJSON(w, http.StatusBadRequest, response)
		return
	}

	data := map[string]interface{}{
		"is_uploaded": true,
	}
	response := helper.APIResponse("Avatar successfully uploaded", http.StatusOK, "succes", data)

	respondJSON(w, http.StatusOK, response)
}
