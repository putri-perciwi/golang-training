package main

import (
	"go-training-restful/routes"
	"os"
)

func main() {
	port := os.Getenv(key: "PORT")
	e := routes.New()

	e.Logger.Fatal(e.Start(":+ PORT"))
}
