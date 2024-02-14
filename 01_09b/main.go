package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"text/tabwriter"
)

const path = "songs.json"

// Song stores all the song related information
type Song struct {
	Name      string `json:"name"`
	Album     string `json:"album"`
	PlayCount int64  `json:"play_count"`
}

// makePlaylist makes the merged sorted list of songs
func makePlaylist(albums [][]Song) []Song {
	songs := []Song{}
	iAlbum := map[string]int{} // index per album name
	nAlbum := map[string]int{} // remaining songs per album
	kAlbum := map[string]int{} // counter  per album name
	// number of songs per album
	// ini index per album
	for i := 0; i < len(albums); i++ {
		albumName := albums[i][0].Album
		nAlbum[albumName] = len(albums[i])
		iAlbum[albumName] = i
		kAlbum[albumName] = 0
	}
	for {
		// take the first song of every album
		s := []Song{}
		for a, i := range iAlbum {
			if nAlbum[a] == 0 { // no more songs left in this album
				continue
			}
			k := kAlbum[a] // next index available for this album
			s = append(s, albums[i][k])
		}
		// all songs from all albums have been picked up
		if len(s) == 0 {
			break
		}

		// sort the available songs according to their playcount
		sort.Slice(s, func(i, j int) bool {
			return s[i].PlayCount > s[j].PlayCount
		})
		// pick up the first one
		songs = append(songs, s[0])
		// update indexes
		nAlbum[s[0].Album] -= 1
		kAlbum[s[0].Album] += 1
	}
	return songs
}

func main() {
	albums := importData()
	printTable(makePlaylist(albums))
}

// printTable prints merged playlist as a table
func printTable(songs []Song) {
	w := tabwriter.NewWriter(os.Stdout, 3, 3, 3, ' ', tabwriter.TabIndent)
	fmt.Fprintln(w, "####\tSong\tAlbum\tPlay count")
	for i, s := range songs {
		fmt.Fprintf(w, "[%d]:\t%s\t%s\t%d\n", i+1, s.Name, s.Album, s.PlayCount)
	}
	w.Flush()

}

// importData reads the input data from file and creates the friends map
func importData() [][]Song {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var data [][]Song
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}
