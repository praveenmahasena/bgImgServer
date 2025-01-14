package server

import (
	"math/rand/v2"
)

func getImg()([]byte,error){
	randInt:=rand.IntN(1)

	switch randInt{
	case 0:
		return queryWallhaven()
	}

	return nil,nil

}
