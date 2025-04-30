package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// output for http service discovery is here https://prometheus.io/docs/prometheus/latest/http_sd/#http_sd-format
type sdOutput struct {
	Targets []string          `json:"targets"`
	Labels  map[string]string `json:"labels"`
}

func getNodes() []sdOutput {
	// here is where the logic could go for service discovery
	// for example if you have a list of ec2 instances you may
	// query them by tag and return them here
	return []sdOutput{{
		[]string{"localhost:9090"},
		map[string]string{
			"host": "localhost",
			"port": "9090",
		},
	}, {
		[]string{"localhost:9091"},
		map[string]string{
			"host": "localhost",
			"port": "9091",
		},
	}}
}

func serviceDiscoveryHandler(w http.ResponseWriter, r *http.Request) {
	nodes := getNodes()
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(nodes); err != nil {
		fmt.Errorf("failed to encode JSON %v", err)
	}
}

func main() {
	port := os.Getenv("PORT")
	portMapping := fmt.Sprintf(":%s", port)
	http.HandleFunc("/service-discovery", serviceDiscoveryHandler)

	fmt.Printf("server starting...")

	err := http.ListenAndServe(portMapping, nil)
	if err != nil {
		fmt.Printf("err: %v", err)
		os.Exit(1)
	}
}
