package templates

import (
	"fmt"
	"os"

	utils "github.com/DevopsGuyXD/Goku/Utils"
)

// ====================================== SIMPLE TEMPLATE
func StarterTemplate(projectName string) {

	folderName := projectName

	utils.CreateFolder(folderName)

	createFolders(folderName)

	fmt.Println()
	utils.InitSpinner("Creating Project")
	fmt.Printf("\rCreating %v ✅ \n", projectName)
}

// ====================================== CREATE FOLDER
func createFolders(folderName string) {
	subFolders := []string{"Routes", "Controller", "Config", "Models", "Utils"}

	for _, subFolder := range subFolders {

		folder := fmt.Sprintf("%v/%v", folderName, subFolder)

		utils.CreateFolder(folder)
		createFiles(folderName, subFolder)
	}
}

// ====================================== CREATE FILES
func createFiles(folderName, subFolder string) {

	mainFile(folderName)
	envFile(folderName)
	modFile(folderName)

	switch {
	case subFolder == "Routes":
		routesFile(folderName, subFolder)

	case subFolder == "Controller":
		controllerFile(folderName, subFolder)

	case subFolder == "Config":
		configFile(folderName, subFolder)

	case subFolder == "Models":
		modelFile(folderName, subFolder)

	case subFolder == "Utils":
		utilFile(folderName, subFolder)
	}
}

// ====================================== .ENV FILE
func envFile(folderName string) {
	file, err := os.Create(fmt.Sprintf("%v/.env", folderName))
	utils.CheckForNil(err)
	defer file.Close()

	data := fmt.Sprintf(
		"PORT=:8000\n" +
			"GOKU_VERSION=v1.0.0\n",
	)

	utils.WriteFile(file, data)
}

// ====================================== MOD FILE
func modFile(folderName string) {
	file, err := os.Create(fmt.Sprintf("%v/go.mod", folderName))
	utils.CheckForNil(err)
	defer file.Close()

	data := fmt.Sprintf(
		"module github.com/DevopsGuyXD/%v\n"+

			"go 1.23.5",

		folderName)

	utils.WriteFile(file, data)
}

// ====================================== MAIN FILE
func mainFile(folderName string) {

	file, err := os.Create(fmt.Sprintf("%v/main.go", folderName))
	utils.CheckForNil(err)
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

//-------------------------- Main
func main() {
	util.InitEnvFile()

	fmt.Printf("Listening on http://localhost" + os.Getenv("PORT") + "\n")
	fmt.Printf("Swagger: http://localhost" + os.Getenv("PORT") + "/swagger/index.html" + "\n")

	server := route.RouteCollection()
	err := http.ListenAndServe(os.Getenv("PORT"), server)
	util.CheckForNil(err)
}

//-------------------------- INIT
func init() {

	models.AppModels()

}`, folderName)
	utils.WriteFile(file, data)
}

// ====================================== ROUTES FILE
func routesFile(folderName, subFolder string) {
	file, err := os.Create(fmt.Sprintf("%v/%v/routes.go", folderName, subFolder))
	utils.CheckForNil(err)
	defer file.Close()

	data := fmt.Sprintf(
		`package routes

import (
	controller "github.com/DevopsGuyXD/%[1]v/Controller"
	_ "github.com/DevopsGuyXD/myapp/docs"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

//-------------------------- Routes
func RouteCollection() chi.Router {

	router := chi.NewRouter()

	router.Get("/", controller.GET_home)
	router.Get("/health", controller.GET_health)
	router.Get("/swagger/*", httpSwagger.WrapHandler)

	return router
}`, folderName)
	utils.WriteFile(file, data)
}

// ====================================== CONTROLLER FILE
func controllerFile(folderName, subFolder string) {
	file, err := os.Create(fmt.Sprintf("%v/%v/home.go", folderName, subFolder))
	utils.CheckForNil(err)
	defer file.Close()

	data := fmt.Sprintf(
		`package controller

import (
	"encoding/json"
	"net/http"
	"os"

	util "github.com/DevopsGuyXD/%[1]v/Utils"
)

// // -------------------------- HOME CONTROLLER
// GET_books godoc
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

// //-------------------------- HEALTH CONTROLLER
// GET_books godoc
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

	`, folderName)
	utils.WriteFile(file, data)
}

// ====================================== CONFIG FILE
func configFile(folderName, subFolder string) {

	file, err := os.Create(fmt.Sprintf("%v/%v/config.go", folderName, subFolder))
	utils.CheckForNil(err)
	defer file.Close()

	data := fmt.Sprintf(
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
}`)
	utils.WriteFile(file, data)
}

// ====================================== MODEL FILE
func modelFile(folderName, subFolder string) {

	file, err := os.Create(fmt.Sprintf("%v/%v/models.go", folderName, subFolder))
	utils.CheckForNil(err)
	defer file.Close()

	data := fmt.Sprintf(
		`package models

//-------------------------- MODELS
func AppModels(){

}`)
	utils.WriteFile(file, data)
}

// ====================================== UTIL FILE
func utilFile(folderName, subFolder string) {
	file, err := os.Create(fmt.Sprintf("%v/%v/utils.go", folderName, subFolder))
	utils.CheckForNil(err)
	defer file.Close()

	data :=
		`package utils

import (
	"log"
	"github.com/joho/godotenv"
)

//-------------------------- Error Handling
func CheckForNil(err error) {
	if err != nil {
		log.Println(err)
		return
	}
}

//-------------------------- .env Init
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
