package main

import (
	"context"
	"fmt"
	"net"
	/* "flag" */
	"os/signal"
	"syscall"
)

func signalHandler() {
	ctx, stop := signal.NotifyContext(
		context.Background(), syscall.SIGINT,
	)
	defer stop()
	fmt.Println("Waiteando signal")
	<-ctx.Done()

	fmt.Println(context.Cause(ctx))
}

func main() {
	ips := []string{"1.2.34.5", "256.1.2.3", "0.1.2.3", "2.3", "23.23.23.23"}

	for _, ip := range ips {
		parsed := net.ParseIP(ip)
		if parsed == nil {
			fmt.Printf("%s is invalid\n", ip)
		} else {
			fmt.Printf("%s is valid\n", ip)
		}
	}

	macs := []string{"3A:7F:91:2C:44:B8","A6:1D:5E:93:70:2F","4C:82:19:AF:6D:E3","D2:58:0B:77:9A:41","e"}
	for _, mac := range macs {
		parsed, _ := net.ParseMAC(mac)
		if parsed == nil {
			fmt.Printf("%s is invalid\n", mac)
		} else {
			fmt.Printf("%s is valid\n", mac)
		}
	}

	fd, _ := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_ICMP)
	// goroutine to handle signal
	// goroutine to spoof loop

	go signalHandler()
	fmt.Println(fd)
}