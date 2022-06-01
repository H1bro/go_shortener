package main

import "github.com/H1bro/go_shortener/internal/app/server"

func main() {

	serverAddr := "localhost:8080"

	s := server.Server{Addr: serverAddr}
	s.StartServer()
}
