package main

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/MediaCrush/DataCrush/agent"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", ":9999")
	sock, err := net.ListenUDP("udp", addr)

	if err != nil {
		fmt.Println(err)
	}

	for {
		buf := make([]byte, 1024)
		bytes, err := sock.Read(buf)
		if err != nil {
			fmt.Println(err)
		}

		evt := &agent.Event{}
		err = json.Unmarshal(buf[:bytes], evt)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(bytes, evt)
	}

}
