package utils

import (
	"bufio"
	"log"
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

// ============================================================================ CREATE FILE
func Create_File(files []string) {
	for _, file := range files {
		_, err := os.Create(file)
		Check_For_Nil(err)
	}
}

// ============================================================================ OPEN FILE
func Open_File(filePath string) *os.File {
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	Check_For_Nil(err)

	return file
}

// ============================================================================ WRITE TO FILE
func Write_File(file *os.File, data string) {
	_, err := file.Seek(0, 0)
	Check_For_Nil(err)

	_, err = file.WriteString(data)
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
	Check_For_Nil(err)

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

	Check_For_Nil(scanner.Err())

	err := os.WriteFile(filePath, []byte(strings.Join(lines, "\n")), 0644)
	Check_For_Nil(err)
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
func UpdateImport() {

}
