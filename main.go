package main

import (
	"os"

	"github.com/diegosanchez/ms_forms/server"
)

func main() {
	port := os.Getenv("PORT")

	router := server.NewMsFormServer(port)

	router.Run()
}
