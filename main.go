package main

import (
	"strings",
	"fmt"
)

func main() {
	tokenString := "d48c948a-7361-4693-b86b-0e14145cdd6a"
	parts = strings.Split(tokenString, ".")
	fmt.Println(len(parts))
}
