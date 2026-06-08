package main

import (
	"log"

	"github.com/xtdlib/decimal"
)

func main() {
	a := decimal.New(1.345)
	b := a
	log.Println(a.Equal(b))
	log.Println(b)
}
