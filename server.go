package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"sync"

	"chatroom/commons"
)

type ChatServer struct {
	Messages []string
	Mutex    sync.Mutex
}

func (s *ChatServer) SendMessage(msg commons.Message, reply *string) error {
	formatted := fmt.Sprintf("%s: %s", msg.Name, msg.Text)

	s.Mutex.Lock()
	s.Messages = append(s.Messages, formatted)
	s.Mutex.Unlock()

	fmt.Println(formatted)

	*reply = "Message received"
	return nil
}

func (s *ChatServer) GetMessages(dummy int, reply *commons.Reply) error {
	s.Mutex.Lock()
	reply.Messages = append([]string{}, s.Messages...)
	s.Mutex.Unlock()
	return nil
}

func main() {
	server := new(ChatServer)
	rpc.Register(server)

	addr := commons.GetServerAddress()
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("Listen error:", err)
	}

	fmt.Println("Chat server running on", addr)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(conn)
	}
}
