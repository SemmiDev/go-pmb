package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func Random() (r string) {
	rand.Seed(time.Now().UnixNano())
	v := rand.Perm(9)
	return arrayToString(v, "")
}

func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}
