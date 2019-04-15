package main

import (
	"flag"
	"log"
	"os"
	"github.com/kevba/kink"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	daysTotal := flag.Int("days", 7, "days to look back")
	flag.Parse()

	api := kink.NewKinkAPI()
	c := NewCounter()

	for i := 0; i <= *daysTotal; i++ {

		d, err := kink.FetchDay(0-i, api)
		if err != nil {
			log.Fatal(err)
		}

		c.countDay(d)

	}

	for _, cs := range sortSongs(c) {
		log.Printf("%v - %v: %v", cs.Artist, cs.Title, cs.count)
	}
}
