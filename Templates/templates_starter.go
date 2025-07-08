package templates

import (
	"fmt"
	"os"
	"strings"

	utils "github.com/DevopsGuyXD/Goku/Utils"
)

// ============================================================================ FILE CONTROLLER
func fileController(project, folder string) (*os.File, string) {

	var file *os.File
	var data string
	defer file.Close()

	switch {
	case strings.Contains(folder, "main.go"):
		data = main_Data(project)
		file = create_open_file(folder)

	case strings.Contains(folder, ".env"):
		data = env_Data
		file = create_open_file(folder)

	case strings.Contains(folder, "go.mod"):
		data = mod_Data(project)
		file = create_open_file(folder)

	case strings.Contains(folder, "Routes"):
		data = routes_Data(project)
		file = create_open_file(folder)

	case strings.Contains(folder, "Controller"):
		data = controller_Data(project)
		file = create_open_file(folder)

	case strings.Contains(folder, "Config"):
		data = config_Data
		file = create_open_file(folder)

	case strings.Contains(folder, "Models"):
		data = model_Data
		file = create_open_file(folder)

	case strings.Contains(folder, "Utils"):
		data = utils_Data
		file = create_open_file(folder)
	}

	return file, data
}

// ============================================================================ CREATE AND OPEN FOLDER
func create_open_file(folder_or_File string) *os.File {

	var filePath string

	if utils.Folder_Exists(folder_or_File) {
		folderParsed := strings.Split(folder_or_File, "\\")
		filePath = folder_or_File + "/" + strings.ToLower(folderParsed[1]) + ".go"
	} else {
		filePath = folder_or_File
	}

	utils.Create_File([]string{filePath})
	file := utils.Open_File(filePath)

	return file
}

// ============================================================================ .env DATA
var env_Data = `
PORT=":8000"
GOKU_VERSION="v1.0.0"`

// ============================================================================ go.mod DATA
func mod_Data(project string) string {

	data := fmt.Sprintf(`
module github.com/DevopsGuyXD/%v
go 1.23.5
`, project)

	return data
}

// ============================================================================ main.go DATA
func main_Data(project string) string {

	data := fmt.Sprintf(`
package main

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

	return data
}

// ============================================================================ routes.go DATA
func routes_Data(project string) string {

	data := fmt.Sprintf(`
package routes

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

	return data
}

// ============================================================================ controller.go DATA
func controller_Data(project string) string {

	data := fmt.Sprintf(`
package controller

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

	return data
}

// ============================================================================ config.go DATA
var config_Data = `
package config

import (
	"database/sql"
	
	utils "github.com/DevopsGuyXD/myapp/Utils"
	_ "modernc.org/sqlite"
)

// -------------------------- INIT DB
func InitDatabase() *sql.DB {

	database, err := sql.Open("sqlite", "./Sqlite/app.db")
	utils.Check_For_Nil(err)

	return database
}`

// ============================================================================ models.go DATA
var model_Data = `
package models

// -------------------------- MODELS
func AppModels(){

}`

// ============================================================================ utils.go DATA
var utils_Data = `
package utils

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
