package main

import (
	"log"
	"net/http"
	"fmt"
)



func new() http.Handler {
	mux := http.NewServeMux()
	// Root
	mux.Handle("/",  http.FileServer(http.Dir("./")))
	// OauthGoogle
	mux.HandleFunc("/auth", oauthByGoogleOauth)
	mux.HandleFunc("/callback", oauthGoogleCallback)

	return mux
}

func main() {
	server := &http.Server{
		Addr: fmt.Sprintf(":80"),
		Handler: new(),
	}

	log.Printf("Starting HTTP Server. Listening at %q", server.Addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("%v", err)
	} else {
		log.Println("Server closed!")
	}
}


