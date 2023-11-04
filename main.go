package main

import (
	"net/http"
	"yusufwdn/golang-restful-api/app"
	"yusufwdn/golang-restful-api/controller"
	"yusufwdn/golang-restful-api/helper"
	"yusufwdn/golang-restful-api/middleware"
	"yusufwdn/golang-restful-api/repository"
	"yusufwdn/golang-restful-api/service"

	"github.com/go-playground/validator/v10"

	_ "github.com/go-sql-driver/mysql"
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
