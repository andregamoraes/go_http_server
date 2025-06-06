package main

import (
	"fmt"
	"os"
	"net"
	"flag"
	"go-http-server/app/router"
)

func main() {
	// Parse command line flags
	directory := flag.String("directory", "", "Directory to serve files from")
	flag.Parse()

	server, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
	defer server.Close()

	fmt.Println("Server is listening on port 4221...")

	// The loop that accepts connections and routes them to the router
	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err.Error())
			continue
		}
		// Here we call the router
		go router.HandleConnection(conn, *directory)
	}
}
