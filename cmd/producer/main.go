package main

import (
	"bytes"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/AIdenTale/go-hack.git/internal/producer"
)

type bpmPayload struct {
	BPM float64 `json:"bpm"`
}

type tracPayload struct {
	Trac float64 `json:"trac"`
}

func main() {
	cfg, err := producer.LoadConfig("config/producer.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	start := time.Now()
	go sendData(cfg.BPM, start, "bpm")
	go sendData(cfg.Trac, start, "trac")

	select {} // run forever
}

func sendData(cfg producer.EndpointConfig, start time.Time, mode string) {
	interval := time.Duration(float64(time.Second) / cfg.FreqHz)
	for i := 0; i < cfg.Count; i++ {
		val := rand.NormFloat64()*5 + cfg.Mean // нормальное распределение, stddev=5
		var payload any
		if mode == "bpm" {
			payload = []bpmPayload{{BPM: val}}
		} else {
			payload = []tracPayload{{Trac: val}}
		}
		b, _ := json.Marshal(payload)
		baseURL := os.Getenv("RECEIVER_URL")
		if baseURL == "" {
			baseURL = "http://localhost:8080"
		}

		log.Println("start sending to", baseURL+cfg.Endpoint)
		_, err := http.Post(baseURL+cfg.Endpoint, "application/json", bytes.NewReader(b))
		if err != nil {
			log.Printf("send error: %v", err)
		}

		time.Sleep(interval)
	}
}
