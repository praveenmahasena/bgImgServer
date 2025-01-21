package server

import (
	"math/rand/v2"
)

func getImg()([]byte,error){
	randInt:=rand.IntN(3)

	switch randInt{
	case 0:
		return queryWallhaven()
	case 1:
		return ipsumImg()
	case 2:
	}

	return nil,nil

}
