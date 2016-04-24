package mcping

import (
	"testing"
	"fmt"
)

const TestAddress = "us.mineplex.com:25565"

func TestPing(t *testing.T) {
	response, err := Ping(TestAddress)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("Latency, %v\n", response.Latency)
		fmt.Printf("Online/Max, %v/%v\n", response.Online, response.Max)
		fmt.Printf("Protocol, %v, Server %v\n", response.Protocol, response.Server)
		fmt.Printf("Motd, %v\n", response.Motd)
		fmt.Printf("Favicon len, %v\n", len(response.Favicon))
	}
}