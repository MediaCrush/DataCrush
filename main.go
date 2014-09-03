package main

import (
	"fmt"
	"github.com/MediaCrush/DataCrush/link/ssh"
	"sync"
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
	data := make(chan string)
	wg := sync.WaitGroup{}

	for _, host := range(hosts) {
		go func(target string) {
			wg.Add(1)

			link := ssh.NewSSHLink()
			link.Connect(target + ":22")
			defer link.Disconnect()
			data <- link.Run("uptime")
			data <- link.Run("hostname")

			wg.Done()
		}(host)
	}

	go func() { wg.Wait(); close(data); }()

	for result := range(data) {
		fmt.Println(result)
	}
}
