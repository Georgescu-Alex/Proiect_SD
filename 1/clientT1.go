package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) == 1 || len(args) > 2 {
		fmt.Println("Usage: go run .\\client.go (nume) (IP):(Port)")
		return
	}

	Connect := args[1]
	c, err := net.Dial("tcp", Connect)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Write (exit) to close connection")
	fmt.Println("Nume fisier problema(1.txt, 3.txt, 4.txt, 12.txt): ")
	for {
		c_message := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := c_message.ReadString('\n')
		fmt.Fprintf(c, text+"\n")

		message, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Println("Raspuns server: " + message)
		if strings.TrimSpace(string(text)) == "exit" {
			fmt.Println("Close server connection")
			return
		}
	}
}
