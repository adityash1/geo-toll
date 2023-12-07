package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/adityash1/toll-calculator/types"
	"net/http"
)

func main() {
	listenAddr := flag.String("listenaddr", ":3000", "the listen address of the HTTP transport handler server")
	flag.Parse()

	store := NewMemoryStore()
	var (
		svc = NewInvoiceAggregator(store)
	)
	makeHTTPTransport(*listenAddr, svc)
	fmt.Println("working fine")
}

func makeHTTPTransport(listenAddr string, svc Aggregator) {
	fmt.Println("http transport running on port ", listenAddr)
	http.HandleFunc("/aggregate", handleAggregate(svc))
	err := http.ListenAndServe(listenAddr, nil)
	if err != nil {
		return
	}
}

func handleAggregate(svc Aggregator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var distance types.Distance
		if err := json.NewDecoder(r.Body).Decode(&distance); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}
