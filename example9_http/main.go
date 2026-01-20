package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var blockHeight int = 0

func blockHeightHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"height":    blockHeight,
		"timestamp": time.Now().Unix(),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func mineHandler(w http.ResponseWriter, r *http.Request) {
	blockHeight++
	response := map[string]interface{}{
		"message":    "Block mined",
		"new_height": blockHeight,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func example9_http() {
	fmt.Println("\n=== Example 9: HTTP Server ===")
	fmt.Println("Starting server on http://localhost:8080")
	fmt.Println("Try: curl http://localhost:8080/height")
	fmt.Println("Try: curl -X POST http://localhost:8080/mine")
	fmt.Printf("Press Ctrl+C to stop")

	http.HandleFunc("/height", blockHeightHandler)
	http.HandleFunc("/mine", mineHandler)

	http.ListenAndServe(":8080", nil)
}

func main() {
	example9_http()
}
