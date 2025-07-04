package templates

import (
	"fmt"

	utils "github.com/DevopsGuyXD/Goku/Utils"
)

// ====================================== ROUTES
func routes(crudName string) string {
	data := fmt.Sprintf(
		`	// -------------------------- %[1]v
	router.Route("/%[1]v", func(r chi.Router) {
		r.Get("/", controller.GET_%[1]v)
		r.Post("/", controller.POST_%[1]v)
		r.Get("/{id}", controller.GET_%[1]v_id)
		r.Put("/{id}", controller.UPDATE_%[1]v)
		r.Delete("/{id}", controller.DELETE_%[1]v)
	})
		`, crudName)

	return data
}

// ====================================== CONTROLLER
func controllers(crudName, project string) string {
	data := fmt.Sprintf(

		`package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	models "github.com/DevopsGuyXD/%[2]v/Models"
	"github.com/go-chi/chi/v5"
)

// -------------------------- %[1]v GET ALL
// @Description All %[1]v
// @Tags %[1]v
// @Produce json
// @Success 200 {array} map[string]interface{}
// @Router /%[1]v [get]
func GET_%[1]v(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := models.GET_%[1]v_all()
	
	if response != nil {
		json.NewEncoder(w).Encode(response)
	} else {
		json.NewEncoder(w).Encode("No records found")
	}
}

// -------------------------- %[1]v GET BY ID
// @Description Single %[1]v
// @Tags %[1]v
// @Produce json
// @Success 200 {array} map[string]interface{}
// @Router /%[1]v/{id} [get]
func GET_%[1]v_id(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
		
	id := chi.URLParam(r, "id")
	response := models.GET_%[1]v_by_id(id)

	if response != nil {
		json.NewEncoder(w).Encode(response)
	} else {
		response := fmt.Sprintf("No record with ID: %%s", id)
		json.NewEncoder(w).Encode(response)
	}
}

// -------------------------- %[1]v CREATE
// @Description Create a %[1]v
// @Tags %[1]v
// @Accept json
// @Produce json
// @Param data body object true "Add a new %[1]v"
// @Success 200 {array} map[string]interface{}
// @Router /%[1]v [post]
func POST_%[1]v(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := r.Body
	response := models.CREATE_%[1]v(request)

	json.NewEncoder(w).Encode(response)
}

// -------------------------- %[1]v UPDATE BY ID
// @Description Update %[1]v
// @Tags %[1]v
// @Produce json
// @Success 200 {array} map[string]interface{}
// @Router /%[1]v/{id} [put]
func UPDATE_%[1]v(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	id := chi.URLParam(r, "id")
	request := r.Body

	response := models.UPDATE_%[1]v(id, request)

	json.NewEncoder(w).Encode(response)
}

// -------------------------- %[1]v DELETE
// @Description Delete %[1]v
// @Tags %[1]v
// @Produce json
// @Success 200 {array} map[string]interface{}
// @Router /%[1]v/{id} [delete]
func DELETE_%[1]v(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := chi.URLParam(r, "id")
	response := models.DELETE_%[1]v(id)

	json.NewEncoder(w).Encode(response)
}`, crudName, project)

	return data
}

// ====================================== MODELS
func models(projectName, crudName string) string {
	data := fmt.Sprintf(
		`package models

import (
	"database/sql"
	"fmt"
	"io"
	"strings"


	config "github.com/DevopsGuyXD/%[1]v/Config"
	utils "github.com/DevopsGuyXD/%[1]v/Utils"
)

// -------------------------- CREATE %[2]v TABLE
func %[2]v() {

	db := initDB()
	defer db.Close()

	// Creating %[2]v table with PRIMARY KEY and DATE TIME
	table := "CREATE TABLE IF NOT EXISTS %[2]v(id INTEGER PRIMARY KEY AUTOINCREMENT, created_at DATETIME DEFAULT CURRENT_TIMESTAMP)"

	createDomainsTable, err := db.Prepare(table)
	utils.Check_For_Nil(err)
	createDomainsTable.Exec()

	columnsToAdd := []string{
		// Add columns here. Example,
		// "title VARCHAR(25)",
		// "author VARCHAR(25)",
		// "language VARCHAR(25)",
		// "pages INTEGER",
	}

	for _, col := range columnsToAdd {
		stmt := "ALTER TABLE %[2]v ADD COLUMN " + col + ";"
		_, err := db.Exec(stmt)
		if err != nil {
			if strings.Contains(err.Error(), "SQL logic error: duplicate column name") {
				continue
			}
			utils.Check_For_Nil(err)
		}
	}
}

// -------------------------- INIT DB
func initDB() *sql.DB {
	var database = config.InitDatabase()
	return database
}

// -------------------------- GET %[2]v ALL
func GET_%[2]v_all() []map[string]interface{} {

	query := "SELECT * FROM %[2]v"
	return get_Handler(query)
}

// -------------------------- GET %[2]v by ID
func GET_%[2]v_by_id(id string) []map[string]interface{} {

	query := fmt.Sprintf("SELECT * FROM %[2]v WHERE id = %%s", id)
	return get_Handler(query)
}

// -------------------------- CREATE %[2]v RECORD
func CREATE_%[2]v(request io.ReadCloser) string {
	return create_Handler(request)
}

// -------------------------- UPDATE %[2]v
func UPDATE_%[2]v(id string, request io.ReadCloser) string {
	return update_Handler(id, request)
}

// -------------------------- DELETE %[2]v
func DELETE_%[2]v(id string) string {
	return delete_Handler(id)
}
`, projectName, crudName)

	return data
}

// ====================================== MODELS HANDLERS
func modelHandlers(crudName string) string {
	data := fmt.Sprintf(`

// -------------------------- GET HANDLER
func get_Handler(query string) []map[string]interface{} {

	db := initDB()
	defer db.Close()

	var response []map[string]interface{}

	data, err := db.Query(query)
	utils.Check_For_Nil(err)
	defer data.Close()

	columns, err := data.Columns()
	utils.Check_For_Nil(err)

	for data.Next() {

		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))
		rowMap := make(map[string]interface{})

		for i := range values {
			valuePtrs[i] = &values[i]
		}

		err := data.Scan(valuePtrs...)
		utils.Check_For_Nil(err)

		for i, col := range columns {
			val := values[i]

			if b, ok := val.([]byte); ok {
				rowMap[col] = string(b)
			} else {
				rowMap[col] = val
			}
		}

		response = append(response, rowMap)
	}

	return response
}

// -------------------------- CREATE HANDLER
func create_Handler(request io.ReadCloser) string {

	db := initDB()
	defer db.Close()

	var response string
	var data map[string]interface{}

	err := json.NewDecoder(request).Decode(&data)
	utils.Check_For_Nil(err)

	columns := []string{}
	placeholders := []string{}
	values := []interface{}{}

	for k, v := range data {
		columns = append(columns, k)
		placeholders = append(placeholders, "?")
		values = append(values, v)
	}

	query := fmt.Sprintf("INSERT INTO %[1]v (%%s) VALUES (%%s)",
		strings.Join(columns, ", "),
		strings.Join(placeholders, ", "))

	res, err := db.Exec(query, values...)
	if err != nil {
		response = "DB insert error"

	} else {
		rowsAffected, err := res.RowsAffected()
		utils.Check_For_Nil(err)

		if rowsAffected != 0 {
			response = "Created successfully"
		} else {
			response = "No records created"
		}
	}

	return response
}

// -------------------------- UPDATE HANDLER
func update_Handler(id string, request io.ReadCloser) string {
	db := initDB()
	defer db.Close()

	var response string
	var data map[string]interface{}

	setParts := []string{}
	setValues := []interface{}{}

	err := json.NewDecoder(request).Decode(&data)
	utils.Check_For_Nil(err)

	for k, v := range data {
		setParts = append(setParts, fmt.Sprintf("%%s = ?", k))
		setValues = append(setValues, v)
	}

	if len(setParts) == 0 {
		return "No fields provided for update"
	}

	query := fmt.Sprintf("UPDATE %[1]v SET %%s WHERE id = ?",
		strings.Join(setParts, ", "))

	values := append(setValues, id)

	res, err := db.Exec(query, values...)
	if err != nil {
		response = "DB update error"
	} else {
		rowsAffected, err := res.RowsAffected()
		utils.Check_For_Nil(err)

		if rowsAffected != 0 {
			response = "Updated successfully"
		} else {
			response = "No records updated"
		}
	}

	return response
}

// -------------------------- DELETE HANDLER
func delete_Handler(id string) string {

	db := initDB()
	defer db.Close()

	var response string

	query := fmt.Sprintf("DELETE FROM %[1]v WHERE id=%%s", id)

	res, err := db.Exec(query)
	if err != nil {
		response = "DB delete error"

	} else {
		rowsAffected, err := res.RowsAffected()
		utils.Check_For_Nil(err)

		if rowsAffected != 0 {
			response = "Deleted successfully"
		} else {
			response = "No records deleted"
		}

	}

	return response
}
`, crudName)

	return data
}

// ====================================== DOCKER FILE
func DockerFile() {

	var data string

	filePath := "./dockerfile"

	file := utils.Create_File(filePath)
	defer file.Close()

	folder := "Sqlite"

	exists := utils.Folder_Exists(folder)

	if exists {

		data =
			`FROM golang:latest AS builder
WORKDIR /app

    ENV CGO_ENABLED=0 \
        GOOS=linux \
        GOARCH=amd64

    COPY go.mod go.sum ./
    RUN go mod download

    COPY . .

    RUN go build -o app

FROM alpine:latest
WORKDIR /app

    RUN addgroup -S appgroup && adduser -S appuser -G appgroup
    USER appuser

    COPY --from=builder /app/app .
    COPY --from=builder /app/.env .
    COPY --from=builder /app/Sqlite ./Sqlite

    EXPOSE 8000

    HEALTHCHECK --interval=30s --timeout=5s --start-period=10s --retries=3 \
    CMD wget --spider -q http://localhost:8000/health || exit 1

    CMD ["./app"]`
	} else {
		data = `FROM golang:1.23.5 AS builder
WORKDIR /app

    ENV CGO_ENABLED=0 \
        GOOS=linux \
        GOARCH=amd64

    COPY go.mod go.sum ./
    RUN go mod download

    COPY . .

    RUN go build -o app

FROM alpine:latest
WORKDIR /app

    RUN addgroup -S appgroup && adduser -S appuser -G appgroup
    USER appuser

    COPY --from=builder /app/app .
    COPY --from=builder /app/.env .

    EXPOSE 8000

    HEALTHCHECK --interval=30s --timeout=5s --start-period=10s --retries=3 \
    CMD wget --spider -q http://localhost:8000/health || exit 1

    CMD ["./app"]`
	}
	utils.Write_File(file, data)

	fmt.Println("\nAdded dockerfile ")
}

func modelImports() string {
	data := fmt.Sprintln(`
	
import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	utils "github.com/DevopsGuyXD/myapp/Utils"
)`)

	return data
}
