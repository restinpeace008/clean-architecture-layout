package main

import (
	cmd "app-module/cmd"
	"app-module/internal/app"
)

func main() {
	cmd.Execute()
	app.Start()
}
