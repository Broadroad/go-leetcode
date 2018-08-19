package main

import (
	"fmt"
	"strings"
)

func uncommonFromSentences(A string, B string) []string {

	na := len(A)
	nb := len(B)

	var n int
	if na > nb {
		n = na
	} else {
		n = nb
	}
	flags := make(map[string]int, n)
	ret := make([]string, 0)
	As := strings.Split(A, " ")
	Bs := strings.Split(B, " ")

	for _, s := range As {
		flags[s] = flags[s] + 1
	}
	for _, s := range Bs {
		flags[s] = flags[s] + 1
	}

	for k, v := range flags {
		if v == 1 {
			ret = append(ret, k)
		}
	}
	return ret

}

func main() {
	a := "this apple is sweet"
	b := "this apple is sor"

	fmt.Println(uncommonFromSentences(a, b))
}
