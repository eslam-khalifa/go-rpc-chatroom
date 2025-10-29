package main

import (
	"bufio"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strings"

	"chatroom/commons"
)

func main() {
	serverAddr := commons.GetServerAddress()
	client, err := rpc.Dial("tcp", serverAddr)
	if err != nil {
		log.Fatal("Error connecting to server:", err)
	}
	defer client.Close()

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Connected to chat server at", serverAddr)

	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Printf("Welcome, %s! Type your messages below (type 'exit' to quit)\n\n", name)

	for {
		fmt.Print("You: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if text == "exit" {
			fmt.Println("Exiting chat...")
			break
		}

		var reply string
		err = client.Call("ChatServer.SendMessage", commons.Message{Name: name, Text: text}, &reply)
		if err != nil {
			log.Println("Error sending message:", err)
			continue
		}

		var chat commons.Reply
		err = client.Call("ChatServer.GetMessages", 0, &chat)
		if err != nil {
			log.Println("Error getting messages:", err)
			continue
		}

		fmt.Println("\n--- Chat History ---")
		for _, msg := range chat.Messages {
			fmt.Println(msg)
		}
		fmt.Println("--------------------\n")
	}
}
