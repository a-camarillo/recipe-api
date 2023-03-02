package main

import (
	"net/http"
	"log"
	"fmt"
)

func RunServer() {
	handler := ApiRouteHandler()
	s := http.Server{
		Addr: ":3000",
		Handler: handler,
	}
	fmt.Println("Server is running on Port", s.Addr)
	log.Fatal(s.ListenAndServe())
}