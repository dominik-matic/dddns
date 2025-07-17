package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
)

func main() {
	tokenLen := 32
	maxTokenLen := 8192
	if len(os.Args) >= 2 {
		arg := os.Args[1]
		n, err := strconv.Atoi(arg)
		if err != nil || n > maxTokenLen {
			fmt.Println(arg, "is not a valid length number.")
		} else {
			tokenLen = n
		}
	}
	token := generateToken(tokenLen)
	fmt.Println(token)
}

func generateToken(nBytes int) string {
	bytes := make([]byte, nBytes)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}
