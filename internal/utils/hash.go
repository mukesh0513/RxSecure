package utils

import "hash/fnv"

func GenerateHash(token string) int{
	h := fnv.New32a()
	h.Write([]byte(token))
	return int(h.Sum32()) % 1000
}
