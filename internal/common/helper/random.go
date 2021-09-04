package helper

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func Random() string {
	rand.Seed(time.Now().UnixNano())
	v := rand.Perm(9)
	return arrayToString(v, "")
}

func GenerateUsername() string {
	rand.Seed(time.Now().UnixNano())
	v := rand.Perm(10)
	return arrayToString(v, "")
}

func GeneratePassword() string {
	rand.Seed(time.Now().UnixNano())
	v := rand.Perm(15)
	return arrayToString(v, "")
}
func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}
