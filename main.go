package main

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/routes"
	// "github.com/tejas-cogo/go-cogoport/cmd"
)

func main() {

	// cmd.Execute()
	e := godotenv.Load()

	if e != nil {
		fmt.Print(e)
	}

	config.Connect()

	r := routes.SetupRouter()

	port := os.Getenv("port")

	if len(os.Args) > 1 {
		reqPort := os.Args[1]
		if reqPort != "" {
			port = reqPort
		}
	}

	if len(os.Args) > 1 {
		reqPort := os.Args[1]
		if reqPort != "" {
			port = reqPort
		}
	}

	if port == "" {
		port = "8080" //localhost
	}
	type Job interface {
		Run()
	}

	r.Run(":" + port)

}
