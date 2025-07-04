package utils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// ====================================== GOKU VERSION
func Version() {
	fmt.Println(
		`    ____    ___    _  __   _   _ 
   / ___|  / _ \  | |/ /  | | | |
  | |  _  | | | | | ' /   | | | |
  | |_| | | |_| | | . \   | |_| |
   \____|  \___/  |_|\_\   \___/  v1.0.0`)
}

// ====================================== CREATOR
func Creator() {
	fmt.Printf("\n%v\n%v\n", "With love ❤️", "Bharath Dundi 🤘")
}

// ====================================== ERROR HANDLING
func CheckForNil(err error) {
	if err != nil {
		log.Println(err)
		return
	}
}

// ====================================== INSTALL DEPENDENCIES
func InstallDependencies() {

	initSwagger()

	done := make(chan bool)
	go Spinner(done, "Installing Dependencies")

	cmd := exec.Command("sh", "-c", "go mod tidy")
	err := cmd.Run()
	if err != nil {
		fmt.Printf("\rInstalling Dependencies ❌\n")
		close(done)
		return
	}

	if !FilsExists(".air.toml") {
		initair()
	}

	close(done)
	fmt.Printf("\rInstalling Dependencies \n")
}

// ====================================== INIT SWAGGER
func initSwagger() {

	done := make(chan bool)
	go Spinner(done, "Updating Swagger")

	calledFrom := CalledFromLocation()

	cmd := exec.Command("sh", "-c", fmt.Sprintf("go run github.com/swaggo/swag/cmd/swag@v1.8.12 init --dir \"%s\"", calledFrom))
	err := cmd.Run()
	if err != nil {
		fmt.Printf("\rUpdating Swagger ❌\n")
		close(done)
		return
	}

	close(done)
	fmt.Print("\rUpdating Swagger\n")
}

// ====================================== INIT Air
func initair() {

	calledFrom := CalledFromLocation()

	_, err := exec.Command("sh", "-c", fmt.Sprintf("go run github.com/air-verse/air@latest init --dir \"%s\"", calledFrom)).Output()
	CheckForNil(err)
}

// ====================================== ERROR HANDLING
func AllOptions() {
	fmt.Printf(`  
  Options:

    -h | --help
    -v | --version
    -i | --install

    goku create-project mytestapp
    goku run dev
    goku run build
    goku run start
    goku add-crud <NAME>
    goku add-docker
    goku build-docker <NAME:TAG> -> Note: TAG will be "latest" if not specified
    goku docker <NAME>
`)
}

// ====================================== OPEN FILE
func OpenFile(filePath string) *os.File {

	file, err := os.Open(filePath)
	CheckForNil(err)

	return file
}

// ====================================== CREATE FOLDER
func CreateSingleFolder(folderName string) {
	err := os.Mkdir(folderName, 0755)
	CheckForNil(err)
}

// ====================================== CREATE FOLDER
func CreateFile(fileName string) *os.File {
	file, err := os.Create(fileName)
	CheckForNil(err)

	return file
}

// ====================================== WRITE TO FILE
func WriteFile(file *os.File, data string) {
	_, err := file.WriteString(data)
	CheckForNil(err)
}

// ====================================== GET PROJECT NAME
func GetProjectName() string {
	dir, err := os.Getwd()
	CheckForNil(err)

	folderName := filepath.Base(dir)

	return folderName
}

// ====================================== CAPITALIZE
func Capitalize(word string) string {
	caser := cases.Title(language.English)

	return caser.String(word)
}

// ====================================== FOLDER EXISTS
func FolderExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}

	if info.IsDir() {
		return true
	}

	return false
}

// ====================================== FOLDER EXISTS
func FilsExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	}

	return true
}

// ====================================== FIND WHICH PROJECT IS CALLING GOKU
func CalledFromLocation() string {
	wd, _ := os.Getwd()
	return wd
}

// ====================================== SPINNER
func Spinner(done chan bool, message string) {
	spinChars := `-\|/`
	i := 0
	for {
		select {
		case <-done:
			return
		default:
			fmt.Printf("\r%s %c", message, spinChars[i%len(spinChars)])
			time.Sleep(100 * time.Millisecond)
			i++
		}
	}
}

// ====================================== APPEND TO LAST LINE
func AppendToLastLine(file, data string) {

	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if _, err := f.WriteString(data); err != nil {
		log.Fatal(err)
	}
}

// ====================================== UPDATE PACKAGES
func UpdatePackages(filePath, packages string) {
	data, err := os.ReadFile(filePath)
	CheckForNil(err)

	content := string(data)
	lines := strings.Split(content, "\n")

	for i, line := range lines {
		if line == "package models" {
			lines = append(lines[:i+1], append([]string{packages}, lines[i+1:]...)...)
			break
		}
	}

	err = os.WriteFile(filePath, []byte(strings.Join(lines, "\n")), 0644)
	CheckForNil(err)
}
