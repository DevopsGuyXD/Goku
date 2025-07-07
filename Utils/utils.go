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

// ====================================== CREATOR
func Creator() {
	fmt.Printf("\n%v\n%v\n", "With love ❤️", "Bharath Dundi 🤘")
}

// ====================================== ERROR HANDLING
func Check_For_Nil(err error) {
	if err != nil {
		log.Println(err)
		return
	}
}

// ====================================== INSTALL DEPENDENCIES
func Install_Dependencies() {

	init_Swagger()

	done := make(chan bool)
	go Spinner(done, "Installing Dependencies")

	cmd := exec.Command("sh", "-c", "go mod tidy")
	err := cmd.Run()
	if err != nil {
		fmt.Printf("\rInstalling Dependencies ❌\n")
		close(done)
		return
	}

	if !Fils_Exists(".air.toml") {
		init_air()
	}

	close(done)
	fmt.Printf("\rInstalling Dependencies ✔\n\n")
}

// ====================================== INIT SWAGGER
func init_Swagger() {

	done := make(chan bool)
	go Spinner(done, "Updating Swagger")

	calledFrom := Called_From_Location()

	cmd := exec.Command("sh", "-c", fmt.Sprintf("go run github.com/swaggo/swag/cmd/swag@v1.8.12 init --dir \"%s\"", calledFrom))
	err := cmd.Run()
	if err != nil {
		fmt.Printf("\rUpdating Swagger ❌\n")
		close(done)
		return
	}

	close(done)

	fmt.Print("\rUpdating Swagger ✔\n\n")
}

// ====================================== INIT Air
func init_air() {

	calledFrom := Called_From_Location()

	_, err := exec.Command("sh", "-c", fmt.Sprintf("go run github.com/air-verse/air@latest init --dir \"%s\"", calledFrom)).Output()
	Check_For_Nil(err)
}

// ====================================== ERROR HANDLING
func All_Options() {
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
func Open_File(filePath string) *os.File {

	file, err := os.Open(filePath)
	Check_For_Nil(err)

	return file
}

// ====================================== CREATE FOLDER
func Create_Single_Folder(project string) {
	err := os.Mkdir(project, 0755)
	Check_For_Nil(err)
}

// ====================================== CREATE FOLDER
func Create_File(fileName string) *os.File {
	file, err := os.Create(fileName)
	Check_For_Nil(err)

	return file
}

// ====================================== WRITE TO FILE
func Write_File(file *os.File, data string) {
	_, err := file.WriteString(data)
	Check_For_Nil(err)
}

// ====================================== GET PROJECT NAME
func Get_Project_Name() string {
	dir, err := os.Getwd()
	Check_For_Nil(err)

	project := filepath.Base(dir)

	return project
}

// ====================================== CAPITALIZE
func Capitalize(word string) string {
	caser := cases.Title(language.English)

	return caser.String(word)
}

// ====================================== FOLDER EXISTS
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

// ====================================== FOLDER EXISTS
func Fils_Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	}

	return true
}

// ====================================== FIND WHICH PROJECT IS CALLING GOKU
func Called_From_Location() string {
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

// ====================================== CHECK IF LINES EXIST IN FILE
func Check_If_Lines_Exist(filePath string, targets map[string]bool) bool {

	file, err := os.Open(filePath)
	Check_For_Nil(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if _, ok := targets[scanner.Text()]; ok {
			targets[scanner.Text()] = true
		}
	}

	for _, found := range targets {
		if !found {
			return false
		}
	}

	return true
}

// ====================================== INSERT INTO FILE
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
