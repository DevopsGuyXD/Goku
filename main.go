package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	configs "github.com/DevopsGuyXD/Goku/Configs"
	templates_curd "github.com/DevopsGuyXD/Goku/Templates/crud"
	templates_starter "github.com/DevopsGuyXD/Goku/Templates/starter"
	utils "github.com/DevopsGuyXD/Goku/Utils"
)

func main() {

	switch {
	// ============================================================================ -o
	case len(os.Args) == 1:
		utils.All_Options()

	// ============================================================================ -v | -i | -o
	case len(os.Args) == 2:
		switch {
		case os.Args[1] == "--version" || os.Args[1] == "-v":
			utils.Version()

		case os.Args[1] == "--install" || os.Args[1] == "-i":
			utils.Install_Dependencies()

		case os.Args[1] == "--help" || os.Args[1] == "-h":
			utils.All_Options()

		case os.Args[1] == "--creator" || os.Args[1] == "-c":
			utils.Creator()

		case os.Args[1] == "moby":
			utils.Create_File([]string{"dockerfile"})
			utils.Write_File(utils.Open_File("dockerfile"), templates_starter.DockerFile_Data())

		case os.Args[1] == "m-build":
			configs.Build_Docker_Image()

		case os.Args[1] == "swag":
			utils.Init_Swagger()

		case os.Args[1] == "scan":
			utils.RunScan()

		case os.Args[1] == "test":
			configs.Run_Tests()

		default:
			fmt.Printf("\nGoku: Invalid option. Please use one of the supported options.\n\n ☆  goku -h\n")
		}

	// ============================================================================ create-project | dev | build | start
	case len(os.Args) == 3:
		switch {
		case os.Args[1] == "run" && os.Args[2] == "dev":
			configs.Run_Dev()

		case os.Args[1] == "run" && os.Args[2] == "build":
			configs.Create_Build("--scan=true")

		case os.Args[1] == "run" && os.Args[2] == "start":
			_, err := os.Stat("./dist")
			if err != nil {
				fmt.Printf("\nBuild step required: Please compile your code before proceeding.\n\n Use the following command:\n goku run build\n")
				os.Exit(0)
			}

			configs.Run_Prod()

		case os.Args[1] == "add-crud" && os.Args[2] != "":
			crudName := os.Args[2]
			templates_curd.CRUD_Project(crudName)

		case os.Args[1] == "create-project" && os.Args[2] != "":
			project := strings.ToLower(os.Args[2])
			templates_starter.Starter_Project(project)

		case os.Args[1] == "m-run" && os.Args[2] != "":
			port, err := strconv.Atoi(os.Args[2])
			utils.Check_For_Err(err)

			configs.Run_Docker_Image(port, "")

		default:
			fmt.Printf("\nGoku: Invalid option. Please use one of the supported options.\n\n ☆  goku -h\n")
		}

	// ============================================================================ create-project | dev | build | start
	case len(os.Args) == 4:
		switch {

		case os.Args[1] == "run" && os.Args[2] == "build" && os.Args[3] == "--scan=false":
			configs.Create_Build(os.Args[3])
		}

	default:
		fmt.Printf("\nGoku: Invalid option. Please use one of the supported options.\n\n ☆  goku -h\n")
	}

}
