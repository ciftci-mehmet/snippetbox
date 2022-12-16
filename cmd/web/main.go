package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/ciftci-mehmet/snippetbox/pkg/models/mysql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golangcollege/sessions"
)

// define application struct that holds application wide dependencies
type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	session       *sessions.Session
	snippets      *mysql.SnippetModel
	templateCache map[string]*template.Template
}

func main() {

	// command line flag addr, default value :4000
	addr := flag.String("addr", ":4000", "HTTP network address")
	// command line flag dsn, mysql dsn string
	dsn := flag.String("dsn", "web:pass1word@tcp(localhost:33066)/snippetbox?parseTime=true", "MySQL data source name")
	// command line flag for session secret 32 bytes long
	secret := flag.String("secret", "s6Ndh+pPbnzHbS*+9Pk8qGWhTzbpa@ge", "Secret key")

	// parse command line flag
	flag.Parse()

	// new logs
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// db
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	// defer to close db if main function exits
	defer db.Close()

	// init new template cache
	templateCache, err := newTemplateCache("./ui/html/")
	if err != nil {
		errorLog.Fatal(err)
	}

	// init new session
	session := sessions.New([]byte(*secret))
	session.Lifetime = 12 * time.Hour

	// application dependencies
	app := &application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		session:       session,
		snippets:      &mysql.SnippetModel{DB: db},
		templateCache: templateCache,
	}

	// init server settings
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	// start listen, serve and log any errors
	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)

}

// openDB function wraps sql.Open() and returnsa a sql.DB connection for given dsn
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
