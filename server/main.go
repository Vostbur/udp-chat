package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {

	addr := net.UDPAddr{
		Port: 10000,
		IP:   net.ParseIP("127.0.0.1"),
	}
	buffer := make([]byte, 1024)
	clients := []net.UDPAddr{}
	contains := false

	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Start chat. Write STOP for end.")
	defer conn.Close()

	for {
		n, client, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(client)
		
		if strings.TrimSpace(string(buffer[n-4:n])) == "STOP" {
			fmt.Println("Stop chat")
			return
		}

		for i := range(clients) {
			if clients[i].String() == client.String() {
				contains = true
				break
			}
		}
		if contains {
			contains = false
		} else {
			clients = append(clients, *client)
		}

		for i := range(clients) {
			if clients[i].String() != client.String() {
				_, err = conn.WriteToUDP(buffer[0:n], &clients[i])
				if err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}
}
