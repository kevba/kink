package kink

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const base string = "https://api.kink.nl"

type Api interface {
	fetchProgramming(day int) ([]byte, error)
}

func NewKinkAPI() Api {
	return &kinkAPI{}
}

type kinkAPI struct {
}

func (k *kinkAPI) fetchProgramming(day int) ([]byte, error) {
	url := fmt.Sprintf("%v/programming?includes=songs&day=%v", base, day)
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, nil
	}

	data, err := ioutil.ReadAll(resp.Body)
	return data, err
}

type programmingJSON struct {
	Included []struct {
		Type       string `json:"type"`
		Attributes struct {
			Title  string `json:"title"`
			Artist string `json:"artist"`
		} `json:"attributes"`
	} `json:"included"`
}
