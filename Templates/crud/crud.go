package templates_curd

import (
	"fmt"
	"strings"

	utils "github.com/DevopsGuyXD/Goku/Utils"
)

// ============================================================================ CRUD TEMPLATE
func CRUD_Project(crudName string) {

	// filePath := "./Models/models.go"

	crud_Route(crudName)
	crud_Controller(crudName)
	crud_Sqlite()
	crud_Model(crudName)
	crud_Model_Handlers(crudName)
	updating_Config_Main(crudName)
	update_Config_Main_Packages()
	// addingTest(crudName)

	// utils.Install_Dependencies()

	// -------------------- DONE STATUS
	fmt.Printf("\rAdding \"%v\" ✔\n\n", crudName)
}

// ============================================================================ CRUD ADD ROUTE
func crud_Route(crudName string) {

	filePath := "./Routes/routes.go"
	data := routes_Data(crudName)
	imports := []string{
		`books_c "github.com/DevopsGuyXD/myapp/Controller/books"`,
	}

	file := utils.Open_File(filePath)
	defer file.Close()

	lines := utils.InsertIntoFileBefore(data, file)

	utils.Write_File(file, strings.Join(lines, "\n"))

	utils.UpdateImport(filePath, imports)

	// =================================================
	// topLine := "import ("
	// data = "\"os\""

	// targets := map[string]bool{
	// 	"\t\"os\"": false,
	// }

	// if !utils.Check_If_Lines_Exist(filePath, targets) {
	// 	utils.InsertIntoFileAfter(topLine, filePath, data)
	// }
}

// ============================================================================ CRUD ADD CONTROLLER
func crud_Controller(crudName string) {

	data := controller_Data(crudName, utils.Project_Name())
	folder := "Controller/" + crudName
	newFile := []string{folder + "/" + crudName + ".go"}

	if !utils.Folder_Exists(folder) {
		utils.Create_Folder([]string{folder})
		utils.Create_File(newFile)

		file := utils.Open_File(newFile[0])
		defer file.Close()

		utils.Write_File(file, data)
	}

	// =================================================
	// file := utils.Open_File("./Routes/routes.go")
	// defer file.Close()

	// utils.Write_File(file, strings.Join(lines, "\n"))

	// _, err := file.WriteString(data)
	// utils.Check_For_Nil(err)

	// utils.InsertIntoFileAfter("import (", "./Routes/routes.go", fmt.Sprintf("%[1]v_c \"github.com/DevopsGuyXD/myapp/Controller/%[1]v\"", crudName))

	// filePath = fmt.Sprintf("./Controller/%[1]v/%[1]v_test.go", crudName)
}

// ============================================================================ CRUD SQLITE
func crud_Sqlite() {
	databaseFolder := "Sqlite"

	if !utils.Folder_Exists(databaseFolder) {
		utils.Create_Folder([]string{databaseFolder})

		utils.Create_File([]string{"./" + databaseFolder + "/app.db"})
	}
}

// ============================================================================ CRUD ADD MODEL
func crud_Model(crudName string) {

	data := model_Data(crudName, utils.Project_Name())
	folder := "Models"
	newFile := []string{folder + "/" + crudName + ".go"}

	utils.Create_File(newFile)

	file := utils.Open_File(newFile[0])
	defer file.Close()

	utils.Write_File(file, data)

}

// ============================================================================ UPDATING CONFIG MAIN
func updating_Config_Main(crudName string) {

	topLine := "func AppModels(){"
	filePath := "./Models/models.go"
	data := fmt.Sprintf("%v()", crudName)

	utils.InsertIntoFileAfter(topLine, filePath, data)
}

// ============================================================================ CRUD ADD MODEL HANDLERS
func crud_Model_Handlers(crudName string) {

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

		data := model_Handler_Data(crudName)

		utils.AppendToFileBottom(filePath, data)
	}

	// update_Config_Main_Packages(filePath)
}

// ============================================================================ UPDATE PACKAGES
func update_Config_Main_Packages() {

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

	//===========================================================
	// 	packages := `
	// import (
	// 	"database/sql"
	// 	"encoding/json"
	// 	"fmt"
	// 	"io"
	// 	"strings"

	// 	config "github.com/DevopsGuyXD/myapp/Config"
	// 	utils "github.com/DevopsGuyXD/myapp/Utils"
	// )`

	// 	data, err := os.ReadFile(filePath)
	// 	utils.Check_For_Nil(err)

	// 	content := string(data)
	// 	lines := strings.Split(content, "\n")

	// 	for i, line := range lines {
	// 		if line == "package models" {
	// 			lines = append(lines[:i+1], append([]string{packages}, lines[i+1:]...)...)
	// 			break
	// 		}
	// 	}

	// 	err = os.WriteFile(filePath, []byte(strings.Join(lines, "\n")), 0644)
	// 	utils.Check_For_Nil(err)
}

// ============================================================================
// func addingTest(crudName string) {

// 	folder := "./Test"
// 	utils.Create_Folder(folder)

// 	utils.Create_Folder(fmt.Sprintf("%v/%[1]v/%[1]v_test.go", folder, crudName))

// 	// testData := Test(crudName)

// 	//   = utils.Create_File(filePath)
// 	// defer file.Close()

// }
