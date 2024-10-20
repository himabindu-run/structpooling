package main

import "fmt"

const poolSize = 30
var messagePool = make(chan *Message, 300)
var counter int

func populatePool() {
	for i := 0; i < poolSize; i++ {
		messagePool <- &Message{}
		counter++
	}
	fmt.Println("Populating the pool with msg pointers")
}

func getMessageFromMessagePool() *Message {
	if len(messagePool) == 0 {
		messagePool <- &Message{}
		counter++
		fmt.Println("Creating new msg pointer", counter)
	}
	msgPointer := <- messagePool
	return msgPointer
}

func putBackMessage(m *Message) {
	messagePool <- m
	fmt.Println("Put the msg pointer back in pool channel")
}