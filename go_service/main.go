package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"os"
)

func main() {
	http.Handle("/metrics", promhttp.Handler())

	port := os.Getenv("PORT")
	portMapping := fmt.Sprintf(":%s", port)
	fmt.Printf("server starting...")

	err := http.ListenAndServe(portMapping, nil)
	if err != nil {
		fmt.Printf("err: %v", err)
		os.Exit(1)
	}
}
