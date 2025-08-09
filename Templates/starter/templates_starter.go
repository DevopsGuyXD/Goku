package templates_starter

import (
	"fmt"
	"os"
	"path/filepath"
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

	case strings.Contains(folder, "go.mod"):
		data = mod_Data(project)
		file = create_open_file(folder)

	case strings.Contains(folder, "dockerfile"):
		data = DockerFile_Data()
		file = create_open_file(folder)

	case strings.Contains(folder, "Routes"):
		data = routes_Data(project)
		file = create_open_file(folder)

	case strings.Contains(folder, "Controller"):
		data = controller_Data(project)
		file = create_open_file(folder)

	case strings.Contains(folder, "Config"):
		data = config_Data(project)
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
		_, folderName := filepath.Split(folder_or_File)
		filePath = filepath.Join(folder_or_File, strings.ToLower(folderName)+".go")
	} else {
		filePath = folder_or_File
	}

	utils.Create_File([]string{filePath})
	file := utils.Open_File(filePath)

	return file
}

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
	"time"

	config "github.com/DevopsGuyXD/myapp/Config"
	models "github.com/DevopsGuyXD/%[1]v/Models"
	routes "github.com/DevopsGuyXD/%[1]v/Routes"
	utils "github.com/DevopsGuyXD/%[1]v/Utils"
)

// -------------------------- INIT
func init() {
	models.AppModels()
}

// -------------------------- Start server
func StartServer(port string) {
	server := &http.Server{
		Addr:         port,
		Handler:      routes.RouteCollection(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	err := server.ListenAndServe()
	utils.Check_For_Err(err)
}

// -------------------------- Main
func main() {
	port := ":" + config.EnvConfig()["PORT"]
	fmt.Printf("\nListening on http://localhost" + port + "\n")

	StartServer(port)
}`, project)

	return data
}

// ============================================================================ routes.go DATA
func routes_Data(project string) string {

	data := fmt.Sprintf(`package routes

import (
	controller "github.com/DevopsGuyXD/%[1]v/Controller"
	_ "github.com/DevopsGuyXD/%[1]v/docs"
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

	config "github.com/DevopsGuyXD/myapp/Config"
	utils "github.com/DevopsGuyXD/%[1]v/Utils"
)

// -------------------------- HOME CONTROLLER
// @Description Home
// @Produce json
// @Success 200 {array} map[string]interface{}
// @Router / [get]
func GET_home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	message := "Welcome to Goku " + config.EnvConfig()["VERSION"]

	err := json.NewEncoder(w).Encode(message)
	utils.Check_For_Err(err)
}

// -------------------------- HEALTH CONTROLLER
// @Description Health
// @Produce json
// @Success 200 {array} map[string]interface{}
// @Router /health [get]
func GET_health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	message := "Healthy!!"

	err := json.NewEncoder(w).Encode(message)
	utils.Check_For_Err(err)
}

// -------------------------- NOT ALLOWED CONTROLLER
func GET_NotAllowed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)

	message := "NA"

	err := json.NewEncoder(w).Encode(message)
	utils.Check_For_Err(err)

}`, project)

	return data
}

// ============================================================================ config.go DATA
func config_Data(project string) string {

	data := fmt.Sprintf(`
package config

import (
	"database/sql"
	
	utils "github.com/DevopsGuyXD/%[1]v/Utils"
	_ "modernc.org/sqlite"
)

// -------------------------- INIT DB
func InitDatabase() *sql.DB {
	database, err := sql.Open("sqlite", "./Sqlite/app.db")
	utils.Check_For_Err(err)

	return database
}

// -------------------------- ENV CONFIG
func EnvConfig() map[string]string {
	env := map[string]string{
		"PORT":    "8000",
		"VERSION": "v0.0.1",
	}

	return env
}`, project)

	return data
}

// ============================================================================ models.go DATA
var model_Data = `package models

// -------------------------- MODELS
func AppModels(){

}`

// ============================================================================ utils.go DATA
var utils_Data = `
package utils

import (
	"log"
)

// -------------------------- Error Handling
func Check_For_Err(err error) {
	if err != nil {
		log.Println(err)
		return
	}
}`

// ============================================================================ dockerfile DATA
func DockerFile_Data() string {

	data := `FROM golang:1.23.5 AS builder
WORKDIR /app

    ENV CGO_ENABLED=0 \
        GOOS=linux \
        GOARCH=amd64

    COPY go.mod go.sum ./
    RUN go mod download

    COPY . .

    RUN go build -o ./dist/app \
        && [ -d "./Sqlite" ] && cp -r ./Sqlite ./dist || echo "Sqlite folder not found"

FROM alpine:latest
WORKDIR /app

    RUN addgroup -S goku && adduser -S goku -G goku

    COPY --from=builder /app/dist/ .

    RUN chown -R goku:goku ./Sqlite \
        && chmod 775 ./Sqlite \
        && chmod 664 ./Sqlite/app.db

    USER goku

    EXPOSE 8000

    HEALTHCHECK --interval=30s --timeout=5s --start-period=10s --retries=3 \
    CMD wget --spider -q http://localhost:8000/health || exit 1

    CMD ["./app"]`

	return data
}
