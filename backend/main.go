package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	dbConfig := DBConfig{
		Host:     os.Getenv("POSTGRES_HOST"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Name:     os.Getenv("POSTGRES_NAME"),
		Port:     os.Getenv("POSTGRES_PORT"),
		SSLMode:  "disable",
		Tz:       os.Getenv("POSTGRES_TZ"),
	}
	/*dbConfig := DBConfig{
		Host:     "localhost",
		User:     "postgres",
		Password: "password",
		Name:     "postgres",
		Port:     "5432",
		SSLMode:  "disable",
		Tz:       "Asia/Novosibirsk",
	}*/
	dbe, _ := NewDBEngine(dbConfig)

	router := mux.NewRouter()
	RegisterRoutes(router, dbe)
	headersOk := handlers.AllowedHeaders([]string{"*"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ON"), handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}
