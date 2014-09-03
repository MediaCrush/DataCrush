package main

import (
	"fmt"

	"github.com/MediaCrush/DataCrush/agent"
	"github.com/MediaCrush/DataCrush/link"
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
	data := make(chan agent.Event)

	for _, host := range hosts {
		go func(target string) {
			ssh := link.NewSSHLink()
			ssh.Connect(target + ":22")

			watch, _ := agent.NewCPUWatchAgent(ssh, target, 0.7, 5)
			defer watch.Stop()

			watch.Run(data)
		}(host)
	}

	for result := range data {
		fmt.Println(result.Source, ":", result.Payload)
	}
}
