package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"net"
)


const PORT = 6881

func main() {
	addr, err := net.ResolveUDPAddr("udp", "router.bittorrent.com:6881")
	if err != nil {
		fmt.Println(err)
		log.Panic(err)
	}
	fmt.Println(addr.IP)
	//listenAddr := new(net.UDPAddr)
	//listener, err := net.ListenUDP("udp", listenAddr)
	//fmt.Println(listener)
	id := randomID()
	fmt.Println(hex.EncodeToString(id))
}
