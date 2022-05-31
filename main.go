package main

import (
	"github.com/hyperyuri/webap-with-go/database"
	"github.com/hyperyuri/webap-with-go/server"
)

func main() {
	database.StartDB()

	server := server.NewServer()

	server.Run()

}
