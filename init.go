package main

import (
	"fmt"
	"strings"
	"regexp"
)


func main() {
	regex, _ := regexp.Compile(`([a-z]+)*([A-Z]+)*([1-9]+)*([@.]+)*`)
	kata2 := regex.FindAllString("!?>!)@_!>?@?>Odosa1?>2-o12.312PROowqm1o12dsa><>@**!?>?>>!", -1)
	saring := strings.Join(kata2, "")
	fmt.Println(saring)
	
}