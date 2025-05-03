package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Neko struct {
	Url string `json:"url"`
}

type Nekos struct {
	Results []Neko `json:"results"`
}

func queryNeko() ([]byte, error) {
	req, reqErr := http.NewRequest(http.MethodGet, "https://nekos.best/api/v2/neko", nil)
	if reqErr != nil {
		return nil, fmt.Errorf("error during fetching neko pics %v", reqErr)
	}
	res, resErr := http.DefaultClient.Do(req)
	if resErr != nil {
		return nil, fmt.Errorf("error during fetching neko pics %v", reqErr)
	}
	n := &Nekos{[]Neko{}}
	err := json.NewDecoder(res.Body).Decode(n)
	if err != nil {
		return nil, fmt.Errorf("error during decoding src link %v", err)
	}
	return fetchNeko(n.Results[0].Url)
}

func fetchNeko(url string) ([]byte, error) {
	res, resErr := http.Get(url)
	if resErr != nil {
		return nil, fmt.Errorf("error during fetching neko boi UwU %v", resErr)
	}
	data, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, fmt.Errorf("error during decoding neko boi UwU %v", err)
	}
	return data, nil
}
