package main

import (
	"fmt"
	"net/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	fmt.Println("Starting ML Monitoring Agent on :8080")
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}