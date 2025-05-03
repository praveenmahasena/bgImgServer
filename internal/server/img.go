package server

import (
	"math/rand/v2"
)

func getImg() ([]byte, error) {
	randInt := rand.IntN(5)

	switch randInt {
	case 0:
		return queryNeko()
	case 1:
		return ipsumImg()
	case 2:
		return queryMaid()
	case 3:
		return queryPicRe()
	case 4:
		return queryWallhaven()
	}
	return nil, nil
}
