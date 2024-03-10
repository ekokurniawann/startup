package main

import (
	"log"
	"net/http"

	"github.com/ekokurniawann/startup/handler"
	"github.com/ekokurniawann/startup/user"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
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

	userHandler := handler.NewUserHandler(userService)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Group(func(r chi.Router) {
		r.Post("/api/v1/users", userHandler.RegisterUser)
		r.Post("/api/v1/sessions", userHandler.Login)
	})

	http.ListenAndServe(":3000", r)

}
