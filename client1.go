package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

//Queue represents a queue that holds a slice
type Queue struct {
	items []string
}

//Enqueue adds the item in the Queue
func (q *Queue) Enqueue(text string) {
	q.items = append(q.items, text)

}

func main() {
	myQueue := Queue{}
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	reader := bufio.NewReader(os.Stdin)
	for {

		fmt.Print("message sent: ")
		//reading the string and holding it in text variable
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text+"\n")
		myQueue.Enqueue(text)
		fmt.Println(myQueue)

		//closing the connection
		if strings.TrimSpace(string(text)) == "close" {
			fmt.Println("closed")
			return
		}
	}
}
