package utils

import (
	"math/rand"
	"time"
)

const charset string = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
const idSize int = 32
const firstJan2014 = 1388534400

func GenerateToken() string {
	nanotime := time.Now().UTC().UnixNano()
	b62 := nanotimeToBaseN(nanotime, charset)
	rand.Seed(nanotime)
	random := rand.Int63()
	base64Rand := baseN(random, charset)

	if len(base64Rand) > idSize-len(b62) {
		base64Rand = base64Rand[(len(base64Rand) - idSize + len(b62)):]
	}
	id := b62 + base64Rand

	return id
}

func baseN(num int64, charset string) string {
	res := ""
	length := int64(len(charset))

	for {
		res = string(charset[num%length]) + res
		num = int64(num / length)
		if num == 0 {
			break
		}
	}

	return res
}

func nanotimeToBaseN(nanotime int64, charset string) string {
	// Timestmap of 1st Jan 2014!!
	// 1388534400
	epochTs := int64(firstJan2014 * 1000 * 1000 * 1000)
	nanotime -= epochTs
	bN := baseN(nanotime, charset)
	return bN
}
