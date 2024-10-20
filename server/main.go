package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Message struct {
	Body string `json:"body"`
	Count int `json:"count"`
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		msg := getMessageFromMessagePool()
		fmt.Println("took msg pointer from pool:", msg)
		if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
			// http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		time.Sleep(10*time.Millisecond)
		fmt.Println("Recieved message: ", msg.Body, msg.Count)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("vankai")
		putBackMessage(msg)
		return
	}

}
func main() {
	populatePool()
	http.HandleFunc("/message", requestHandler)
	fmt.Println("Serving on 8080 localhost")
	http.ListenAndServe(":8080", nil)
	// log.Fatal(http.ListenAndServe(":8080", nil))
}