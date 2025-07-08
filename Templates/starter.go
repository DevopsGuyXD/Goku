package templates

import (
	"fmt"

	utils "github.com/DevopsGuyXD/Goku/Utils"
)

// ============================================================================ SIMPLE TEMPLATE
func Starter_Project(project string) {

	// ==== FOLDER LIST
	folders := []string{
		project,
		project + "\\main.go",
		project + "\\.env",
		project + "\\go.mod",
		project + "\\Routes",
		project + "\\Controller",
		project + "\\Config",
		project + "\\Models",
		project + "\\Utils"}

	// ==== CREATE FOLDERS
	utils.Create_Folder(folders)

	// === CREATE MAIN FILES
	// file(project, "/main.go", main_Data(project))
	// file(project, "/.env", env_Data)
	// file(project, "/go.mod", mod_Data(project))

	// === CREATE OTHER FILES
	for _, folder := range folders {
		otherFiles(project, folder)
	}

	// === DONE STATUS
	fmt.Printf("\n\rCreating %v ✔\n\n", project)
}

// ============================================================================ MAIN FILES
func file(project, fileName, data string) {

	filePath := project + fileName

	utils.Create_File([]string{filePath})
	file := utils.Open_File(filePath)
	defer file.Close()

	utils.Write_File(file, data)
}

// ============================================================================ OTHER FILES
func otherFiles(project, folder string) {

	file, data := fileController(project, folder)

	if file != nil || data != "" {
		utils.Write_File(file, data)
	}
}
