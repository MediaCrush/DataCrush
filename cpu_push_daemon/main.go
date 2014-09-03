package main

import (
	"fmt"
	"net"
	"os"

	"github.com/MediaCrush/DataCrush/agent"
	"github.com/MediaCrush/DataCrush/link"
	"labix.org/v2/mgo/bson"
)

var hosts = []string{
	"vox.mediacru.sh",
	"cdn-us-1.mediacru.sh",
	"cdn-us-2.mediacru.sh",
	"cdn-us-3.mediacru.sh",
	"cdn-eu-1.mediacru.sh",
	"cdn-asia-1.mediacru.sh",
}

func main() {
	con, err := net.Dial("udp", "82.130.235.132:9999")
	if err != nil {
		fmt.Println(err)
	}

	data := make(chan agent.Event)

	//for _, host := range hosts {
	//	go func(target string) {
	//		ssh := link.NewSSHLink()
	//		defer ssh.Disconnect()
	//		ssh.Connect(target + ":22")

	//		watch, _ := agent.NewCPUWatchAgent(ssh, target, 0.7, 5)
	//		defer watch.Stop()

	//		watch.Run(data)
	//	}(host)
	//}

	hostname, _ := os.Hostname()

	go func() {
		local := link.NewLocalLink()

		watch, err := agent.NewCPUWatchAgent(local, hostname, 0.7, 1)
		if err != nil {
			panic(err)
		}
		defer watch.Stop()

		watch.Run(data)
	}()

	for result := range data {
		bytes, _ := bson.Marshal(result)
		fmt.Println(result)
		con.Write(bytes)
	}
}
