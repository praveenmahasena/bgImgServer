package server

import (
	"fmt"
	"io"
	"net/http"
)

func queryPicRe()([]byte,error){
	res,resErr:=http.Get("https://pic.re/image")
	if resErr!=nil{
		return nil,fmt.Errorf("error during fetching up pic re %v",resErr)
	}
	defer res.Body.Close()
	data,err:=io.ReadAll(res.Body)
	if err!=nil{
		return nil,fmt.Errorf("error during picking up data from res body %v",err)
	}
	return data,nil
}
