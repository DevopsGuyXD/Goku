package templates_curd

import (
	"fmt"
	"strings"

	utils "github.com/DevopsGuyXD/Goku/Utils"
)

// ============================================================================ CRUD TEMPLATE
func CRUD_Project(crudName string) {
	route(crudName)
	controller(crudName)
	sqlite()
	model(crudName)
	model_Handlers(crudName)
	model_Imports()
	utils.Install_Dependencies()

	// -------------------- DONE STATUS
	fmt.Printf("\rAdding \"%v\" ✔\n\n", crudName)
}

// ============================================================================ CRUD ADD ROUTE
func route(crudName string) {
	filePath := "./Routes/routes.go"
	imports := []string{
		`books_c "github.com/DevopsGuyXD/myapp/Controller/books"`,
	}

	file := utils.Open_File(filePath)
	defer file.Close()

	lines := utils.InsertIntoFileBefore(routes_Data(crudName), file)

	utils.Write_File(file, strings.Join(lines, "\n"))
	utils.UpdateImport(filePath, imports)
}

// ============================================================================ CRUD ADD CONTROLLER
func controller(crudName string) {
	folder := "Controller/" + crudName
	newFile := []string{folder + "/" + crudName + ".go"}

	if !utils.Folder_Exists(folder) {
		utils.Create_Folder([]string{folder})
		utils.Create_File(newFile)

		file := utils.Open_File(newFile[0])
		defer file.Close()

		utils.Write_File(file, controller_Data(crudName, utils.Project_Name()))
	}
}

// ============================================================================ CRUD SQLITE
func sqlite() {
	databaseFolder := "Sqlite"

	if !utils.Folder_Exists(databaseFolder) {
		utils.Create_Folder([]string{databaseFolder})

		utils.Create_File([]string{"./" + databaseFolder + "/app.db"})
	}
}

// ============================================================================ CRUD ADD MODEL
func model(crudName string) {
	folder := "Models"
	newFile := []string{folder + "/" + crudName + ".go"}

	utils.Create_File(newFile)

	file := utils.Open_File(newFile[0])
	defer file.Close()

	utils.Write_File(file, model_Data(crudName, utils.Project_Name()))
	utils.UpdateAppConfig(crudName)
}

// ============================================================================ CRUD ADD MODEL HANDLERS
func model_Handlers(crudName string) {

	filePath := "./Models/models.go"

	targets := map[string]bool{
		"// -------------------------- GET HANDLER":    false,
		"// -------------------------- CREATE HANDLER": false,
		"// -------------------------- UPDATE HANDLER": false,
		"// -------------------------- DELETE HANDLER": false,
	}

	if !utils.Check_If_Lines_Exist(filePath, targets) {
		modelFile := utils.Open_File(filePath)
		defer modelFile.Close()

		utils.AppendToFileBottom(filePath, model_Handler_Data(crudName))
	}
}

// ============================================================================ UPDATE WITH REQUIRED IMPORTS
func model_Imports() {

	filePath := "./Models/models.go"
	imports := []string{
		`"database/sql"`,
		`"encoding/json"`,
		`"fmt"`,
		`"io"`,
		`"strings"`,

		`config "github.com/DevopsGuyXD/myapp/Config"`,
		`utils "github.com/DevopsGuyXD/myapp/Utils"`,
	}

	utils.UpdateImport(filePath, imports)
}
