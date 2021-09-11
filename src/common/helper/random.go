package helper

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// RandomString rand string from alphabet with length n.
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner generates a random owner name.
func RandomOwner() string {
	return RandomString(6)
}

// GenerateUsername for generate rand numbers with length 10.
func GenerateUsername() string {
	rand.Seed(time.Now().UnixNano())
	v := rand.Perm(10)

	return arrayToString(v, "")
}

// GeneratePassword for generate rand numbers with length 15.
func GeneratePassword() string {
	rand.Seed(time.Now().UnixNano())
	v := rand.Perm(15)

	return arrayToString(v, "")
}

func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}

// Hash for hash a string using bcrypt.
func Hash(s string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	return string(bytes)
}
