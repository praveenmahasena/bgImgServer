package server

import (
	"io"
	"net/http"
)


func ipsumImg()([]byte,error){
	req,reqErr:=http.NewRequest(http.MethodGet,"https://picsum.photos/1920/1080",nil)
	if reqErr!=nil{
		return nil,reqErr
	}

	res,resErr:=http.DefaultClient.Do(req)
	if resErr!=nil{
		return nil,reqErr
	}
	defer res.Body.Close()

	return io.ReadAll(res.Body)
}
