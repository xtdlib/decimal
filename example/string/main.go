package main

import (
	"log"

	"github.com/xtdlib/decimal"
)

func main() {
	log.Println("start")
	d := decimal.Float64(1.234567)
	log.Println(d)
	log.Println(d.StringFixed(3))
	log.Println(d.Trunc(3).StringFixed(3))

}
