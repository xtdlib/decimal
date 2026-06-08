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
		decimal.SetDefaultPrecision(4)
		v := decimal.Str("0.000000000333333")
		log.Println(v.Div(3))
	}
	//
	// {
	// 	v := decimal.Str("-1.15")
	// 	log.Println(v.RoundHAZ(1))
	// }
}
