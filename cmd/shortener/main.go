package main

import "github.com/H1bro/yandex-practicum-go/internal/app/server"

func main() {

	serverAddr := "localhost:8080"

	s := server.Server{Addr: serverAddr}
	s.StartServer()
}
