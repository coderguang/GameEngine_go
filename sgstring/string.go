package sgstring

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

var numRunes = []rune("0123456789")

var letterAndNumRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandStringAndNumRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterAndNumRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func RandNumStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = numRunes[rand.Intn(len(numRunes))]
	}
	return string(b)
}

func RandomFromList(params []string) string {
	index := rand.Intn(len(params))
	return params[index]
}

func ContainsWithAnd(targetStr string, subs []string) bool {
	for _, v := range subs {
		if !strings.Contains(targetStr, v) {
			return false
		}
	}
	return true
}

func ContainsWithOr(targetStr string, subs []string) bool {
	for _, v := range subs {
		if strings.Contains(targetStr, v) {
			return true
		}
	}
	return false
}

func EqualWithOr(targetStr string, subs []string) bool {
	for _, v := range subs {
		if strings.EqualFold(targetStr, v) {
			return true
		}
	}
	return false
}

func RemoveSpaceAndLineEnd(rawStr string) string {
	if strings.Contains(rawStr, "\r\n") {
		rawStr = strings.Replace(rawStr, "\r\n", "", -1)
	}
	if strings.Contains(rawStr, " ") {
		rawStr = strings.Replace(rawStr, " ", "", -1)
	}
	if strings.Contains(rawStr, "\n") {
		rawStr = strings.Replace(rawStr, "\n", "", -1)
	}
	if strings.Contains(rawStr, "\r") {
		rawStr = strings.Replace(rawStr, "\r", "", -1)
	}
	return rawStr
}
