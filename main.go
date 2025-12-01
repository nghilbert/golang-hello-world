package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numberProducer(channel chan int) {
    for { // Infinite loop
        num :=  rand.Intn(1000) // Generate an int (0-999)
        channel <- num // Waits for consumer
        time.Sleep(5 * time.Second) // Wait 5 seconds
    }
}

func main() {
    channel := make(chan int) // Create an unbuffered channel
    go numberProducer(channel) // launches a Goroutine
    
    // Consumer runs at the same time as producer in main
    for { // Infinite loop
         num := <- channel // Waits for producer
         fmt.Println("Random number from producer:", num)
     }
}