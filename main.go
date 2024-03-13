package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/ekokurniawann/startup/auth"
	"github.com/ekokurniawann/startup/handler"
	"github.com/ekokurniawann/startup/helper"
	"github.com/ekokurniawann/startup/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=ekokurniawan password=123456 dbname=startup port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService()

	userHandler := handler.NewUserHandler(userService, authService)

	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/users", userHandler.RegisterUser)
	mux.HandleFunc("/api/v1/sessions", userHandler.Login)
	mux.HandleFunc("/api/v1/email_checkers", userHandler.CheckEmailAvailability)
	mux.HandleFunc("/api/v1/avatars", func(w http.ResponseWriter, r *http.Request) {
		authMiddleware(authService, userService, http.HandlerFunc(userHandler.UploadAvatar)).ServeHTTP(w, r)
	})

	server := &http.Server{
		Addr:    ":3000",
		Handler: mux,
	}

	log.Println("Starting server on :3000")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Error starting server: %s\n", err)
	}
}

func respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func authMiddleware(authService auth.Service, userService user.Service, nextHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if !strings.Contains(authHeader, " ") {
			response := helper.APIResponse("Authorization header format is invalid", http.StatusUnauthorized, "error", nil)
			respondJSON(w, http.StatusUnauthorized, response)
			return
		}

		parsedToken := ""
		stringToken := strings.Split(authHeader, " ")
		if len(stringToken) == 2 {
			parsedToken = stringToken[1]
		} else {
			response := helper.APIResponse("Token not found in Authorization header", http.StatusUnauthorized, "error", nil)
			respondJSON(w, http.StatusUnauthorized, response)
			return
		}

		token, err := authService.ValidateToken(parsedToken)
		if err != nil {
			response := helper.APIResponse("Token is invalid or expired", http.StatusUnauthorized, "error", nil)
			respondJSON(w, http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.APIResponse("Invalid token format", http.StatusUnauthorized, "error", nil)
			respondJSON(w, http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.APIResponse("User not found", http.StatusUnauthorized, "error", nil)
			respondJSON(w, http.StatusUnauthorized, response)
			return
		}

		r = r.WithContext(context.WithValue(r.Context(), handler.UserContextKey, user))

		nextHandler.ServeHTTP(w, r)
	})
}
