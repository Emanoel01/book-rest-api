package main

import (
	"book-rest-api/database"
	"book-rest-api/server"
	"fmt"
)

func main() {
	database.StartDB()
	server := server.NewServer()

	fmt.Println("OLA MUNDO")
	server.Run()
}