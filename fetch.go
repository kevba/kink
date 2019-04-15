package kink

import "encoding/json"

func FetchDay(dayOffset int, kink Api) (Day, error) {
	data, err := kink.fetchProgramming(dayOffset)
	if err != nil {
		return Day{}, err
	}

	day, err := parseDay(data)
	return day, err
}

func parseDay(data []byte) (Day, error) {
	dataJSON := programmingJSON{}
	err := json.Unmarshal(data, &dataJSON)
	if err != nil {
		return Day{}, err
	}

	songs := parseSongs(dataJSON)

	return Day{songs}, nil
}

func parseSongs(p programmingJSON) []Song {
	songs := []Song{}

	for _, s := range p.Included {
		if s.Type != "played-song" {
			continue
		}
		parsedSong := Song{s.Attributes.Title, s.Attributes.Artist}
		songs = append(songs, parsedSong)
	}

	return songs
}

type Day struct {
	Songs []Song
}

type Song struct {
	Title  string
	Artist string
}
