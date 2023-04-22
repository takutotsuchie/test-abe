package main

import (
	"fmt"

	"github.com/takutotsuchie/test-abe/ABE"
)

func main() {
	a := ABE.AbeHiroshi()
	fmt.Println(a)
}
func AbeHiroshi() string {
	return "http://abehiroshi.la.coocan.jp/"
}
