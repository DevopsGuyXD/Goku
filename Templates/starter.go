package templates

import (
	"fmt"

	utils "github.com/DevopsGuyXD/Goku/Utils"
)

// ============================================================================ SIMPLE TEMPLATE
func Starter_Project(project string) {

	// -------------------- FOLDER LIST
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

	// -------------------- CREATE FOLDERS
	utils.Create_Folder(folders)

	// -------------------- CREATE FILES
	for _, folder := range folders {
		file, data := fileController(project, folder)

		if file != nil || data != "" {
			utils.Write_File(file, data)
		}
	}

	// -------------------- DONE STATUS
	fmt.Printf("\n\rCreating %v ✔\n\n", project)
}
