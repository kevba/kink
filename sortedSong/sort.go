package main

import (
	"sort"
)

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
