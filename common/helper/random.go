package helper

import (
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func Random() string {
	rand.Seed(time.Now().UnixNano())
	v := rand.Perm(9)
	return arrayToString(v, "")
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

func GenerateID() string {
	return uuid.NewString()
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

func Hash(s string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	return string(bytes)
}
