package main

import "encoding/json"

func fetchDay(dayOffset int, kink api) (day, error) {
	data, err := kink.fetchProgramming(dayOffset)
	if err != nil {
		return day{}, err
	}

	day, err := parseDay(data)
	return day, err
}

func parseDay(data []byte) (day, error) {
	dataJSON := programmingJSON{}
	err := json.Unmarshal(data, &dataJSON)
	if err != nil {
		return day{}, err
	}

	songs := parseSongs(dataJSON)

	return day{songs}, nil
}

func parseSongs(p programmingJSON) []song {
	songs := []song{}

	for _, s := range p.Included {
		if s.Type != "played-song" {
			continue
		}
		parsedSong := song{s.Attributes.Title, s.Attributes.Artist}
		songs = append(songs, parsedSong)
	}

	return songs
}

type day struct {
	songs []song "json:`included`"
}

type song struct {
	title  string
	artist string
}
