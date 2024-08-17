package main

import (
	"fmt"

	"github.com/yanyiwu/gosimhash"
)

func main() {
	hasher := gosimhash.New()
	defer hasher.Free()
	s := "我来到北京清华大学"
	top_n := 5
	fingerprint := hasher.MakeSimhash(s, top_n)
	fmt.Printf("%s simhash: %x\n", s, fingerprint)
}
