package main

import (
	"github.com/labstack/gommon/log"
	"net/http"
	"github.com/rugern/go-workshop/pkg/server"
)

func main() {
	log.Fatal(http.ListenAndServe(":5000", server.MakeRouter()))
}
