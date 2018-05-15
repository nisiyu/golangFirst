package main

import (
	"fmt"
	"github.com/nisiyu/golangFirst/stringutil"
)

func Reverse(x string) string {
	return stringutil.Reverse(x)
}

func main() {
    fmt.Println(Reverse("hello, world"))
}
