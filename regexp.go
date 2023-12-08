package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regexp *regexp.Regexp = regexp.MustCompile(`a([a-z])n`)

	fmt.Println(regexp.MatchString("alwn"))
	fmt.Println(regexp.MatchString("aln"))
	fmt.Println(regexp.MatchString("agn"))

	fmt.Println(regexp.FindAllString("alin aan agan alban alann alw4n", 10))

}
