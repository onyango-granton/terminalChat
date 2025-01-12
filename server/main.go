package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"strings"
	"sync"
)

type Client struct {
	name string
	Conn net.Conn
}

var (
	clients      = make(map[net.Conn]Client)
	clientsMutex sync.Mutex
	messages     []string

	
)

func main() {
	port := "8989"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	// creating a listener
	listener, err := net.Listen("tcp", port)

	// error Handling
	if err != nil{
		log.Fatalf("Failed to start listener"+err.Error())
	}

	defer listener.Close()

	log.Printf("Server started on port %s\n", port)

	for {
		conn, err := listener.Accept()
		
		// error for incoming connections
		if err != nil{
			log.Printf("Error connection %s", err.Error())
			continue
		}


		
	}
}

func handleConnection(conn net.Conn){
	defer conn.Close()

	//get client name
	reader := bufio.NewReader(conn)
	conn.Write([]byte("Enter your name: "))
	name, err := reader.ReadString('\n')

	// error handling on obtaining name
	if err != nil{
		log.Printf("Error encountered on obtaining client name %v", err.Error())
		return
	}
	name = strings.TrimSpace(name)

	client := Client{name: name, Conn: conn}

	// add new client to client list
	addClient(conn, client)
	

	
}


func addClient(conn net.Conn, client Client){
	clientsMutex.Lock()
	defer clientsMutex.Unlock()
	clients[conn] = client
}

func broadcast(message string, excludeConn net.Conn){
	clientsMutex.Lock()
	defer clientsMutex.Unlock()
	for conn := range clients{
		// broadcast message to everyone save sender
		if conn != excludeConn{			
			conn.Write([]byte(message))
		}
	}
}