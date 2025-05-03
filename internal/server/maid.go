package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func queryMaid() ([]byte, error) {
	// visit https://nekosia.cat/documentation?page=api-endpoints for docs
	res, resErr := http.DefaultClient.Get("https://api.nekosia.cat/api/v1/images/maid")
	if resErr != nil {
		return nil, fmt.Errorf("error during dialing to maid API %v", resErr)
	}
	defer res.Body.Close()

	// I am so sorry for doing this horrible struct definition but this works ILY <3
	n := struct {
		Image struct {
			Original struct {
				Url string `json:"url"`
			} `json:"original"`
		} `json:"image"`
	}{}

	err := json.NewDecoder(res.Body).Decode(&n)

	if err!=nil{
		return nil,fmt.Errorf("error during decoding maid png %v",err)
	}

	return fetchMaid(n.Image.Original.Url)
}

func fetchMaid(url string)([]byte,error){
	res, resErr := http.DefaultClient.Get(url)
	if resErr != nil {
		return nil, fmt.Errorf("error during dialing to maid API %v", resErr)
	}
	defer res.Body.Close()

	data,err:=io.ReadAll(res.Body)
	if err!=nil{
		return nil,fmt.Errorf("error during decoing pic %v",err)
	}
	return data,err
}
