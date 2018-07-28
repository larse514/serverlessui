package main

import (
	"log"
	"os"

	"github.com/larse514/serverlessui/serverless-ui/config"
)

func main() {
	app := config.CreateApp()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
