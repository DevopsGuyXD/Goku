package templates

import (
	"fmt"

	utils "github.com/DevopsGuyXD/Goku/Utils"
)

// ====================================== SIMPLE TEMPLATE
func StarterTemplate(project string) {

	utils.CreateSingleFolder(project)

	createFolders(project)

	fmt.Println()
	fmt.Printf("\rCreating %v \n", project)
}

// ====================================== CREATE FOLDER
func createFolders(project string) {
	subFolders := []string{"Routes", "Controller", "Config", "Models", "Utils"}

	for _, subFolder := range subFolders {

		folder := fmt.Sprintf("%v/%v", project, subFolder)

		utils.CreateSingleFolder(folder)
		createFiles(project, subFolder)
	}
}

// ====================================== CREATE FILES
func createFiles(project, subFolder string) {

	mainFile(project)
	envFile(project)
	modFile(project)

	switch {
	case subFolder == "Routes":
		routesFile(project, subFolder)

	case subFolder == "Controller":
		controllerFile(project, subFolder)

	case subFolder == "Config":
		configFile(project, subFolder)

	case subFolder == "Models":
		modelFile(project, subFolder)

	case subFolder == "Utils":
		utilFile(project, subFolder)
	}
}

// ====================================== .ENV FILE
func envFile(project string) {

	filePath := fmt.Sprintf("%v/.env", project)

	file := utils.CreateFile(filePath)
	defer file.Close()

	data := fmt.Sprintf(
		"PORT=:8000\n" +
			"GOKU_VERSION=v1.0.0\n",
	)

	utils.WriteFile(file, data)
}

// ====================================== MOD FILE
func modFile(project string) {

	filePath := fmt.Sprintf("%v/go.mod", project)

	file := utils.CreateFile(filePath)
	defer file.Close()

	data := fmt.Sprintf(
		"module github.com/DevopsGuyXD/%v\n"+

			"go 1.23.5",

		project)

	utils.WriteFile(file, data)
}

// ====================================== MAIN FILE
func mainFile(project string) {

	filePath := fmt.Sprintf("%v/main.go", project)

	file := utils.CreateFile(filePath)
	defer file.Close()

	data := fmt.Sprintf(
		`package main

import (
	"fmt"
	"net/http"
	"os"

	models "github.com/DevopsGuyXD/%[1]v/Models"
	route "github.com/DevopsGuyXD/%[1]v/Routes"
	util "github.com/DevopsGuyXD/%[1]v/Utils"
)

// -------------------------- Main
func main() {
	util.InitEnvFile()

	fmt.Printf("Listening on http://localhost" + os.Getenv("PORT") + "\n")
	fmt.Printf("Swagger: http://localhost" + os.Getenv("PORT") + "/swagger/index.html" + "\n")

	server := route.RouteCollection()
	err := http.ListenAndServe(os.Getenv("PORT"), server)
	util.CheckForNil(err)
}

// -------------------------- INIT
func init() {

	models.AppModels()

}`, project)
	utils.WriteFile(file, data)
}

// ====================================== ROUTES FILE
func routesFile(project, subFolder string) {

	filePath := fmt.Sprintf("%v/%v/routes.go", project, subFolder)

	file := utils.CreateFile(filePath)
	defer file.Close()

	data := fmt.Sprintf(
		`package routes

import (
	controller "github.com/DevopsGuyXD/%[1]v/Controller"
	_ "github.com/DevopsGuyXD/myapp/docs"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

// -------------------------- Routes
func RouteCollection() chi.Router {

	router := chi.NewRouter()

	router.Get("/", controller.GET_home)
	router.Get("/health", controller.GET_health)
	router.Get("/swagger/*", httpSwagger.WrapHandler)

	return router
}`, project)
	utils.WriteFile(file, data)
}

// ====================================== CONTROLLER FILE
func controllerFile(project, subFolder string) {

	filePath := fmt.Sprintf("%v/%v/home.go", project, subFolder)

	file := utils.CreateFile(filePath)
	defer file.Close()

	data := fmt.Sprintf(
		`package controller

import (
	"encoding/json"
	"net/http"
	"os"

	util "github.com/DevopsGuyXD/%[1]v/Utils"
)

// -------------------------- HOME CONTROLLER
// @Description Home
// @Tags Home
// @Produce json
// @Success 200 {array} map[string]interface{}
// @Router / [get]
func GET_home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	util.InitEnvFile()

	message := "Welcome to Goku " + os.Getenv("GOKU_VERSION")

	json.NewEncoder(w).Encode(message)
}

// -------------------------- HEALTH CONTROLLER
// @Description Health
// @Tags Health
// @Produce json
// @Success 200 {array} map[string]interface{}
// @Router /health [get]
func GET_health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	util.InitEnvFile()

	message := "Healthy!!"

	json.NewEncoder(w).Encode(message)
}

	`, project)
	utils.WriteFile(file, data)
}

// ====================================== CONFIG FILE
func configFile(project, subFolder string) {

	filePath := fmt.Sprintf("%v/%v/config.go", project, subFolder)

	file := utils.CreateFile(filePath)
	defer file.Close()

	data :=
		`package config

import (
	"database/sql"
	utils "github.com/DevopsGuyXD/myapp/Utils"
	_ "modernc.org/sqlite"
)

// -------------------------- INIT DB
func InitDatabase() *sql.DB {
	database, err := sql.Open("sqlite", "./Sqlite/app.db")
	utils.CheckForNil(err)

	return database
}`
	utils.WriteFile(file, data)
}

// ====================================== MODEL FILE
func modelFile(project, subFolder string) {

	filePath := fmt.Sprintf("%v/%v/models.go", project, subFolder)

	file := utils.CreateFile(filePath)
	defer file.Close()

	data :=
		`package models

// -------------------------- MODELS
func AppModels(){

}`
	utils.WriteFile(file, data)
}

// ====================================== UTIL FILE
func utilFile(project, subFolder string) {

	filePath := fmt.Sprintf("%v/%v/utils.go", project, subFolder)

	file := utils.CreateFile(filePath)
	defer file.Close()

	data :=
		`package utils

import (
	"log"
	"github.com/joho/godotenv"
)

// -------------------------- Error Handling
func CheckForNil(err error) {
	if err != nil {
		log.Println(err)
		return
	}
}

// -------------------------- .env Init
func InitEnvFile() {
	err := godotenv.Load(".env")
	CheckForNil(err)
}`
	utils.WriteFile(file, data)
}

// // ====================================== SWAGGER FILES
// func swagger() {
// 	calledFrom := utils.CalledFromLocation()

// 	fmt.Println("Called from:", calledFrom)

// 	_, err := exec.Command("sh", "-c", fmt.Sprintf("swag init --dir %s", calledFrom)).Output()
// 	utils.CheckForNil(err)
// }
