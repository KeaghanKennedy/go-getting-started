package main

import (
	"net/http"

	"github.com/KeaghanKennedy/go-getting-started/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
	// notes.PrimitiveDataTypes()
	// notes.Collections()
	// notes.Functions()
}
