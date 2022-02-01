package main

import (
	"fmt"

	"gihub.com/Emanoel01/book-rest-api/src/server"
)

func main() {
	server := server.NewServer()

	fmt.Println("OLA MUNDO")
	server.Run()
}