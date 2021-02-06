package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	//start the server
	ln, err := net.Listen("tcp", ":8080")
	fmt.Println("server is ready now.")
	//error handling
	if err != nil {
		fmt.Println(err)
		return
	}
	//closing the listener
	defer ln.Close()
	for {
		//Accept waits and returns the next listener
		c, err := ln.Accept()
		//error handling
		if err != nil {
			fmt.Println(err)
			return

		}

		go handleClientRequest(c)

	}
}

func handleClientRequest(con net.Conn) {
	defer con.Close()

	for {
		//reading the string
		netData, err := bufio.NewReader(con).ReadString('\n')
		//error handling
		if err != nil {
			fmt.Println(err)
			return
		}
		//closing the connection
		if strings.TrimSpace(netData) == "close" {
			fmt.Println("closed")
			return
		}
		fmt.Print("message received:")
		fmt.Print(string(netData))
		
	}
}
