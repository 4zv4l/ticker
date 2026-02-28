package main

import (
	"bytes"
	"log"
	"net/http"
	"time"

	"github.com/starfederation/datastar/sdk/go"
)

func main() {
	router := http.NewServeMux()

	// Handler for the initial page load
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ClockPage().Render(r.Context(), w)
	})

	// SSE endpoint for live updates
	router.HandleFunc("/clock", func(w http.ResponseWriter, r *http.Request) {
		sse := datastar.NewSSE(w, r)
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-r.Context().Done(): // Stop the goroutine when the client disconnects
				log.Printf("Client is gone")
				return
			case tick := <-ticker.C: // Render a templ fragment and merge it into the DOM
				log.Printf("Tick: %v", tick)
				buf := &bytes.Buffer{}
				now := time.Now().Format("15:04:05")
				if err := ClockFragment(now).Render(r.Context(), buf); err != nil {
					return
				}
				log.Printf("buf: %v", buf)
				// Send the raw HTML fragment through the SSE generator
				sse.MergeFragments(buf.String())
			}
		}
	})

	log.Printf("Starting server on :3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
