package util

import (
	random "crypto/rand"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func NIMAndEmailGenerator(fullName string, path uint, year uint32) (nim, email string) {

	// YEAR   PATH   RAND
	// 200    311    3948
	// 170    312    3941
	// 170    313    3943

	// YEAR
	yearInString := strconv.Itoa(int(year))
	y := string(yearInString[2]) + string(yearInString[3]) + "0"

	// path
	var p string
	switch path {
	case 1:
		p = "312"
	case 2:
		p = "311"
	case 3:
		p = "313"
	}

	// rand
	r := fmt.Sprintf(EncodeToString(4))

	nim = y + p + r

	validName := strings.ToLower(fullName)
	splitFullName := strings.Split(validName, " ")
	if len(splitFullName) == 1 {
		email = splitFullName[0] + r
	}else{
		email = splitFullName[0] + "_" + splitFullName[1] + r
	}

	email += "@student.unri.ac.id"
	return
}

func EncodeToString(max int) string {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(random.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}