package templates

import (
	"fmt"

	utils "github.com/DevopsGuyXD/Goku/Utils"
)

// ====================================== SIMPLE TEMPLATE
func Starter_Template(project string) {

	utils.Create_Single_Folder(project)

	createFolders(project)

	fmt.Println()
	fmt.Printf("\rCreating %v ✔\n\n", project)
}

// ====================================== CREATE FOLDER
func createFolders(project string) {
	subFolders := []string{"Routes", "Controller", "Config", "Models", "Utils"}

	for _, subFolder := range subFolders {

		folder := fmt.Sprintf("%v/%v", project, subFolder)

		utils.Create_Single_Folder(folder)
		Create_Files(project, subFolder)
	}
}

// ====================================== CREATE FILES
func Create_Files(project, subFolder string) {

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

	file := utils.Create_File(filePath)
	defer file.Close()

	data :=
		`PORT=":8000"
GOKU_VERSION="v1.0.0"`
	utils.Write_File(file, data)
}

// ====================================== MOD FILE
func modFile(project string) {

	filePath := fmt.Sprintf("%v/go.mod", project)

	file := utils.Create_File(filePath)
	defer file.Close()

	data := fmt.Sprintf(
		"module github.com/DevopsGuyXD/%v\n"+

			"go 1.23.5",

		project)

	utils.Write_File(file, data)
}

// ====================================== MAIN FILE
func mainFile(project string) {

	filePath := fmt.Sprintf("%v/main.go", project)

	file := utils.Create_File(filePath)
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
	util.Check_For_Nil(err)
}

// -------------------------- INIT
func init() {

	models.AppModels()

}`, project)
	utils.Write_File(file, data)
}

// ====================================== ROUTES FILE
func routesFile(project, subFolder string) {

	filePath := fmt.Sprintf("%v/%v/routes.go", project, subFolder)

	file := utils.Create_File(filePath)
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
	utils.Write_File(file, data)
}

// ====================================== CONTROLLER FILE
func controllerFile(project, subFolder string) {

	filePath := fmt.Sprintf("%v/%v/home.go", project, subFolder)

	file := utils.Create_File(filePath)
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
// @Produce json
// @Success 200 {array} map[string]interface{}
// @Router /health [get]
func GET_health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	util.InitEnvFile()

	message := "Healthy!!"

	json.NewEncoder(w).Encode(message)
}

// -------------------------- NOT ALLOWED CONTROLLER
func GET_NotAllowed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)

	message := "NA"

	json.NewEncoder(w).Encode(message)

}`, project)
	utils.Write_File(file, data)
}

// ====================================== CONFIG FILE
func configFile(project, subFolder string) {

	filePath := fmt.Sprintf("%v/%v/config.go", project, subFolder)

	file := utils.Create_File(filePath)
	defer file.Close()

	data :=
		`package config

import (
	"database/sql"
	"strings"
	
	utils "github.com/DevopsGuyXD/myapp/Utils"
	_ "modernc.org/sqlite"
)

// -------------------------- INIT DB
func InitDatabase(env string) *sql.DB {

	var database *sql.DB
	var err error

	if strings.ToLower(env) == "prod" {
		database, err = sql.Open("sqlite", "./Sqlite/app.db")
		utils.Check_For_Nil(err)
		
	} else if strings.ToLower(env) == "dev" {
		database, err = sql.Open("sqlite", "./Sqlite/test.db")
		utils.Check_For_Nil(err)
	}

	return database
}`
	utils.Write_File(file, data)
}

// ====================================== MODEL FILE
func modelFile(project, subFolder string) {

	filePath := fmt.Sprintf("%v/%v/models.go", project, subFolder)

	file := utils.Create_File(filePath)
	defer file.Close()

	data :=
		`package models

// -------------------------- MODELS
func AppModels(){

}`
	utils.Write_File(file, data)
}

// ====================================== UTIL FILE
func utilFile(project, subFolder string) {

	filePath := fmt.Sprintf("%v/%v/utils.go", project, subFolder)

	file := utils.Create_File(filePath)
	defer file.Close()

	data :=
		`package utils

import (
	"log"
	"github.com/joho/godotenv"
)

// -------------------------- Error Handling
func Check_For_Nil(err error) {
	if err != nil {
		log.Println(err)
		return
	}
}

// -------------------------- .env Init
func InitEnvFile() {
	err := godotenv.Load(".env")
	Check_For_Nil(err)
}`
	utils.Write_File(file, data)
}

// // ====================================== SWAGGER FILES
// func swagger() {
// 	calledFrom := utils.Called_From_Location()

// 	fmt.Println("Called from:", calledFrom)

// 	_, err := exec.Command("sh", "-c", fmt.Sprintf("swag init --dir %s", calledFrom)).Output()
// 	utils.Check_For_Nil(err)
// }
