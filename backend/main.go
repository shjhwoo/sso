package main

import (
	"fmt"
	"log"
	"net/http"
	"sso/authorizationserver"

	"github.com/rs/cors"
)
var corsHandler = cors.New(cors.Options{
	AllowedOrigins:   []string{"http://localhost:8080","http://localhost:3000"},
	AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
	AllowedHeaders:   []string{"Origin", "Accept", "Content-Type", "X-Requested-With"},
	AllowCredentials: true,
	MaxAge:           0,
	Debug:            true,
})

var handler = corsHandler.Handler(authorizationserver.NewHttpHandler())

var port = "8080"

func main () {
	err := http.ListenAndServe(":"+port, handler)
	fmt.Println("Please open your webbrowser at http://localhost:" + port)
	if err != nil {
		log.Fatal(err)
	}
}
