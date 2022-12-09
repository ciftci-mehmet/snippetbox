package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// define application struct that holds application wide dependencies
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {

	// command line flag addr, default value :4000
	addr := flag.String("addr", ":4000", "HTTP network address")

	// parse command line flag
	flag.Parse()

	// new logs
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// application
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	// init server settings
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	// start listen, serve and log any errors
	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)

}
