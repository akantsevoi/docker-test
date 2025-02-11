package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type PingResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}

type Config struct {
	Endpoints []string `json:"endpoints"`
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	response := PingResponse{
		Status:    "ok",
		Timestamp: time.Now(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func pingEndpoints(endpoints []string) {
	client := &http.Client{
		Timeout: 2 * time.Second,
	}

	for _, endpoint := range endpoints {
		go func(url string) {
			for {
				resp, err := client.Get(url)
				if err != nil {
					log.Printf("Error pinging %s: %v", url, err)
				} else {
					log.Printf("Successfully pinged %s, status: %d", url, resp.StatusCode)
					resp.Body.Close()
				}
				time.Sleep(5 * time.Second)
			}
		}(endpoint)
	}
}

func getServiceEndpoints() []string {
	endpointsStr := os.Getenv("ENDPOINTS")
	if endpointsStr == "" {
		panic("should ENDPOINTS")
	}

	return strings.Split(endpointsStr, ",")
}

func main() {
	endpoints := getServiceEndpoints()

	pingEndpoints(endpoints)

	http.HandleFunc("/pings", pingHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server starting on port %s...\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatal(err)
	}
}
