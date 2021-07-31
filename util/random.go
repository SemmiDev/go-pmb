package util

import (
	"fmt"
	"strings"
	"time"
)

func RandomVirtualAccount(phone string) string {
	time := time.Now().UnixNano()
	split := strings.Split(fmt.Sprint(time), "")
	va := strings.Join(split, "") + phone
	return va
}
