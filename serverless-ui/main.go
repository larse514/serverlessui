package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	// Setup seelog for amazon-ecr-credential-helper

	app := cli.NewApp()
	err := app.Run(os.Args)
	if err != nil {
		println("poop")
	}
}
