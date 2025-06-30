package utils

import (
	"bufio"
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

// ====================================== ERROR HANDLING
func CheckForNil(err error) {
	if err != nil {
		log.Println(err)
		return
	}
}

// ====================================== INSTALL DEPENDENCIES
func InstallDependencies() {

	Swagger()

	done := make(chan bool)
	go Spinner(done, "Installing Dependencies")

	cmd := exec.Command("sh", "-c", "go mod tidy")
	err := cmd.Run()
	if err != nil {
		fmt.Printf("\rInstalling Dependencies ❌\n")
		close(done)
		return
	}

	close(done)
	fmt.Printf("\rInstalling Dependencies ✅\n")
}

// ====================================== INIT SWAGGER
func Swagger() {

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
	fmt.Print("\rUpdating Swagger ✅ \n")
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

// ====================================== CREATE FOLDER
func CreateSingleFolder(folderName string) {
	err := os.Mkdir(folderName, 0755)
	CheckForNil(err)
}

// ====================================== CREATE FOLDER
func CreateFile(fileName string) {
	file, err := os.Create(fileName)
	CheckForNil(err)
	defer file.Close()
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

// ====================================== UPDATING MAIN WITH DATABASE
func UpdatingMain(crudName string) {

	var lines []string
	filePath := "./Models/models.go"

	data := fmt.Sprintf("%v()", crudName)

	file, err := os.Open(filePath)
	CheckForNil(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)

		if strings.Contains(line, "func AppModels(){") {
			lines = append(lines, "\t"+data)
		}
	}

	CheckForNil(scanner.Err())

	err = os.WriteFile(filePath, []byte(strings.Join(lines, "\n")), 0644)
	CheckForNil(err)
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

// ====================================== INIT SPINNER
// func InitSpinner(message string) {
// 	done := make(chan bool)
// 	go Spinner(done, message)
// 	time.Sleep(1 * time.Second)
// 	done <- true
// }
