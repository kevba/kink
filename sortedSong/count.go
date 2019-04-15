package main

import (
	"github.com/kevba/kink"
)

func NewCounter() *counter {
	return &counter{
		counters: map[string]*CountedSong{},
	}
}

type counter struct {
	counters map[string]*CountedSong
}

func (c *counter) countDay(d kink.Day) {
	for _, s := range d.Songs {
		countedSong, ok := c.counters[s.Title]

		if !ok {
			countedSong = &CountedSong{0, s}
			c.counters[s.Title] = countedSong
		}

		countedSong.count = countedSong.count + 1
	}
}

type CountedSong struct {
	count int
	kink.Song
}
