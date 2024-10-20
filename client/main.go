package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Message struct {
	Body  string `json:"body"`
	Count int    `json:"count"`
}

func main() {
	wg := sync.WaitGroup{}
	body := Message{Body: "Hello, World!", Count: 12}
	jsonMessage, err := json.Marshal(body)
	if err != nil {
		fmt.Println("Error marshalling JSON")
		return
	}
	url := "http://localhost:8080/message"
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonMessage))
		if err != nil {
			fmt.Println("Error creating HTTP request")
			return
		}

		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error sending request: ", err)
			return
		}
		defer resp.Body.Close()

		var resBody bytes.Buffer
		resBody.ReadFrom(resp.Body)
		fmt.Println("Response and status code:", resBody.String(), resp.Status)
		wg.Done()
		}()
	}
	wg.Wait()
}
