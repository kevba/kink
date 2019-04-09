package main

import (
	"flag"
	"log"
	"os"
	"sort"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	daysTotal := flag.Int("days", 7, "days to look back")
	flag.Parse()

	api := NewKinkAPI()
	c := NewCounter()

	for i := 0; i <= *daysTotal; i++ {

		d, err := fetchDay(0-i, api)
		if err != nil {
			log.Fatal(err)
		}

		c.countDay(d)

	}

	for _, cs := range sortSongs(c) {
		log.Printf("%v - %v: %v", cs.artist, cs.title, cs.count)
	}
}

func NewCounter() *counter {
	return &counter{
		counters: map[string]*CountedSong{},
	}
}

type counter struct {
	counters map[string]*CountedSong
}

func (c *counter) countDay(d day) {
	for _, s := range d.songs {
		countedSong, ok := c.counters[s.title]

		if !ok {
			countedSong = &CountedSong{0, s}
			c.counters[s.title] = countedSong
		}

		countedSong.count = countedSong.count + 1
	}
}

type CountedSong struct {
	count int
	song
}

func sortSongs(c *counter) []*CountedSong {
	songList := []*CountedSong{}

	for _, countedSong := range c.counters {
		songList = append(songList, countedSong)
	}

	sort.Sort(SongSorter(songList))

	return songList
}

type SongSorter []*CountedSong

func (s SongSorter) Len() int           { return len(s) }
func (s SongSorter) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s SongSorter) Less(i, j int) bool { return s[i].count < s[j].count }
