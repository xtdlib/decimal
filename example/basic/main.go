package main

import (
	"log"

	"github.com/xtdlib/decimal"
)

func main() {
	d := decimal.Str("3.0000000")
	log.Println(d.Equal(3))

	x := decimal.Float64(0.1)
	y := decimal.Float64(0.2)
	log.Println(x.AddDec(y).Equal(0.3))
}
