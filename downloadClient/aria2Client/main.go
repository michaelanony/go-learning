package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/oliverpool/argo/daemon"
	"github.com/oliverpool/argo/rpc/http"
)

func main() {
	// Error check omitted for concision
	// Start the daemon (without blocking)
	startTime := time.Now()
	defer func(time1 time.Time) {
		log.Printf("cost Seconds:%v\n", time.Since(time1))
	}(startTime)
	cmd := daemon.New().Cmd()
	cmd.Start()

	fmt.Println("Waiting for daemon to listen")
	for !daemon.IsRunningOn(":6800") {
		time.Sleep(time.Second)
	}

	// client to send commands (argo.rpc.http subpackage)
	client := http.NewClient("http://localhost:6800/jsonrpc", "")
	defer client.Close()

	uri := []string{"http://example.com/"}
	_, err := client.AddURI(uri)
	if err != nil {
		panic(err)
	}
	fmt.Println("URI added")

	// let some time for the download
	time.Sleep(time.Second)

	// ask the daemon to shutdown (and wait for its completion)
	client.Shutdown()
	cmd.Wait()

	// remove the downloaded file
	os.Remove("index.html")

}
