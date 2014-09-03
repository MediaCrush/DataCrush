package main

import (
	"fmt"
	"net"

	"github.com/MediaCrush/DataCrush/agent"
	"labix.org/v2/mgo/bson"
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
		bson.Unmarshal(buf, evt)

		fmt.Println(bytes, evt)
	}

}
