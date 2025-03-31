package main

import (
	"fmt"
	"time"

	probing "github.com/andrewsjg/pro-bing"
)

func pingHost(ip string) bool {
	pinger, err := probing.NewPinger(ip)
	if err != nil {
		return false
	}

	pinger.Count = 1
	pinger.Timeout = time.Second
	pinger.SetPrivileged(true) //Windows must have

	err = pinger.Run()
	if err != nil {
		return false
	}

	stats := pinger.Statistics()
	return stats.PacketsRecv > 0
}

func main() {
	subnet := "192.168.1" //TODO Find subnet

	fmt.Println("Skanowanie sieci:", subnet+".x")

	for i := 1; i <= 254; i++ {
		ip := fmt.Sprintf("%s.%d", subnet, i)
		if pingHost(ip) {
			fmt.Println("Host aktywny:", ip)
		}
	}
}
