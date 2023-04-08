package main

import "github.com/met4tron/doctors/internal/server"

func main() {
	server := server.Init()

	server.Listen()
}
