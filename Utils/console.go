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
  Options:

    -h | --help
    -v | --version
    -i | --install

    goku create-project mytestapp
    goku run dev
    goku run build
    goku run start
    goku add-crud <NAME>
	goku swag
	goku add-docker
    goku docker <NAME:TAG> -> Note: TAG will be "latest" if not specified
	goku dl <NAME>
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
	fmt.Printf("\n%v\n\n", message)
}
