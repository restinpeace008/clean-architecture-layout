package main

import (
	cmd "app-module/cmd"
	"app-module/internal/app"
)

func main() {
	// Cobra run
	cmd.Execute()

	// App run
	app.Start()
}
