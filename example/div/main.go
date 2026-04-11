package main

import (
	"log"

	"github.com/xtdlib/decimal"
)

func main() {
	// b2 := decimal.Str("0.000000001")
	// _ = b2
	// c := decimal.Str("1330.00000002")

	// log.Println(a.Add(b))
	// log.Println(a.Add(b).Add(c))
	// log.Println(a.Add(b).Add(c).Add(b))
	// log.Println(a.Add(b).Add(c).Add(b))
	// log.Println(a.Add(b).Add(c).Add(b).Div(decimal.Int64(2)))

	{
		v := decimal.Str("1234")
		for i := 0; i < 20; i++ {
			v = v.Mul(0.001)
			log.Println(v)
		}
	}

	{
		v := decimal.Str("1234")
		for i := 0; i < 20; i++ {
			v = v.DivInt(1000)
			log.Println(v)
		}
	}
}
