package main

import (
	"authentication/data"
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
	"net/http"
	"os"
	"time"
)

const webPort = "80"

var counts int64

type Config struct {
	Repository data.Repository
	Client     *http.Client
}

func main() {
	log.Println("Starting auth service")
	conn := connectToDb()
	if conn == nil {
		log.Panic("Cant connect to pgsql")
	}

	//setup config
	app := Config{
		Client: &http.Client{},
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}
	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func openDb(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func connectToDb() *sql.DB {
	dsn := os.Getenv("DSN")

	for {
		connection, err := openDb(dsn)
		if err != nil {
			log.Println("Pgsql not ready...")
			counts++
		} else {
			log.Println("Pgsql connected")
			return connection
		}

		if counts > 10 {
			log.Println(err)
			return nil
		}

		log.Println("Backing off two seconds...")
		time.Sleep(2 * time.Second)
		continue
	}
}

func (app *Config) setupRepo(conn *sql.DB) {
	db := data.NewPostgresRepository(conn)
	app.Repository = db
}
