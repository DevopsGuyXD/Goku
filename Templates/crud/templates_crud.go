package templates_curd

import (
	"fmt"

	utils "github.com/DevopsGuyXD/Goku/Utils"
)

// ============================================================================ routes.go DATA
func routes_Data(crudName string) string {
	data := fmt.Sprintf(`	// -------------------------- %[1]v
	router.Route("/%[1]v", func(r chi.Router) {
		r.Get("/", %[1]v_c.GET_%[1]v)
		r.Post("/", %[1]v_c.POST_%[1]v)
		r.Get("/{id}", %[1]v_c.GET_%[1]v_id)
		r.Put("/{id}", %[1]v_c.UPDATE_%[1]v)
		r.Delete("/{id}", %[1]v_c.DELETE_%[1]v)
	})
		`, crudName)

	return data
}

// ============================================================================ CONTROLLER
func controller_Data(crudName, project string) string {
	data := fmt.Sprintf(
		`package %[1]v_c

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	models "github.com/DevopsGuyXD/%[2]v/Models"
	utils "github.com/DevopsGuyXD/%[2]v/Utils"
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
	w.WriteHeader(http.StatusOK)

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
	w.WriteHeader(http.StatusOK)

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	utils.Check_For_Nil(err)
	response := models.GET_%[1]v_by_id(id)

	if response != nil {
		json.NewEncoder(w).Encode(response)
	} else {
		response := fmt.Sprintf("No record with ID: %%d", id)
		json.NewEncoder(w).Encode(response)
	}
}

// -------------------------- %[1]v CREATE
// @Description Create a %[1]v
// @Tags %[1]v
// @Accept json
// @Produce json
// @Param data body object true "Add new record"
// @Success 200 {array} map[string]interface{}
// @Router /%[1]v [post]
func POST_%[1]v(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := r.Body
	response := models.CREATE_%[1]v(request)

	if response == "Created successfully" {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	} else{
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
	} 
}

// -------------------------- %[1]v UPDATE BY ID
// @Description Update %[1]v
// @Tags %[1]v
// @Accept json
// @Produce json
// @Param data body object true "Update record"
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

// ============================================================================ MODELS
func model_Data(crudName, projectName string) string {
	data := fmt.Sprintf(
		`package models

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	utils "github.com/DevopsGuyXD/%[2]v/Utils"
)

// -------------------------- %[1]v STRUCT
type %[3]v struct {
	Title    string `+"`json:\"title\"`"+`
	Author   string `+"`json:\"author\"`"+`
	Language string `+"`json:\"language\"`"+`
	Pages    int    `+"`json:\"pages\"`"+`
}

// -------------------------- CREATE %[1]v TABLE
func %[1]v() {

	db := initDB()
	defer db.Close()

	// Creating %[1]v table with PRIMARY KEY and DATE TIME
	table := "CREATE TABLE IF NOT EXISTS %[1]v(id INTEGER PRIMARY KEY AUTOINCREMENT, created_at DATETIME DEFAULT CURRENT_TIMESTAMP)"

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
		stmt := "ALTER TABLE %[1]v ADD COLUMN " + col + ";"
		_, err := db.Exec(stmt)
		if err != nil {
			if strings.Contains(err.Error(), "SQL logic error: duplicate column name") {
				continue
			}
			utils.Check_For_Nil(err)
		}
	}
}

// -------------------------- GET %[1]v ALL
func GET_%[1]v_all() []map[string]interface{} {
	switch {
	case os.Getenv("TEST_MODE") == "Y":
		// return getAll()
	default:
		query := "SELECT * FROM %[1]v"
		return get_Handler(query)
	}
}

// -------------------------- GET %[1]v by ID
func GET_%[1]v_by_id(id int) []map[string]interface{} {
	switch {
	case os.Getenv("TEST_MODE") == "Y":
		// return byId(id)
	default:
		query := fmt.Sprintf("SELECT * FROM %[1]v WHERE id = %%d", id)
		return get_Handler(query)
	}
}

// -------------------------- CREATE %[1]v RECORD
func CREATE_%[1]v(request io.ReadCloser) string {

	var data %[3]v

	switch {
	case os.Getenv("TEST_MODE") == "Y":
		if err := json.NewDecoder(request).Decode(&data); err != nil {
			return "Invalid JSON"
		} else {
			return "Created successfully"
		}
		
	default:
		err := json.NewDecoder(request).Decode(&data)
		utils.Check_For_Nil(err)

		return create_Handler(data)
	}
}

// -------------------------- UPDATE %[1]v
func UPDATE_%[1]v(id string, request io.ReadCloser) string {
	return update_Handler(id, request)
}

// -------------------------- DELETE %[1]v
func DELETE_%[1]v(id string) string {
	return delete_Handler(id)
}`, crudName, projectName, utils.Capitalize(crudName))

	return data
}

// ============================================================================ MODELS HANDLERS
func model_Handler_Data(crudName string) string {
	data := fmt.Sprintf(`
// -------------------------- INIT DB
func initDB() *sql.DB {
	var database = config.InitDatabase()
	return database
}

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
func create_Handler(data Books) string {

	db := initDB()
	defer db.Close()

	var cols, vals []string
	args := []interface{}{}

	for i := 0; i < reflect.TypeOf(data).NumField(); i++ {
		field := reflect.TypeOf(data).Field(i)
		val := reflect.ValueOf(data).Field(i)

		tag := field.Tag.Get("json")
		if tag != "" {
			cols = append(cols, tag)
			vals = append(vals, "?")
			args = append(args, val.Interface())
		}
	}

	query := fmt.Sprintf("INSERT INTO %[1]v (%%s) VALUES (%%s)",
		strings.Join(cols, ", "), strings.Join(vals, ", "))

	if res, err := db.Exec(query, args...); err != nil {
		return "DB error"
	} else if rows, _ := res.RowsAffected(); rows > 0 {
		return "Created successfully"
	}

	return "No records created"
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
}`, crudName)

	return data
}

// ============================================================================ TEST
func test_Data(crudName, projectName string) string {
	data := fmt.Sprintf(
		`package app_tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	models "github.com/DevopsGuyXD/%[2]v/Models"
	routes "github.com/DevopsGuyXD/%[2]v/Routes"
	utils "github.com/DevopsGuyXD/%[2]v/Utils"
	"github.com/stretchr/testify/assert"
)

// -------------------------- GET TESTS COMMON
func test_cases(rr *httptest.ResponseRecorder, t *testing.T, opertaion string, allRecords bool) {
	switch {
	case opertaion == "GET":
		var book []models.Books
		err := json.Unmarshal(rr.Body.Bytes(), &book)
		assert.NoError(t, err, "Expected valid JSON object")

		assert.Contains(t, rr.Header().Get("Content-Type"), "application/json")

		assert.Equal(t, http.StatusOK, rr.Code)

		switch {
		case allRecords:
			assert.GreaterOrEqual(t, len(book), 2, "Expected more than one book record")
		default:
			assert.Equal(t, len(book), 1, "Expected book with ID 1")
		}

	case opertaion == "POST":
		assert.Equal(t, http.StatusCreated, rr.Code)
	}
}

// -------------------------- POST TEST COMMON

// -------------------------- TEST INIT
func setup(opertaion, route string, payload []byte) *httptest.ResponseRecorder {
	os.Setenv("TEST_MODE", "Y")

	rr := httptest.NewRecorder()
	router := routes.RouteCollection()

	req, err := http.NewRequest(opertaion, route, bytes.NewBuffer(payload))
	utils.Check_For_Nil(err)
	router.ServeHTTP(rr, req)

	return rr
}

// -------------------------- POST /%[1]v
func Test_%[1]v_POST(t *testing.T) {
	opertaion := "POST"
	route := "/%[1]v"
	allRecords := false
	newBook := models.Books{Title: "New Book", Author: "New Author", Language: "English", Pages: 200}
	payload, _ := json.Marshal(newBook)

	rr := setup(opertaion, route, payload)
	test_cases(rr, t, opertaion, allRecords)
}

// -------------------------- GET /%[1]v
func Test_%[1]v_GET(t *testing.T) {
	opertaion := "GET"
	route := "/%[1]v"
	allRecords := true

	rr := setup(opertaion, route, nil)
	test_cases(rr, t, opertaion, allRecords)
}

// -------------------------- GET /%[1]v/{id}
func Test_%[1]v_GET_ID(t *testing.T) {
	opertaion := "GET"
	route := "/%[1]v/1"
	allRecords := false

	rr := setup(opertaion, route, nil)
	test_cases(rr, t, opertaion, allRecords)
}

// // -------------------------- PUT /%[1]v/{id}
// func Test_%[1]v_PUT(t *testing.T) {

// 	os.Setenv("TEST_MODE", "Y")
// 	rr := httptest.NewRecorder()
// 	router := routes.RouteCollection()

// 	updatedBook := Book{
// 		Title:    "Updated Book",
// 		Author:   "Updated Author",
// 		Language: "Hindi",
// 		Pages:    300,
// 	}
// 	payload, _ := json.Marshal(updatedBook)

// 	req, err := http.NewRequest("PUT", "/%[1]v/1", bytes.NewBuffer(payload))
// 	req.Header.Set("Content-Type", "application/json")
// 	utils.Check_For_Nil(err)
// 	router.ServeHTTP(rr, req)

// 	assert.Equal(t, http.StatusOK, rr.Code)
// }

// // -------------------------- DELETE /%[1]v/{id}
// func Test_%[1]v_DELETE(t *testing.T) {

// 	os.Setenv("TEST_MODE", "Y")
// 	rr := httptest.NewRecorder()
// 	router := routes.RouteCollection()

// 	req, err := http.NewRequest("DELETE", "/%[1]v/1", nil)
// 	utils.Check_For_Nil(err)
// 	router.ServeHTTP(rr, req)

// 	assert.Equal(t, http.StatusOK, rr.Code)
// }

// // -------------------------- Error Case: GET Nonexistent
// func Test_%[1]v_GET_NotFound(t *testing.T) {

// 	os.Setenv("TEST_MODE", "Y")
// 	rr := httptest.NewRecorder()
// 	router := routes.RouteCollection()

// 	req, err := http.NewRequest("GET", "/%[1]v/9999", nil)
// 	utils.Check_For_Nil(err)
// 	router.ServeHTTP(rr, req)

// 	assert.Equal(t, http.StatusNotFound, rr.Code)
// }

// // -------------------------- Error Case: DELETE Nonexistent
// func Test_%[1]v_DELETE_NotFound(t *testing.T) {

// 	os.Setenv("TEST_MODE", "Y")
// 	rr := httptest.NewRecorder()
// 	router := routes.RouteCollection()

// 	req, err := http.NewRequest("DELETE", "/%[1]v/9999", nil)
// 	utils.Check_For_Nil(err)
// 	router.ServeHTTP(rr, req)

// 	assert.Equal(t, http.StatusNotFound, rr.Code)
// }
`, crudName, projectName)

	return data
}
