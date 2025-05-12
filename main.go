package main

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/helmigandi/go-restapi-1/app"
	"github.com/helmigandi/go-restapi-1/controller"
	"github.com/helmigandi/go-restapi-1/helper"
	"github.com/helmigandi/go-restapi-1/middleware"
	"github.com/helmigandi/go-restapi-1/repository"
	"github.com/helmigandi/go-restapi-1/service"
	"net/http"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
