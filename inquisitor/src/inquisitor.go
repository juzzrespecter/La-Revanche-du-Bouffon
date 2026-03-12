package main

import (
	"context"
	"fmt"

	/* "flag" */
	"os/signal"
	"syscall"
)

func signalHandler(s chan int) {
	ctx, stop := signal.NotifyContext(
		context.Background(), syscall.SIGINT,
	)
	defer stop()
	fmt.Println("Waiteando signal")
	<-ctx.Done()

	fmt.Println(context.Cause(ctx))
	fmt.Println("End of signal handler")
	s <- 1
}

func init() {

}

func main() {
	// ips
	// macs

	fd, _ := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_ICMP)
	// goroutine to handle signal
	// goroutine to spoof loop

	sigEnd := make(chan int)
	go signalHandler(sigEnd)
	end := <-sigEnd
	fmt.Println(fd, end)
	return
}
