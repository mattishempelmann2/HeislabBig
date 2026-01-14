package main

import (
	//"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {

	const address = "10.100.23.11:33546"
	ConnectMessage := "Connect to: 10.100.23.24:33546\x00"
	SecondMessage := "Hello from group 13\x00"
	Port := ":33546"

	l, err := net.Listen("tcp", Port)
	if err != nil {
		fmt.Printf("Error listener: %s \n", err)
		return
	}
	defer l.Close()
	fmt.Printf("Listening on port %s\n", Port)

	timeout := 5 * time.Second

	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		fmt.Printf("Connection failed: %s\n", err)
		return
	}
	defer conn.Close()
	fmt.Printf("Connected to %s\n", address)

	_, err = conn.Write([]byte(ConnectMessage))
	if err != nil {
		fmt.Printf("Failed to send message: %s\n", err)
		return
	}
	fmt.Printf("Connect Message sent\n")
	fmt.Printf("Waiting for reply\n")

	newConn, err := l.Accept()
	if err != nil {
		log.Fatal(err)
	}
	defer newConn.Close()

	fmt.Printf("Accepted connection from %s\n", newConn.RemoteAddr())

	_, err = conn.Write([]byte(SecondMessage))

	if err != nil {
		fmt.Printf("Failed to send Second message: %s\n", err)
		return
	}
	fmt.Printf("Message2 sent\n")

	io.Copy(os.Stdout, newConn)

}
