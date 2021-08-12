package adapters

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type IRandomAdapter interface {
	GenerateRandom() string
}

type RandomAdapter struct {
}

func NewRandomAdapter() IRandomAdapter {
	return &RandomAdapter{}
}

func (u *RandomAdapter) GenerateRandom() string {
	rand.Seed(time.Now().UnixNano())
	v := rand.Perm(9)
	return arrayToString(v, "")
}

func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}
