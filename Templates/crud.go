package templates

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	utils "github.com/DevopsGuyXD/Goku/Utils"
)

// ====================================== CRUD TEMPLATE
func CRUDTemplate(crudName string) {
	crudRoute(crudName)
	crudController(crudName)
	// crudConfig()
	crudModel(crudName)
}

// ====================================== CRUD ADD ROUTE
func crudRoute(crudName string) {

	var lines []string
	filePath := "./Routes/routes.go"

	data := fmt.Sprintf(
		`	//-------------------------- %[1]v
	router.Route("/%[1]v", func(r chi.Router) {
		r.Get("/", controller.GET_%[1]v)
		r.Post("/", controller.POST_%[1]v)
		r.Get("/{id}", controller.GET_%[1]v_id)
		r.Put("/{id}", controller.UPDATE_%[1]v)
		r.Delete("/{id}", controller.DELETE_%[1]v)
	})
		`, crudName)

	file, err := os.Open(filePath)
	utils.CheckForNil(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "return router") {
			lines = append(lines, data)
		}
		lines = append(lines, line)
	}

	err = scanner.Err()
	utils.CheckForNil(err)

	err = os.WriteFile(filePath, []byte(strings.Join(lines, "\n")), 0644)
	utils.CheckForNil(err)
}

// ====================================== CRUD ADD CONTROLLER
func crudController(crudName string) {

	project := utils.GetProjectName()

	data := fmt.Sprintf(

		`package controller

import (
	"encoding/json"
	"net/http"

	models "github.com/DevopsGuyXD/%[2]v/Models"
	"github.com/go-chi/chi/v5"
)

//-------------------------- %[1]v GET ALL
// @Description All %[1]v
// @Tags %[1]v
// @Produce json
// @Success 200 {array} map[string]interface{}
// @Router /%[1]v [get]
func GET_%[1]v(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	results := models.GET_%[1]v_all()
	
	if results != nil {
		json.NewEncoder(w).Encode(results)
	} else {
		json.NewEncoder(w).Encode("No records found")
	}
}

//-------------------------- %[1]v CREATE
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

//-------------------------- %[1]v GET BY ID
// @Description Single %[1]v
// @Tags %[1]v
// @Produce json
// @Success 200 {array} map[string]interface{}
// @Router /%[1]v/{id} [get]
func GET_%[1]v_id(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
		
	id := chi.URLParam(r, "id")
	response := models.GET_%[1]v_by_id(id)

	json.NewEncoder(w).Encode(response)
}

//-------------------------- %[1]v UPDATE BY ID
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

//-------------------------- %[1]v DELETE
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

	file, err := os.Create(fmt.Sprintf("./Controller/%v.go", crudName))
	utils.CheckForNil(err)
	defer file.Close()

	_, err = file.WriteString(data)
	utils.CheckForNil(err)
}

// ====================================== CRUD ADD MODEL
func crudModel(crudName string) {

	projectName := utils.GetProjectName()
	databaseFolder := "Sqlite"
	databaseFile := "./Sqlite/app.db"

	utils.CreateFolder(databaseFolder)
	utils.CreateFile(databaseFile)

	file, err := os.Create(fmt.Sprintf("./Models/%v.go", crudName))
	utils.CheckForNil(err)
	defer file.Close()

	data := fmt.Sprintf(
		`package models

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"


	config "github.com/DevopsGuyXD/%[1]v/Config"
	utils "github.com/DevopsGuyXD/%[1]v/Utils"
)

// -------------------------- CREATE %[2]v TABLE
func %[2]v() {

	database := config.InitDatabase()
	defer database.Close()

	table := "CREATE TABLE IF NOT EXISTS %[2]v(id INTEGER PRIMARY KEY AUTOINCREMENT, created_at DATETIME DEFAULT CURRENT_TIMESTAMP)"

	createDomainsTable, err := database.Prepare(table)
	utils.CheckForNil(err)
	createDomainsTable.Exec()

	columnsToAdd := []string{
		// ADD COLUMNS HERE. Example,
		// "title VARCHAR(25)",
		// "author VARCHAR(25)",
		// "language VARCHAR(25)",
		// "pages INTEGER",
	}

	for _, col := range columnsToAdd {
		stmt := "ALTER TABLE %[2]v ADD COLUMN " + col + ";"
		_, err := database.Exec(stmt)
		if err != nil {
			if strings.Contains(err.Error(), "SQL logic error: duplicate column name") {
				continue
			}
			utils.CheckForNil(err)
		}
	}
}

// // -------------------------- GET %[2]v ALL
func GET_%[2]v_all() []map[string]interface{} {

	var results []map[string]interface{}

	var database = config.InitDatabase()
	defer database.Close()

	data, err := database.Query("SELECT * FROM %[2]v")
	utils.CheckForNil(err)
	defer data.Close()

	columns, err := data.Columns()
	utils.CheckForNil(err)

	for data.Next() {

		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))
		rowMap := make(map[string]interface{})

		for i := range values {
			valuePtrs[i] = &values[i]
		}

		err := data.Scan(valuePtrs...)
		utils.CheckForNil(err)

		for i, col := range columns {
			val := values[i]

			if b, ok := val.([]byte); ok {
				rowMap[col] = string(b)
			} else {
				rowMap[col] = val
			}
		}

		results = append(results, rowMap)
	}

	return results
}

// // -------------------------- CREATE %[2]v RECORD
func CREATE_%[2]v(request io.ReadCloser) string {

	var response string
	var data map[string]interface{}

	var database = config.InitDatabase()
	defer database.Close()

	err := json.NewDecoder(request).Decode(&data)
	utils.CheckForNil(err)

	columns := []string{}
	placeholders := []string{}
	values := []interface{}{}

	for k, v := range data {
		columns = append(columns, k)
		placeholders = append(placeholders, "?")
		values = append(values, v)
	}

	query := fmt.Sprintf("INSERT INTO %[2]v (%%s) VALUES (%%s)",
		strings.Join(columns, ", "),
		strings.Join(placeholders, ", "))

	res, err := database.Exec(query, values...)
	if err != nil {
		response = "DB insert error"

	} else {
		rowsAffected, err := res.RowsAffected()
		utils.CheckForNil(err)

		if rowsAffected != 0 {
			response = "Created successfully"
		} else {
			response = "No records created"
		}
	}

	return response
}

// -------------------------- GET %[2]v by ID
func GET_%[2]v_by_id(id string) []map[string]interface{} {
	var results []map[string]interface{}

	var database = config.InitDatabase()
	defer database.Close()

	data, err := database.Query("SELECT * FROM %[2]v WHERE id = ?", id)
	utils.CheckForNil(err)
	defer data.Close()

	columns, err := data.Columns()
	utils.CheckForNil(err)

	for data.Next() {
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))
		rowMap := make(map[string]interface{})

		for i := range values {
			valuePtrs[i] = &values[i]
		}

		err := data.Scan(valuePtrs...)
		utils.CheckForNil(err)

		for i, col := range columns {
			val := values[i]

			switch v := val.(type) {
			case []byte:
				rowMap[col] = string(v)
			default:
				rowMap[col] = v
			}
		}

		results = append(results, rowMap)
	}

	if err = data.Err(); err != nil {
		utils.CheckForNil(err)
	}

	return results
}

// -------------------------- UPDATE %[2]v
func UPDATE_%[2]v(id string, request io.ReadCloser) string {

	var response string
	var data map[string]interface{}

	setParts := []string{}
	setValues := []interface{}{}

	var database = config.InitDatabase()
	defer database.Close()

	err := json.NewDecoder(request).Decode(&data)
	utils.CheckForNil(err)

	for k, v := range data {
		setParts = append(setParts, fmt.Sprintf("%%s = ?", k))
		setValues = append(setValues, v)
	}

	if len(setParts) == 0 {
		return "No fields provided for update"
	}

	query := fmt.Sprintf("UPDATE %[2]v SET %%s WHERE id = ?",
		strings.Join(setParts, ", "))

	values := append(setValues, id)

	res, err := database.Exec(query, values...)
	if err != nil {
		response = "DB update error"
	} else {
		rowsAffected, err := res.RowsAffected()
		utils.CheckForNil(err)

		if rowsAffected != 0 {
			response = "Updated successfully"
		} else {
			response = "No records updated"
		}
	}

	return response
}

// -------------------------- DELETE %[2]v
func DELETE_%[2]v(id string) string {

	var response string

	var database = config.InitDatabase()
	defer database.Close()

	query := fmt.Sprintf("DELETE FROM %[2]v WHERE id=%%s", id)

	res, err := database.Exec(query)
	if err != nil {
		response = "DB delete error"

	} else {
		rowsAffected, err := res.RowsAffected()
		utils.CheckForNil(err)

		if rowsAffected != 0 {
			response = "Deleted successfully"
		} else {
			response = "No records deleted"
		}
		
	}

	return response
}
`, projectName, crudName)

	file, err = os.Create(fmt.Sprintf("./Models/%v.go", crudName))
	utils.CheckForNil(err)
	defer file.Close()

	_, err = file.WriteString(data)
	utils.CheckForNil(err)

	utils.UpdatingMain(crudName)

	fmt.Println()
	utils.InitSpinner("Adding CRUD")
	fmt.Printf("\rAdding \"%v\" CRUD ✅ \n", crudName)

	utils.InstallDependencies()
}
