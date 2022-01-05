package main

import (
	"crypto/sha256"
	"fmt"
	"time"
)


func main() {
	text := "Mantap Boss"
	salt := fmt.Sprintf("%d", time.Now().UnixNano())
	combine := fmt.Sprintf("%s %s", text, salt)
	fmt.Println(combine)
	sha := sha256.New()
	sha.Write([]byte(combine))
	encText := sha.Sum(nil)
	result := fmt.Sprintf("%x", encText)
	fmt.Println(result)
}