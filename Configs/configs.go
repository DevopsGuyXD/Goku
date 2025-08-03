package configs

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"

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
	utils.Check_For_Err(err)

	os.Exit(0)
}

// ============================================================================ CREATE BUILD
func Create_Build(scan string) {

	done := make(chan bool)

	if scan == "--scan=true" {
		utils.RunScan()
	} else {
		fmt.Printf("\nWarning: Scan skipped. It is strongly recommended not to skip this step\n")
	}

	if runtime.GOOS == "windows" {
		fmt.Println()
		go utils.Spinner(done, "Building your app")

		cmd := exec.Command("sh", "-c", "go build -o ./dist/app.exe")
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

		cmd := exec.Command("sh", "-c", "go build -o ./dist/app")
		err := cmd.Run()
		if err != nil {
			fmt.Printf("\rBuilding your app ❌\n")
			close(done)
			return
		}

		close(done)
		fmt.Println()
		fmt.Print("\rBuilding your app ✔\n")
	}
}

// ============================================================================ RUN PRODUCTION
func Run_Prod() {

	var shell, flag string

	if runtime.GOOS == "windows" {
		shell = "cmd.exe"
		flag = "/C"

		file, err := filepath.Glob(utils.Called_From_Location() + "/dist/*.exe")
		utils.Check_For_Err(err)

		utils.Message("🔥 Running in Production mode")

		cmd := exec.Command(shell, flag, file[0])
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin

		err = cmd.Run()
		utils.Check_For_Err(err)

		os.Exit(0)
	} else {
		shell = "sh"
		flag = "-c"

		file, err := filepath.Glob(utils.Called_From_Location() + "/dist/*")
		utils.Check_For_Err(err)

		utils.Message("🔥 Running in Production mode")

		cmd := exec.Command(shell, flag, file[0])
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin

		err = cmd.Run()
		utils.Check_For_Err(err)

		os.Exit(0)
	}
}

// ============================================================================ RUN TESTS
func Run_Tests() {
	var shell, flag string
	// calledFrom := utils.Called_From_Location()

	if runtime.GOOS == "windows" {
		shell = "cmd.exe"
		flag = "/C"
	} else {
		shell = "sh"
		flag = "-c"
	}

	utils.Message("🧪 Running tests")

	cmd := exec.Command(shell, flag, "go test ./Tests -count=1")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err != nil {
		fmt.Println("\nTests will be implemented for the CRUD functionality")
	}

	os.Exit(0)
}

// ============================================================================ DOCKER BUILD
func Build_Docker_Image() {

	data := utils.Open_File(utils.Called_From_Location() + "/.env")
	line := utils.ReturnLineFromFile(data)
	parts := strings.Split(line, `"`)

	image := strings.Split(utils.Called_From_Location(), `\`)

	cmd := exec.Command("sh", "-c", fmt.Sprintf("docker build --build-arg PORT=%v -t %s .", parts[1], image[len(image)-1]))

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

	buildkitPrefix := regexp.MustCompile(`^#\d+\s*`)

	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			line := buildkitPrefix.ReplaceAllString(scanner.Text(), "")
			fmt.Println(line)
		}
	}()

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

// ============================================================================ RUN DOCKER IMAGE
func Run_Docker_Image(port int, image string) {

	var shell, flag string
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		shell = "cmd.exe"
		flag = "/C"
	} else {
		shell = "sh"
		flag = "-c"
	}

	utils.Message("🐳 Running container")

	if image == "" {
		default_image := strings.Split(utils.Called_From_Location(), `\`)
		cmd = exec.Command(shell, flag, fmt.Sprintf("docker run -p %d:8000 %s", port, default_image[len(default_image)-1]))

	} else {
		cmd = exec.Command(shell, flag, fmt.Sprintf("docker run -p %d:8000 %s", port, image))
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	utils.Check_For_Err(err)

	os.Exit(0)
}
