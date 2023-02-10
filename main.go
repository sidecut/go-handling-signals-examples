package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	// Errors omitted for brevity.

	// Make an HTTP server.
	server := http.Server{
		Addr: ":8080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Simulate a slow HTTP response.
			time.Sleep(10 * time.Second)
			io.WriteString(w, "Hello")
		}),
	}

	// Start the HTTP server in a separate Go routine.
	go func() {
		fmt.Println("Listening for HTTP connections.")
		server.ListenAndServe()
	}()

	// Make a signal channel. Register SIGINT.
	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, os.Interrupt)

	// Wait for signal.
	<-sigch

	fmt.Println("Interrupted. Exiting.")

	// Trigger a shutdown and allow 13 seconds to drain connections. Ignoring
	// CancelFunc for brevity.
	ctx, _ := context.WithTimeout(context.Background(), 13*time.Second)
	server.Shutdown(ctx)
}
