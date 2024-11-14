package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type config struct {
	addr string
}

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	var cfg config

	flag.StringVar(&cfg.addr, "addr", ":4000", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Lshortfile)
	errorLog := log.New(os.Stderr, "Error\t", log.Ldate|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	srv := &http.Server{
		Addr:     cfg.addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting Server on %s", cfg.addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
