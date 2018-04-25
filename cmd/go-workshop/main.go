package main

import (
	"github.com/labstack/gommon/log"
	"net/http"
	"github.com/rugern/go-workshop/pkg/server"
	"github.com/rugern/go-workshop/pkg/database"
)

func main() {
	todoDB := database.NewTodoDB()
	log.Fatal(http.ListenAndServe(":5000", server.MakeRouter(todoDB)))
}
