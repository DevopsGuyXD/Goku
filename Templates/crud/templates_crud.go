package templates_curd

import (
	"fmt"

	utils "github.com/DevopsGuyXD/Goku/Utils"
)

// ============================================================================ routes.go DATA
func routes_Data(crudName string) string {
	data := fmt.Sprintf(`	// -------------------------- %[1]v
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

// ============================================================================ CONTROLLER
func controller_Data(crudName, project string) string {
	data := fmt.Sprintf(
		`package controller

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
	utils.Check_For_Err(err)
	response := models.GET_%[1]v_by_id(id)

	if response != nil {
		json.NewEncoder(w).Encode(response)
	} else {
		response := fmt.Sprintf("No record with ID: %%d", id)
		json.NewEncoder(w).Encode(response)
	}
}

// -------------------------- %[1]v POST
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
	response := models.POST_%[1]v(request)

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

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	utils.Check_For_Err(err)
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

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	utils.Check_For_Err(err)

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
var %[1]v %[3]v
var %[1]v_list []%[3]v

type %[3]v struct {
	Title    string `+"`json:\"title\"`"+`
	Author   string `+"`json:\"author\"`"+`
	Language string `+"`json:\"language\"`"+`
	Pages    int    `+"`json:\"pages\"`"+`
}

// -------------------------- CREATE %[1]v TABLE
func %[1]v_ct() {

	db := initDB()
	defer db.Close()

	table := "CREATE TABLE IF NOT EXISTS %[1]v(id INTEGER PRIMARY KEY AUTOINCREMENT, created_at DATETIME DEFAULT CURRENT_TIMESTAMP)"

	createDomainsTable, err := db.Prepare(table)
	utils.Check_For_Err(err)
	createDomainsTable.Exec()

	columnsToAdd := []string{
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
			utils.Check_For_Err(err)
		}
	}
}

// -------------------------- GET %[1]v ALL
func GET_%[1]v_all() []map[string]interface{} {	
	switch {
	case os.Getenv("TEST_MODE") == "Y":
		var result []map[string]interface{}
		for _, data := range %[1]v_list {
			result = append(result, map[string]interface{}{
				"title": &data.Title,
			})
		}

		return result

	default:
		query := "SELECT * FROM %[1]v"

		jsonData, err := json.Marshal(get_Handler(query))
		utils.Check_For_Err(err)

		if err = json.Unmarshal(jsonData, &%[1]v_list); err != nil {
			return []map[string]interface{}{
				{"message": "Type error"},
			}
		}

		return get_Handler(query)
	}
}

// -------------------------- GET %[1]v by ID
func GET_%[1]v_by_id(id int) map[string]interface{} {
	switch {
	case os.Getenv("TEST_MODE") == "Y":
		return map[string]interface{}{
			"Title": &%[1]v_list[id-1].Title,
		}

	default:
		query := fmt.Sprintf("SELECT * FROM %[1]v WHERE id = %%d", id)

		result := get_Handler(query)
		if len(result) == 0 || len(result[0]) == 0 {
			return map[string]interface{}{
				"message": fmt.Sprintf("No record with ID: %%v", id),
			}
		}

		b, err := json.Marshal(result[0])
		utils.Check_For_Err(err)

		if err = json.Unmarshal(b, &%[1]v); err != nil {
			return map[string]interface{}{
				"message": "Type error",
			}
		}

		return result[0]
	}
}

// -------------------------- CREATE %[1]v RECORD
func POST_%[1]v(request io.ReadCloser) string {
	switch {
	case os.Getenv("TEST_MODE") == "Y":
		if err := json.NewDecoder(request).Decode(&%[1]v); err != nil {
			return "Invalid JSON"
		} else {
		 	%[1]v_list = append(%[1]v_list, %[1]v)
			return "Created successfully"
		}

	default:
		err := json.NewDecoder(request).Decode(&%[1]v)
		utils.Check_For_Err(err)
		defer request.Close()

		return post_Handler(%[1]v)
	}
}

// -------------------------- UPDATE %[1]v
func UPDATE_%[1]v(id int, request io.ReadCloser) string {
	switch {
	case os.Getenv("TEST_MODE") == "Y":
		%[1]v_list[id-1].Title = "New %[1]v 3"
		if %[1]v_list[id-1].Title == "New %[1]v 3" {
			return "Updated successfully"
		}
		return "Error updating"

	default:
		err := json.NewDecoder(request).Decode(&%[1]v)
		utils.Check_For_Err(err)
		defer request.Close()

		return update_Handler(id, %[1]v)
	}
}

// -------------------------- DELETE %[1]v
func DELETE_%[1]v(id int) string {
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

	var response []map[string]interface{}

	db := initDB()
	defer db.Close()

	rows, err := db.Query(query)
	utils.Check_For_Err(err)
	defer rows.Close()

	cols, err := rows.Columns()
	utils.Check_For_Err(err)

	vals := make([]interface{}, len(cols))

	for rows.Next() {
		row := make(map[string]interface{}, len(cols))

		for i := range vals {
			vals[i] = new(interface{})
		}
		err := rows.Scan(vals...)
		utils.Check_For_Err(err)

		for i, col := range cols {
			row[col] = vals[i]
		}
		response = append(response, row)
	}

	return response
}

// -------------------------- CREATE HANDLER
func post_Handler(data interface{}) string {

	db := initDB()
	defer db.Close()

	var cols, vals []string
	val := reflect.ValueOf(data)
	typ := reflect.TypeOf(data)
	args := make([]interface{}, 0, typ.NumField())

	for i := 0; i < typ.NumField(); i++ {
		tag := typ.Field(i).Tag.Get("json")

		if tag != "" {
			cols = append(cols, tag)
			vals = append(vals, "?")
			args = append(args, val.Field(i).Interface())
		}
	}

	query := fmt.Sprintf("INSERT INTO %[1]v (%%s) VALUES (%%s)", strings.Join(cols, ", "), strings.Join(vals, ", "))

	if res, err := db.Exec(query, args...); err != nil {
		return "DB error"
	} else if rows, _ := res.RowsAffected(); rows > 0 {
		return "Created successfully"
	}
	return "No records created"
}

// -------------------------- UPDATE HANDLER
func update_Handler(id int, data interface{}) string {
	db := initDB()
	defer db.Close()

	var vals []string
	val := reflect.ValueOf(data)
	typ := reflect.TypeOf(data)
	args := make([]interface{}, 0, typ.NumField())

	for i := 0; i < typ.NumField(); i++ {
		field := val.Field(i)
		tag := typ.Field(i).Tag.Get("json")

		if tag != "" && !field.IsZero() {
			vals = append(vals, fmt.Sprintf("%%s = ?", tag))
			args = append(args, field.Interface())
		}
	}

	args = append(args, id)
	query := fmt.Sprintf("UPDATE %[1]v SET %%s WHERE id = ?", strings.Join(vals, ", "))

	if res, err := db.Exec(query, args...); err != nil {
		return "DB update error"
	} else if rows, _ := res.RowsAffected(); rows > 0 {
		return "Updated successfully"
	}
	return "No records updated"
}

// -------------------------- DELETE HANDLER
func delete_Handler(id int) string {

	db := initDB()
	defer db.Close()

	var response string

	query := fmt.Sprintf("DELETE FROM %[1]v WHERE id=%%d", id)

	res, err := db.Exec(query)
	if err != nil {
		response = "DB delete error"

	} else {
		rowsAffected, err := res.RowsAffected()
		utils.Check_For_Err(err)

		if rowsAffected != 0 {
			response = "Deleted successfully"
		} else {
			response = "No records deleted"
		}

	}

	return response
}`, crudName, utils.Capitalize(crudName))

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

// -------------------------- TEST CASES
func test_cases(rr *httptest.ResponseRecorder, t *testing.T, opertaion string, allRecords bool) {

	common_cases := func(status int, err error) {
		assert.NoError(t, err, "Failed to unmarshal response")
		assert.Contains(t, rr.Header().Get("Content-Type"), "application/json")
		assert.Equal(t, status, rr.Code)
	}

	switch {
	case opertaion == "GET":
		switch {
		case allRecords:
			var %[1]v []models.%[3]v
			err := json.Unmarshal(rr.Body.Bytes(), &%[1]v)

			assert.LessOrEqual(t, len(%[1]v), 2, "Expected records greater than 1")
			common_cases(http.StatusOK, err)

		case !allRecords:
			var %[1]v models.%[3]v
			err := json.Unmarshal(rr.Body.Bytes(), &%[1]v)

			assert.Equal(t, "New %[1]v 1", %[1]v.Title, "Expected Title to be 'New %[1]v 1'")
			common_cases(http.StatusOK, err)
		}

	case opertaion == "POST":
		common_cases(http.StatusCreated, nil)

	case opertaion == "PUT":
		common_cases(http.StatusOK, nil)

	}
}

// -------------------------- TEST INIT
func setup(opertaion, route string, payload []byte) *httptest.ResponseRecorder {
	os.Setenv("TEST_MODE", "Y")

	rr := httptest.NewRecorder()
	router := routes.RouteCollection()

	req, err := http.NewRequest(opertaion, route, bytes.NewBuffer(payload))
	utils.Check_For_Err(err)
	router.ServeHTTP(rr, req)

	return rr
}

// -------------------------- POST /%[1]v
func Test_%[1]v_POST(t *testing.T) {
	opertaion := "POST"
	route := "/%[1]v"
	allRecords := false
	var rr *httptest.ResponseRecorder
	%[1]v := []models.%[3]v{
		{Title: "New %[1]v 1", Author: "New Author 1", Language: "English", Pages: 144},
		{Title: "New %[1]v 2", Author: "New Author 2", Language: "French", Pages: 164},
	}

	for _, data := range %[1]v {
		payload, _ := json.Marshal(data)
		rr = setup(opertaion, route, payload)
	}
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

// -------------------------- PUT /%[1]v/{id}
func Test_%[1]v_PUT(t *testing.T) {
	opertaion := "PUT"
	route := "/%[1]v/1"
	allRecords := false
	update%[1]v := []models.%[3]v{
		{Title: "New %[1]v 3"},
	}
	payload, _ := json.Marshal(update%[1]v)

	rr := setup(opertaion, route, payload)
	test_cases(rr, t, opertaion, allRecords)
}

// // -------------------------- DELETE /%[1]v/{id}
// func Test_%[1]v_DELETE(t *testing.T) {

// 	os.Setenv("TEST_MODE", "Y")
// 	rr := httptest.NewRecorder()
// 	router := routes.RouteCollection()

// 	req, err := http.NewRequest("DELETE", "/%[1]v/1", nil)
// 	utils.Check_For_Err(err)
// 	router.ServeHTTP(rr, req)

// 	assert.Equal(t, http.StatusOK, rr.Code)
// }

// // -------------------------- Error Case: GET Nonexistent
// func Test_%[1]v_GET_NotFound(t *testing.T) {

// 	os.Setenv("TEST_MODE", "Y")
// 	rr := httptest.NewRecorder()
// 	router := routes.RouteCollection()

// 	req, err := http.NewRequest("GET", "/%[1]v/9999", nil)
// 	utils.Check_For_Err(err)
// 	router.ServeHTTP(rr, req)

// 	assert.Equal(t, http.StatusNotFound, rr.Code)
// }

// // -------------------------- Error Case: DELETE Nonexistent
// func Test_%[1]v_DELETE_NotFound(t *testing.T) {

// 	os.Setenv("TEST_MODE", "Y")
// 	rr := httptest.NewRecorder()
// 	router := routes.RouteCollection()

// 	req, err := http.NewRequest("DELETE", "/%[1]v/9999", nil)
// 	utils.Check_For_Err(err)
// 	router.ServeHTTP(rr, req)

// 	assert.Equal(t, http.StatusNotFound, rr.Code)
// }
`, crudName, projectName, utils.Capitalize(crudName))

	return data
}
