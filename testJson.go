package main

import (
	"encoding/json"
	"fmt"
)

type Token struct {
	Token string
}

func main() {

	data := Token {
		Token: "123",
	}

	json2, _ := json.Marshal(data)
}