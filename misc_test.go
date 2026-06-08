package decimal

import (
	"testing"
)

func TestMisc(t *testing.T) {
	num := Int64(0)

	if num.EqualDec(Str("0")) != true {
		t.Fatal("should be equal")
	} else {
		t.Log("equal")
	}
}

func TestLongFloat(t *testing.T) {
	SetDefaultPrecision(3)
	t.Log(Float64(1.23456789))
}
