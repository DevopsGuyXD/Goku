package utils

import (
	"fmt"
	"time"
)

// ============================================================================ GOKU VERSION
func Version() {
	fmt.Println(
		`    ____    ___    _  __   _   _ 
   / ___|  / _ \  | |/ /  | | | |
  | |  _  | | | | | ' /   | | | |
  | |_| | | |_| | | . \   | |_| |
   \____|  \___/  |_|\_\   \___/  v1.0.0`)
}

// ============================================================================ CREATOR
func Creator() {
	fmt.Printf("\n%v\n%v\n", "With love ❤️", "Bharath Dundi 🤘")
}

// ============================================================================ ERROR HANDLING
func All_Options() {
	fmt.Printf(`  
  ⭐ Options:

      -h, --help        Show help information  
      -v, --version     Show CLI version  
      -i, --install     Install project dependencies


  🚀 Project Commands:

      goku create-project <project-name> | Create a new Goku project


  🔧 Run & Build:

      goku run dev                       | Start the project in development mode
      goku run build                     | Build/compile the project for production. Use --scan=false to skip gosec from scanning your code
      goku run start                     | Run the compiled project in production mode


  ⚙️  Feature Additions:

      goku add-crud <name>               | Generate CRUD logic for the specified resource
      goku swag                          | Generate or update Swagger documentation


  🧪 Testing & Validation:

      goku test                          | Run all project tests to validate functionality

      
  🐳 Docker Integration:

    ☆ Please ensure that Docker is installed and actively running on your system ☆

      goku dock                          | Add a Dockerfile to the project
      goku dock-build                    | Build a Docker image for the project
      goku dock-list                     | List Docker images associated with the project
      goku dock-tag <name:tag>           | Tag Docker images associated with the project
      goku dock-run <port> <name:tag>    | Run the Docker image. Defaults to the project image if <name:tag> is empty
`)
}

// ============================================================================ SPINNER
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

// ============================================================================ MESSAGE
func Message(message string) {
	fmt.Printf("\n%v\n", message)
}
