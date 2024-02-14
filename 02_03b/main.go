package main

import (
	"log"
	"sync"
)

// the number of attendees we need to serve lunch to
const consumerCount = 300

// foodCourses represents the types of resources to pass to the consumers
var foodCourses = []string{
	"Caprese Salad",
	"Spaghetti Carbonara",
	"Vanilla Panna Cotta",
}

// takeLunch is the consumer function for the lunch simulation
// Change the signature of this function as required
func takeLunch(chCourses []chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	var wgConsumer sync.WaitGroup
	wgConsumer.Add(consumerCount)
	for i := 0; i < consumerCount; i++ {
		go func(i int) {
			defer wgConsumer.Done()
			// Every consumer wants to get a course
			for k := range foodCourses {
				<-chCourses[k]
			}
		}(i)
	}
	wgConsumer.Wait()
}

// serveLunch is the producer function for the lunch simulation.
// Change the signature of this function as required
func serveLunch(chCourses []chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for k := 0; k < consumerCount; k++ {
		var wgLunch sync.WaitGroup
		// Serve courses simultaneously
		wgLunch.Add(len(foodCourses))
		for i, course := range foodCourses {
			go func(i int, course string) {
				defer wgLunch.Done()
				chCourses[i] <- struct{}{}
				log.Printf("Served course %s.\n", course)
			}(i, course)
		}
		wgLunch.Wait()
	}
}

func main() {
	log.Printf("Welcome to the conference lunch! Serving %d attendees.\n", consumerCount)

	var chCourses = make([]chan struct{}, len(foodCourses))
	for i := range foodCourses {
		chCourses[i] = make(chan struct{})
	}
	var wg sync.WaitGroup
	wg.Add(2)
	go serveLunch(chCourses, &wg)
	go takeLunch(chCourses, &wg)
	wg.Wait()

	log.Println("END")
}
