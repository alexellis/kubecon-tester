package function

import (
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// Handle a serverless request
func Handle(req []byte) string {
	if os.Getenv("Http_Query") == "action=decode" {
		decoded := decode(req)
		return string(decoded)
	}

	res, _ := bcrypt.GenerateFromPassword(req, bcrypt.DefaultCost)

	return string(res)
}

func decode(req []byte) string {
	res := decodeValue(req)
	if res == true {
		return "true"
	}
	return "false"
}

func decodeValue(req []byte) bool {
	data := string(req)
	index := strings.Index(data, " ")
	if index == -1 {
		return false
	}

	plain := data[:index]
	hash := data[index+1:]
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain)); err != nil {
		return false
	}

	return true
}
