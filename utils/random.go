package utils

import (
	"fmt"
	"math/rand"
	"time"
)

const numset = "0123456789"
const alphaset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var seedRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

//RandomAlphanumeric generate random alphanumeric data
func RandomAlphanumeric(length int) string {
	charset := fmt.Sprintf("%s%s", alphaset, numset)
	randStr := make([]byte, length)
	for i := range randStr {
		randStr[i] = charset[seedRand.Intn(len(charset))]
	}
	return string(randStr)
}

//RandomAlpha generate random alphabet data
func RandomAlpha(length int) string {
	randStr := make([]byte, length)
	for i := range randStr {
		randStr[i] = alphaset[seedRand.Intn(len(alphaset))]
	}
	return string(randStr)
}

//RandomNumeric generate random numeric data
func RandomNumeric(length int) string {
	randStr := make([]byte, length)
	for i := range randStr {
		randStr[i] = numset[seedRand.Intn(len(numset))]
	}
	return string(randStr)
}
