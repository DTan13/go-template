package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/DTan13/template/cmd"
	"github.com/DTan13/template/server"
)

func main() {
	switch strings.ToLower(os.Getenv("MODE")) {
	case "cli":
		// Start cli mode
		cmd.Start()

	case "web":
		// Start the web server
		server.Start()

	default:
		fmt.Println("Please Enter a valid argument")
	}
}
