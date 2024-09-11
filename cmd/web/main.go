package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {
	//Parsing command line switches
	addr := flag.String("addr", ":4000", "HTTP Port to Listen On")
	flag.Parse()

	//Creating a new structured logger to use throughout, just prints to the standard output
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	//Creating a new application struct to pass for now just the structured logger around as a dependency injection
	app := &application{
		logger: logger,
	}

	logger.Info("starting server", "addr", *addr)

	//Launching server and calling routes.go to create the routes or error out and fail
	err := http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
