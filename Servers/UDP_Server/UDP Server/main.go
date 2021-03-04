package main

import (
	"fmt"
	"net"
	"encoding/hex"
)

func handleConnection(con *net.UDPConn) {
	var buf [2048]byte
	n, err := con.Read(buf[0:])
	if err != nil {
		fmt.Println("Error reading")
		return
	} else {
		fmt.Println(hex.EncodeToString(buf[0:n]))
		fmt.Println("Done!")
	}
}

func StartServer() {
	port := ":8080"
	protocol := "udp"
	
	adr,err := net.ResolveUDPAddr(protocol, port)
	if err != nil {
		fmt.Println("Wrong addr")
		return
	}
	
	fmt.Println("Server working on addr " + adr.String())
	
	list, err := net.ListenUDP(protocol, adr)
	if err != nil {
		fmt.Println("Error creating server")
		return
	}
	
	for {
		handleConnection(list)
	}
}

func main() {
	StartServer()
}
