package main

import (
	"fmt"
	"log"
	"sync"
)

// setup constants
const baristaCount = 3
const customerCount = 20
const maxOrderCount = 40

// the total amount of drinks that the bartenders have made
type coffeeShop struct {
	orderCount int

	orderCoffee  chan struct{}
	finishCoffee chan struct{}

	mx sync.RWMutex
}

// registerOrder ensures that the order made by the baristas is counted
func (p *coffeeShop) registerOrder() {
	defer p.mx.Unlock()
	p.mx.Lock()
	p.orderCount++
}

func (p *coffeeShop) timeToClose() bool {
	defer p.mx.RUnlock()
	p.mx.RLock()
	if p.orderCount >= maxOrderCount {
		return true
	}
	return false
}

// barista is the resource producer of the coffee shop
func (p *coffeeShop) barista(name string) {
	for {
		log.Println(p.orderCount)
		if p.timeToClose() {
			return
		}
		select {
		case <-p.orderCoffee:
			p.registerOrder()
			log.Printf("%s makes a coffee.\n", name)
			p.finishCoffee <- struct{}{}
		}
	}
}

// customer is the resource consumer of the coffee shop
func (p *coffeeShop) customer(name string) {
	for {
		if p.timeToClose() {
			return
		}
		select {
		case p.orderCoffee <- struct{}{}:
			log.Printf("%s orders a coffee!", name)
			<-p.finishCoffee
			log.Printf("%s enjoys a coffee!\n", name)
		}
	}
}

func main() {
	log.Println("Welcome to the Level Up Go coffee shop!")
	orderCoffee := make(chan struct{}, baristaCount)
	finishCoffee := make(chan struct{}, baristaCount)
	p := coffeeShop{
		orderCoffee:  orderCoffee,
		finishCoffee: finishCoffee,
	}
	var wg sync.WaitGroup
	wg.Add(baristaCount + customerCount)
	for i := 0; i < baristaCount; i++ {
		go func(i int) {
			defer wg.Done()
			p.barista(fmt.Sprint("Barista-", i))
		}(i)
	}
	for i := 0; i < customerCount; i++ {
		go func(i int) {
			defer wg.Done()
			go p.customer(fmt.Sprint("Customer-", i))
		}(i)
	}
	wg.Wait()
	log.Println("The Level Up Go coffee shop has closed! Bye!")
}
