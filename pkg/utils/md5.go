package utils

import (
	"crypto/md5"
	"fmt"
)

func Hashing(key string) string {

	data := []byte(key)
	hashed := fmt.Sprintf("%x", md5.Sum(data))
	return hashed
}
