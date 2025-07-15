package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// ============================================================================ ERROR HANDLING
func Check_For_Err(err error) {
	if err != nil {
		log.Println(err)
		return
	}
}

// ============================================================================ INSTALL DEPENDENCIES
func Install_Dependencies() {

	Init_Swagger()

	done := make(chan bool)
	fmt.Println()
	go Spinner(done, "Installing Dependencies")

	cmd := exec.Command("sh", "-c", "go mod tidy")
	err := cmd.Run()
	if err != nil {
		fmt.Printf("\rInstalling Dependencies ❌\n")
		close(done)
		return
	}

	if !Files_Exists(".air.toml") {
		init_air()
	}

	close(done)
	fmt.Printf("\rInstalling Dependencies ✔\n")
}

// ============================================================================ INIT SWAGGER
func Init_Swagger() {

	done := make(chan bool)
	fmt.Println()
	go Spinner(done, "Updating")

	calledFrom := Called_From_Location()

	cmd := exec.Command("sh", "-c", fmt.Sprintf("go run github.com/swaggo/swag/cmd/swag@v1.8.12 init --dir \"%s\"", calledFrom))
	err := cmd.Run()
	if err != nil {

		fmt.Printf("\rUpdating ❌\n")
		close(done)
		return
	}

	close(done)
	fmt.Print("\rUpdating ✔\n")
}

// ============================================================================ INIT Air
func init_air() {

	calledFrom := Called_From_Location()

	_, err := exec.Command("sh", "-c", fmt.Sprintf("go run github.com/air-verse/air@latest init --dir \"%s\"", calledFrom)).Output()
	Check_For_Err(err)
}

// ============================================================================ GET PROJECT NAME
func Project_Name() string {
	dir, err := os.Getwd()
	Check_For_Err(err)

	project := filepath.Base(dir)

	return project
}

// ============================================================================ CAPITALIZE
func Capitalize(word string) string {
	caser := cases.Title(language.English)

	return caser.String(word)
}

// ============================================================================ FIND WHICH PROJECT IS CALLING GOKU
func Called_From_Location() string {
	wd, _ := os.Getwd()
	return wd
}

// ============================================================================ CHECK IF LINES EXIST IN FILE
func Check_If_Lines_Exist(filePath string, targets map[string]bool) bool {

	file, err := os.Open(filePath)
	Check_For_Err(err)
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
