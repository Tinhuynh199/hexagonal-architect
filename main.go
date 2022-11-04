package main

import (
	"fmt"
	"strconv"

	"hexrestapi1/internal/app"
)

func main() {

	// Loading Config
	config := app.GetConfig()

	// Loading App
	app := app.App{}
	app.Initialize(config)

	// Start Server
	server := ""
	if config.Server.Port != nil {
		server = ":" + strconv.FormatInt(*config.Server.Port, 10)
	}
	fmt.Println("Start server")
	app.Run(server)
}
