package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"goapp/internal/pkg/watcher"

	"github.com/gorilla/websocket"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetFlags(0)
}

func main() {

	wg := sync.WaitGroup{}
	done := make(chan struct{})
	exitChannel := make(chan os.Signal, 1)
	signal.Notify(exitChannel, syscall.SIGINT, syscall.SIGTERM)

	addr := flag.String("addr", "localhost:8080", "http service address")
	n := flag.Int("n", 5, "number of connections")
	path := flag.String("path", "/goapp/ws", "path to websocket")
	help := flag.Bool("help", false, "display help message")

	flag.Usage = func() {
		fmt.Println("Usage: goapp [options]")
		fmt.Println("Options:")
		flag.PrintDefaults()
	}

	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	u := url.URL{Scheme: "ws", Host: *addr, Path: *path}

	for i := 0; i < *n; i++ {
		c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
		if err != nil {
			log.Fatalf("Error connecting to server: %v", err)
		}
		wg.Add(1)

		go func(i int, c *websocket.Conn) {
			defer wg.Done()
			defer c.Close()
			for {
				select {
				default:
					var counter watcher.Counter

					_, message, err := c.ReadMessage()
					if err != nil {
						log.Println("Error reading message:", err)
						return
					}

					err = json.Unmarshal(message, &counter)
					if err != nil {
						log.Println("Error unmarshalling JSON:", err)
						return
					}

					log.Printf("[conn #%d] iteration: %d, value: %s\n", i, counter.Iteration, counter.Value)

				case <-done:
					log.Printf("[conn #%d] closing...", i)
					// Send close message to the server before closing the connection
					if err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "")); err != nil {
						log.Println("Error sending close message:", err)
					}
					return
				}
			}
		}(i, c)
	}

	<-exitChannel
	log.Println("Shutting down...")
	close(done)
	wg.Wait()
	log.Println("All connections closed")
}
