package templates_starter

import (
	"fmt"
	"path/filepath"

	utils "github.com/DevopsGuyXD/Goku/Utils"
)

// ============================================================================ SIMPLE TEMPLATE
func Starter_Project(project string) {

	// -------------------- FOLDER LIST
	folders := []string{
		project,
		filepath.Join(project, "main.go"),
		filepath.Join(project, "go.mod"),
		filepath.Join(project, "dockerfile"),
		filepath.Join(project, "Routes"),
		filepath.Join(project, "Controller"),
		filepath.Join(project, "Config"),
		filepath.Join(project, "Models"),
		filepath.Join(project, "Utils"),
	}

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
	fmt.Println()
	fmt.Printf("\rCreating %v ✔\n", project)
}
