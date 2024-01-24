package app

import (
	"log"
	"net/http"

	"github.com/dexguitar/chatapp/configs"
	"github.com/dexguitar/chatapp/db"
)

func Run(c *configs.Config) {
	// Postgres
	dbConn, err := db.NewDatabase(c.Postgres)
	if err != nil {
		log.Fatalf("could not initialize database connection: %s", err.Error())
	}
	err = dbConn.GetDB().Ping()
	if err != nil {
		log.Fatalf("failed to ping DB: %s", err.Error())
	}
	defer dbConn.Close()

	// Handler
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello there!"))
	})

	http.ListenAndServe(c.Host+c.Port, mux)
}
