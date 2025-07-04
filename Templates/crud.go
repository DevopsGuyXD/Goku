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
	crudSqlite()
	crudModel(crudName)
	crudModelHandlers(crudName)
	updatingConfigMain(crudName)
	utils.InstallDependencies()
	fmt.Printf("\rAdding \"%v\" CRUD \n", crudName)
}

// ====================================== CRUD ADD ROUTE
func crudRoute(crudName string) {

	var lines []string

	filePath := "./Routes/routes.go"
	data := routes(crudName)

	file := utils.OpenFile(filePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "return router") {
			lines = append(lines, data)
		}
		lines = append(lines, line)
	}

	err := scanner.Err()
	utils.CheckForNil(err)

	err = os.WriteFile(filePath, []byte(strings.Join(lines, "\n")), 0644)
	utils.CheckForNil(err)
}

// ====================================== CRUD ADD CONTROLLER
func crudController(crudName string) {

	project := utils.GetProjectName()

	filePath := fmt.Sprintf("./Controller/%v.go", crudName)
	data := controllers(crudName, project)

	file := utils.CreateFile(filePath)
	defer file.Close()

	_, err := file.WriteString(data)
	utils.CheckForNil(err)
}

// ====================================== CRUD SQLITE
func crudSqlite() {
	databaseFolder := "Sqlite"
	fileName := "./" + databaseFolder + "/app.db"

	if !utils.FolderExists(databaseFolder) {
		utils.CreateSingleFolder(databaseFolder)

		dbFile := utils.CreateFile(fileName)
		defer dbFile.Close()
	}
}

// ====================================== CRUD ADD MODEL
func crudModel(crudName string) {

	projectName := utils.GetProjectName()

	file := fmt.Sprintf("./Models/%v.go", crudName)

	modelFile := utils.CreateFile(file)
	defer modelFile.Close()

	data := models(projectName, crudName)

	_, err := modelFile.WriteString(data)
	utils.CheckForNil(err)

}

// ====================================== CRUD ADD MODEL HANDLERS
func crudModelHandlers(crudName string) {

	file := "./Models/models.go"

	modelFile := utils.OpenFile(file)
	defer modelFile.Close()

	data := modelHandlers(crudName)

	utils.AppendToLastLine(file, data)
}

// ====================================== UPDATING CONFIG MAIN
func updatingConfigMain(crudName string) {

	var lines []string
	filePath := "./Models/models.go"

	data := fmt.Sprintf("%v()", crudName)

	file := utils.OpenFile(filePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)

		if strings.Contains(line, "func AppModels(){") {
			lines = append(lines, "\t"+data)
		}
	}

	utils.CheckForNil(scanner.Err())

	err := os.WriteFile(filePath, []byte(strings.Join(lines, "\n")), 0644)
	utils.CheckForNil(err)

	utils.UpdatePackages(filePath, tester())
}
