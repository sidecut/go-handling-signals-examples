// From https://hjr265.me/blog/go-tidbit-handling-signals-exitting-gracefully/

package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	fmt.Println("Waiting for signal.")

	// Make a buffered channel.
	sigch := make(chan os.Signal, 1)
	// Register the signals that you want to handle.
	signal.Notify(sigch, os.Interrupt)

	// Wait for the signal.
	<-sigch

	fmt.Println("Received interrupt. Exiting.")
}
