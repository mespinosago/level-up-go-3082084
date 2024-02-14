package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

// the amount of bidders we have at our auction
const bidderCount = 10

// initial wallet value for all bidders
const walletAmount = 250

// items is the map of auction items
var items = []string{
	"The \"Best Gopher\" trophy",
	"The \"Learn Go with Adelina\" experience",
	"Two tickets to a Go conference",
	"Signed copy of \"Beautiful Go code\"",
	"Vintage Gopher plushie",
}

// bid is a type that pairs the bidder id and the amount they want to bid
type bid struct {
	bidderID string
	amount   int
}

// auctioneer receives bids and announces winners
type auctioneer struct {
	bidders map[string]*bidder
	winners map[string]string //winner (bidderID per items)
}

// runAuction and manages the auction for all the items to be sold
// Change the signature of this function as required
func (a *auctioneer) runAuction() {
	for _, item := range items {
		log.Printf("Opening bids for %s!\n", item)
		maxBid := 0
		winner := ""
		for id, bidder := range a.bidders {
			bd := <-bidder.bid
			if bd.amount > maxBid {
				maxBid = bd.amount
				winner = id
			}
		}
		a.winners[item] = winner
		a.bidders[winner].payBid(maxBid)
		log.Printf("Winner for %s is %s, BID: %d\n", item, winner, maxBid)
	}
}

// bidder is a type that holds the bidder id and wallet
type bidder struct {
	id     string
	wallet int
	bid    chan bid
}

// placeBid generates a random amount and places it on the bids channels
// Change the signature of this function as required
func (b *bidder) placeBid() {
	for _, item := range items {
		log.Printf("Bidder %s, bidding for %s!\n", b.id, item)
		bidAmound := 0
		if b.wallet > 0 {
			bidAmound = rand.Intn(b.wallet + 1)
		}
		b.bid <- bid{
			bidderID: b.id,
			amount:   bidAmound,
		}
	}
}

// payBid subtracts the bid amount from the wallet of the auction winner
func (b *bidder) payBid(amount int) {
	b.wallet -= amount
}

func main() {
	rand.Seed(time.Now().UnixNano())
	log.Println("Welcome to the LinkedIn Learning auction.")
	bidders := make(map[string]*bidder, bidderCount)
	for i := 0; i < bidderCount; i++ {
		id := fmt.Sprint("Bidder ", i)
		b := bidder{
			id:     id,
			wallet: walletAmount,
			bid:    make(chan bid),
		}
		bidders[id] = &b
		go b.placeBid()
	}
	a := auctioneer{
		bidders: bidders,
		winners: map[string]string{},
	}
	a.runAuction()
	log.Println("WINNERS", a.winners)
	for item, b := range a.bidders {
		log.Println("item ", item, "  BIDDER ", b.id, "  WALLET ", b.wallet)
	}
}

// getRandomAmount generates a random integer amount up to max
func getRandomAmount(max int) int {
	return rand.Intn(int(max))
}
