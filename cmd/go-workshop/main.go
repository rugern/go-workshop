package main

import (
	"github.com/labstack/gommon/log"
	"net/http"
	"github.com/rugern/go-workshop/pkg/server"
	"github.com/rugern/go-workshop/pkg/database"
)

func main() {
	db, err := database.Setup("mysql", "root:test@tcp(localhost:3306)/workshop", 10, 20)
	if err != nil {
		log.Errorf("Failed to setup DB instance: %v", err)
		panic(err)
	}
	todoDB := database.NewTodoDB(db)
	log.Fatal(http.ListenAndServe(":5000", server.MakeRouter(todoDB)))
}
