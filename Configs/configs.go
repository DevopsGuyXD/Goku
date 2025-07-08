package configs

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"runtime"

	utils "github.com/DevopsGuyXD/Goku/Utils"
)

// ============================================================================ RUN DEV
func Run_Dev() {

	var shell, flag string
	calledFrom := utils.Called_From_Location()

	if runtime.GOOS == "windows" {
		shell = "cmd.exe"
		flag = "/C"
	} else {
		shell = "sh"
		flag = "-c"
	}

	utils.Message("🔧 Running in Dev mode")

	cmd := exec.Command(shell, flag, fmt.Sprintf("go run github.com/air-verse/air@latest air --dir \"%s\"", calledFrom))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	utils.Check_For_Nil(err)

	os.Exit(0)
}

// ============================================================================ CREATE BUILD
func Create_Build() {

	done := make(chan bool)

	if runtime.GOOS == "windows" {
		go utils.Spinner(done, "Building your app")

		cmd := exec.Command("sh", "-c", "go build -o ./dist/app.exe . && cp .env ./dist")
		err := cmd.Run()
		if err != nil {
			fmt.Printf("\rBuilding your app ❌\n")
			close(done)
			return
		}

		close(done)
		fmt.Print("\rBuilding your app ✔\n")

	} else {
		go utils.Spinner(done, "Building your app")

		cmd := exec.Command("sh", "-c", "go build -o ./dist/app . && cp .env ./dist")
		err := cmd.Run()
		if err != nil {
			fmt.Printf("\rBuilding your app ❌\n")
			close(done)
			return
		}

		close(done)
		fmt.Print("\rBuilding your app ✔\n")
	}
}

// ============================================================================ DOCKER BUILD
func Create_Docker_Image(dockerImageName string) {
	cmd := exec.Command("sh", "-c", fmt.Sprintf("docker build -t %s .", dockerImageName))

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error getting StdoutPipe:", err)
		return
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Println("Error getting StderrPipe:", err)
		return
	}

	if err := cmd.Start(); err != nil {
		fmt.Println("Error starting command:", err)
		return
	}

	// Regex to remove BuildKit prefixes like: #5 or #12
	buildkitPrefix := regexp.MustCompile(`^#\d+\s*`)

	// Stream stdout
	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			line := buildkitPrefix.ReplaceAllString(scanner.Text(), "")
			fmt.Println(line)
		}
	}()

	// Stream stderr
	go func() {
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			line := buildkitPrefix.ReplaceAllString(scanner.Text(), "")
			fmt.Println(line)
		}
	}()

	if err := cmd.Wait(); err != nil {
		fmt.Println("Command finished with error:", err)
		return
	}

	fmt.Printf("\nDocker build completed successfully \n")
}

func List_Docker_Image(dockerImageName string) {
	res, err := exec.Command("sh", "-c", fmt.Sprintf("docker image ls %s", dockerImageName)).Output()
	utils.Check_For_Nil(err)

	fmt.Printf("\n%v", string(res))
}

// ============================================================================ RUN PRODUCTION
func Run_Prod() {

	var shell, flag string

	if runtime.GOOS == "windows" {
		shell = "cmd.exe"
		flag = "/C"

		utils.Message("🔥 Running in Production mode")

		cmd := exec.Command(shell, flag, "go run .")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin

		err := cmd.Run()
		utils.Check_For_Nil(err)

		os.Exit(0)
	} else {
		shell = "sh"
		flag = "-c"

		utils.Message("🔥 Running in Production mode")

		cmd := exec.Command(shell, flag, "go run .")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin

		err := cmd.Run()
		utils.Check_For_Nil(err)

		os.Exit(0)
	}
}
