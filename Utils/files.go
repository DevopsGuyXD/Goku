package utils

import (
	"os"
	"strings"
)

// ============================================================================ CREATE FOLDER
func Create_Folder(folders []string) {
	for _, folder := range folders {
		if !strings.Contains(folder, ".") && !Folder_Exists(folder) {
			err := os.Mkdir(folder, 0755)
			Check_For_Nil(err)
		}
	}
}

// ============================================================================ CREATE FOLDER
func Create_File(files []string) {
	for _, file := range files {
		_, err := os.Create(file)
		Check_For_Nil(err)
	}
}

// ============================================================================ OPEN FILE
func Open_File(filePath string) *os.File {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	Check_For_Nil(err)

	return file
}

// ============================================================================ WRITE TO FILE
func Write_File(file *os.File, data string) {
	_, err := file.WriteString(data)
	Check_For_Nil(err)
}

// ============================================================================ FOLDER EXISTS
func Folder_Exists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}

	if info.IsDir() {
		return true
	}

	return false
}

// ============================================================================ FOLDER EXISTS
func Files_Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	}

	return true
}
