package main

import (
	"fmt"
	"os"
	"strings"

	configs "github.com/DevopsGuyXD/Goku/Configs"
	templates "github.com/DevopsGuyXD/Goku/Templates"
	utils "github.com/DevopsGuyXD/Goku/Utils"
)

func main() {

	switch {
	// ====================================== -o
	case len(os.Args) == 1:
		utils.AllOptions()

	// ====================================== -v | -i | -o
	case len(os.Args) == 2:
		switch {
		case os.Args[1] == "--version" || os.Args[1] == "-v":
			utils.Version()

		case os.Args[1] == "--install" || os.Args[1] == "-i":
			utils.InstallDependencies()

		case os.Args[1] == "--help" || os.Args[1] == "-h":
			utils.AllOptions()

		case os.Args[1] == "--creator" || os.Args[1] == "-c":
			utils.Creator()

		case os.Args[1] == "add-docker":
			templates.DockerFile()

		default:
			fmt.Printf("Go1: Bad option\n")
		}

	// ====================================== create-project | dev | build | start
	case len(os.Args) == 3:
		switch {
		case os.Args[1] == "run" && os.Args[2] == "dev":
			configs.RunDev()

		case os.Args[1] == "run" && os.Args[2] == "build":
			configs.CreateBuild()

		case os.Args[1] == "run" && os.Args[2] == "start":
			_, err := os.Stat("./dist")
			if err != nil {
				fmt.Printf("\nYou will need to \"build\" your code first\n\n \tRun \"goku run build\"\n")
				os.Exit(0)
			}

			configs.RunProd()

		case os.Args[1] == "docker" && os.Args[2] != "":
			dockerImageName := strings.ToLower(os.Args[2])
			configs.ListDockerImage(dockerImageName)

		case os.Args[1] == "add-crud" && os.Args[2] != "":
			crudName := os.Args[2]
			templates.CRUDTemplate(crudName)

		case os.Args[1] == "create-project" && os.Args[2] != "":
			project := strings.ToLower(os.Args[2])
			templates.StarterTemplate(project)

		case os.Args[1] == "build-docker" && os.Args[2] != "":
			dockerImageName := strings.ToLower(os.Args[2])
			configs.CreateDockerImage(dockerImageName)

		default:
			fmt.Printf("Go1: Bad option\n")
			utils.AllOptions()
		}

	default:
		fmt.Printf("Go1: Bad option\n")
		utils.AllOptions()
	}

}
