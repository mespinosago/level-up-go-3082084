package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

const path = "entries.json"

// raffleEntry is the struct we unmarshal raffle entries into
type raffleEntry struct {
	Id, Name string
}

// importData reads the raffle entries from file and creates the entries slice.
func importData() []raffleEntry {
	var result []raffleEntry
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(data, &result)
	if err != nil {
		fmt.Print(err)
	}
	return result
}

// getWinner returns a random winner from a slice of raffle entries.
func getWinner(entries []raffleEntry) raffleEntry {
	rand.Seed(time.Now().Unix())
	wi := rand.Intn(len(entries))
	return entries[wi]
}

func main() {
	entries := importData()
	log.Println("And... the raffle winning entry is...")
	winner := getWinner(entries)
	time.Sleep(500 * time.Millisecond)
	log.Println(winner)
}
