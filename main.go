package main

import (
	"log"
	"net/http"

	"github.com/ekokurniawann/startup/auth"
	"github.com/ekokurniawann/startup/handler"
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
	mux.HandleFunc("/api/v1/avatars", userHandler.UploadAvatar)

	server := &http.Server{
		Addr:    ":3000",
		Handler: mux,
	}

	log.Println("Starting server on :3000")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Error starting server: %s\n", err)
	}
}
