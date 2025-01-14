package server

import (
	"io"
	"net/http"
)

func queryWallhaven() ([]byte, error) {
	req, reqErr := http.Get("https://wallhaven.cc/api/v1/search?categories=anime&purity=100&sorting=random&page=1")

	if reqErr != nil {
		return nil, reqErr
	}

	defer req.Body.Close()

	res, resErr := io.ReadAll(req.Body)

	if resErr != nil {
		return nil, resErr
	}

	leng := len(res)
	url := []byte{}

outer:
	for i := 0; i < leng; i += 4 {
		window := string(res[i : i+4])
		if window == "path" {
		inner:
			for idx := i + 7; idx < leng; idx += 1 {
				if res[idx] == '"' {
					break outer
				}
				if res[idx] == '\\' {
					continue inner
				}
				url = append(url, res[idx])
			}
			break outer
		}
	}

	imgReq,imgReqErr:=	http.Get(string(url))

	if imgReqErr!=nil{
		return nil,imgReqErr
	}
	defer imgReq.Body.Close()

	return io.ReadAll(imgReq.Body)
}
