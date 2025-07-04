package templates

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	utils "github.com/DevopsGuyXD/Goku/Utils"
)

// ====================================== CRUD TEMPLATE
func CRUD_Template(crudName string) {

	filePath := "./Models/models.go"

	crud_Route(crudName)
	crud_Controller(crudName)
	crud_Sqlite()
	crud_Model(crudName)
	updating_Config_Main(crudName)

	if !utils.Check_If_Lines_Exist(filePath) {
		crud_Model_Handlers(crudName)
	}

	utils.Install_Dependencies()

	fmt.Printf("\rAdding \"%v\" CRUD ✔\n\n", crudName)
}

// ====================================== CRUD ADD ROUTE
func crud_Route(crudName string) {

	var lines []string

	filePath := "./Routes/routes.go"
	data := routes(crudName)

	file := utils.Open_File(filePath)
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
	utils.Check_For_Nil(err)

	err = os.WriteFile(filePath, []byte(strings.Join(lines, "\n")), 0644)
	utils.Check_For_Nil(err)
}

// ====================================== CRUD ADD CONTROLLER
func crud_Controller(crudName string) {

	project := utils.Get_Project_Name()

	filePath := fmt.Sprintf("./Controller/%v.go", crudName)
	data := controllers(crudName, project)

	file := utils.Create_File(filePath)
	defer file.Close()

	_, err := file.WriteString(data)
	utils.Check_For_Nil(err)
}

// ====================================== CRUD SQLITE
func crud_Sqlite() {
	databaseFolder := "Sqlite"
	fileName := "./" + databaseFolder + "/app.db"

	if !utils.Folder_Exists(databaseFolder) {
		utils.Create_Single_Folder(databaseFolder)

		dbFile := utils.Create_File(fileName)
		defer dbFile.Close()
	}
}

// ====================================== CRUD ADD MODEL
func crud_Model(crudName string) {

	projectName := utils.Get_Project_Name()

	file := fmt.Sprintf("./Models/%v.go", crudName)

	modelFile := utils.Create_File(file)
	defer modelFile.Close()

	data := models(projectName, crudName)

	_, err := modelFile.WriteString(data)
	utils.Check_For_Nil(err)

}

// ====================================== UPDATING CONFIG MAIN
func updating_Config_Main(crudName string) {

	var lines []string
	filePath := "./Models/models.go"

	data := fmt.Sprintf("%v()", crudName)

	file := utils.Open_File(filePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)

		if strings.Contains(line, "func AppModels(){") {
			lines = append(lines, "\t"+data)
		}
	}

	utils.Check_For_Nil(scanner.Err())

	err := os.WriteFile(filePath, []byte(strings.Join(lines, "\n")), 0644)
	utils.Check_For_Nil(err)
}

// ====================================== CRUD ADD MODEL HANDLERS
func crud_Model_Handlers(crudName string) {

	filePath := "./Models/models.go"

	modelFile := utils.Open_File(filePath)
	defer modelFile.Close()

	data := modelHandlers(crudName)

	utils.AppendToLastLine(filePath, data)

	update_Config_Main_Packages(filePath)
}

// ====================================== UPDATE PACKAGES
func update_Config_Main_Packages(filePath string) {

	packages := `
import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	config "github.com/DevopsGuyXD/myapp/Config"
	utils "github.com/DevopsGuyXD/myapp/Utils"
)`

	data, err := os.ReadFile(filePath)
	utils.Check_For_Nil(err)

	content := string(data)
	lines := strings.Split(content, "\n")

	for i, line := range lines {
		if line == "package models" {
			lines = append(lines[:i+1], append([]string{packages}, lines[i+1:]...)...)
			break
		}
	}

	err = os.WriteFile(filePath, []byte(strings.Join(lines, "\n")), 0644)
	utils.Check_For_Nil(err)
}
