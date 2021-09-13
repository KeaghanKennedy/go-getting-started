package notes

import "fmt"

// 6.2 - 6.5 FUNCTIONS

// Paramaters are declared in the method signature.
// Arguments are passed in when the functions is used, but they're basically the same.
func startWebServer(port, numRetries int) (int, error) {
	fmt.Println("Starting server...")
	fmt.Println("Server started on port", port)
	return port, nil
}

func Functions() {
	port := 3000
	// Underscore here allows us to ignore the returned port value by dumping it into a write only variable.
	// If we didn't do this we'd have to use it or the compiles would complain.
	_, err := startWebServer(port, 3)
	fmt.Println(err)

}
