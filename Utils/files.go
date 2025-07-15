package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/DevopsGuyXD/Goku/Templates/common"
)

// ============================================================================ CREATE FOLDER
func Create_Folder(folders []string) {
	for _, folder := range folders {
		if !strings.Contains(folder, ".") && !strings.Contains(folder, "dockerfile") && !Folder_Exists(folder) {
			err := os.Mkdir(folder, 0755)
			Check_For_Err(err)
		}
	}
}

// ============================================================================ CREATE FILE
func Create_File(files []string) {
	for _, file := range files {
		_, err := os.Create(file)
		Check_For_Err(err)
	}
}

// ============================================================================ OPEN FILE
func Open_File(filePath string) *os.File {
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	Check_For_Err(err)

	return file
}

// ============================================================================ READ FILE
func Read_File(filePath string) []byte {
	data, err := os.ReadFile(filePath)
	Check_For_Err(err)

	return data
}

// ============================================================================ WRITE TO FILE
func Write_File(file *os.File, data string) {
	_, err := file.Seek(0, 0)
	Check_For_Err(err)

	_, err = file.WriteString(data)
	Check_For_Err(err)
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

// ============================================================================ INSERT INTO FILE BEFORE
func InsertIntoFileBefore(data string, file *os.File) []string {
	var lines []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "return router") {
			lines = append(lines, data)
		}
		lines = append(lines, line)
	}

	err := scanner.Err()
	Check_For_Err(err)

	return lines
}

// ============================================================================ INSERT INTO FILE AFTER
func InsertIntoFileAfter(topLine, filePath, data string) {
	var lines []string

	file := Open_File(filePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)

		if strings.Contains(line, topLine) {
			lines = append(lines, "\t"+data)
		}
	}

	Check_For_Err(scanner.Err())

	err := os.WriteFile(filePath, []byte(strings.Join(lines, "\n")), 0644)
	Check_For_Err(err)
}

// ============================================================================ APPEND TO LAST LINE
func AppendToFileBottom(file, data string) {
	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if _, err := f.WriteString(data); err != nil {
		log.Fatal(err)
	}
}

// ============================================================================ UPDATE IMPORT
func UpdateImport(filePath string, imports []string) {

	var import_Block_Exists bool
	file := Read_File(filePath)
	lines := strings.Split(string(file), "\n")

	//--------------------------------------
	for _, line := range lines {
		if line == "import (" {
			import_Block_Exists = true
			break
		}
	}

	//--------------------------------------
	if !import_Block_Exists {
		InsertIntoFileAfter("package", filePath, common.Import_Data())
	}

	//--------------------------------------
	for _, importData := range imports {
		InsertIntoFileAfter("import (", filePath, importData)
	}
}

// ============================================================================ UPDATE IMPORT
func UpdateAppConfig(crudName string) {
	topLine := "func AppModels(){"
	filePath := "./Models/models.go"
	data := fmt.Sprintf("%v()", crudName)

	InsertIntoFileAfter(topLine, filePath, data)
}

// ============================================================================ RETURN LINE FROM FILE
func ReturnLineFromFile(data *os.File) string {

	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "PORT") {
			return line
		}
	}

	return ""
}
