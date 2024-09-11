package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	logger *slog.Logger
}

func main() {
	//Parsing command line switches
	addr := flag.String("addr", ":4000", "HTTP Port to Listen On")
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL data source name")
	flag.Parse()

	//Creating a new structured logger to use throughout, just prints to the standard output
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	//Open MySQL connection
	db, err := openDB(*dsn)

	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	defer db.Close()

	//Creating a new application struct to pass for now just the structured logger around as a dependency injection
	app := &application{
		logger: logger,
	}

	logger.Info("starting server", "addr", *addr)

	//Launching server and calling routes.go to create the routes or error out and fail
	err = http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}
	return db, err
}
