package main

import (
	"net/http"
	"log"
	"fmt"
)

func RunServer() {
	handler := MuxHandler()
	s := http.Server{
		Addr: ":3000",
		Handler: handler,
	}
	fmt.Println("Server is running on Port", s.Addr)
	log.Fatal(s.ListenAndServe())
}