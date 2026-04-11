package decimal

import (
	"encoding/json"
	"math/big"
	"testing"
)

func TestInt64ZeroAlloc(t *testing.T) {
	allocs := testing.AllocsPerRun(1000, func() {
		_ = Int64(42)
	})
	if allocs != 0 {
		t.Errorf("Int64(42): got %v allocs, want 0", allocs)
	}
}

func BenchmarkInt64(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = Int64(42)
	}
}

func BenchmarkAdd(b *testing.B) {
	a := Int64(100)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = a.Add(int64(200))
	}
}

func BenchmarkAddInt64(b *testing.B) {
	a := Int64(100)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = a.AddInt64(200)
	}
}

func BenchmarkAddInt(b *testing.B) {
	a := Int64(100)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = a.AddInt(200)
	}
}

func BenchmarkBigRatInt(b *testing.B) {
	x := big.NewInt(100)
	y := big.NewInt(200)
	b.ReportAllocs()
	b.ResetTimer()
	z := new(big.Int)

	for i := 0; i < b.N; i++ {
		_ = z.Add(x, y)
	}
}

func TestNew(t *testing.T) {
	if got := New(3).String(); got != "3" {
		t.Errorf("New(3) = %s, want 3", got)
	}
	if got := New(int32(3)).String(); got != "3" {
		t.Errorf("New(int32(3)) = %s, want 3", got)
	}
	if got := New(int16(3)).String(); got != "3" {
		t.Errorf("New(int16(3)) = %s, want 3", got)
	}
	if got := New(int64(3)).String(); got != "3" {
		t.Errorf("New(int64(3)) = %s, want 3", got)
	}
	if got := New(1.5).String(); got != "1.5" {
		t.Errorf("New(1.5) = %s, want 1.5", got)
	}
	if got := New(float32(1.5)).String(); got != "1.5" {
		t.Errorf("New(float32(1.5)) = %s, want 1.5", got)
	}
	if got := New("1.5").String(); got != "1.5" {
		t.Errorf(`New("1.5") = %s, want 1.5`, got)
	}
	if got := New(Int64(3)).String(); got != "3" {
		t.Errorf("New(Decimal) = %s, want 3", got)
	}
}

func TestInt64Func(t *testing.T) {
	if got := Int64(42).String(); got != "42" {
		t.Errorf("Int64(42) = %s, want 42", got)
	}
}

func TestStrFunc(t *testing.T) {
	if got := Str("3.14").String(); got != "3.14" {
		t.Errorf(`Str("3.14") = %s, want 3.14`, got)
	}
}

func TestFloat64Func(t *testing.T) {
	if got := Float64(1.5).String(); got != "1.5" {
		t.Errorf("Float64(1.5) = %s, want 1.5", got)
	}
}

func TestAdd(t *testing.T) {
	got := New(1).Add(2).String()
	if got != "3" {
		t.Errorf("1+2 = %s, want 3", got)
	}
}

func TestAddInt64(t *testing.T) {
	got := New(1).AddInt64(2).String()
	if got != "3" {
		t.Errorf("1+2 = %s, want 3", got)
	}
}

func TestAddInt(t *testing.T) {
	got := New(1).AddInt(2).String()
	if got != "3" {
		t.Errorf("1+2 = %s, want 3", got)
	}
}

func TestAddStr(t *testing.T) {
	got := New(1).AddStr("2").String()
	if got != "3" {
		t.Errorf("1+2 = %s, want 3", got)
	}
}

func TestAddDec(t *testing.T) {
	got := New(1).AddDec(Int64(2)).String()
	if got != "3" {
		t.Errorf("1+2 = %s, want 3", got)
	}
}

func TestSub(t *testing.T) {
	got := New(5).Sub(3).String()
	if got != "2" {
		t.Errorf("5-3 = %s, want 2", got)
	}
}

func TestSubInt64(t *testing.T) {
	got := New(5).SubInt64(3).String()
	if got != "2" {
		t.Errorf("5-3 = %s, want 2", got)
	}
}

func TestSubInt(t *testing.T) {
	got := New(5).SubInt(3).String()
	if got != "2" {
		t.Errorf("5-3 = %s, want 2", got)
	}
}

func TestSubStr(t *testing.T) {
	got := New(5).SubStr("3").String()
	if got != "2" {
		t.Errorf("5-3 = %s, want 2", got)
	}
}

func TestSubDec(t *testing.T) {
	got := New(5).SubDec(Int64(3)).String()
	if got != "2" {
		t.Errorf("5-3 = %s, want 2", got)
	}
}

func TestMul(t *testing.T) {
	got := New(3).Mul(4).String()
	if got != "12" {
		t.Errorf("3*4 = %s, want 12", got)
	}
}

func TestMulInt64(t *testing.T) {
	got := New(3).MulInt64(4).String()
	if got != "12" {
		t.Errorf("3*4 = %s, want 12", got)
	}
}

func TestMulInt(t *testing.T) {
	got := New(3).MulInt(4).String()
	if got != "12" {
		t.Errorf("3*4 = %s, want 12", got)
	}
}

func TestMulStr(t *testing.T) {
	got := New(3).MulStr("4").String()
	if got != "12" {
		t.Errorf("3*4 = %s, want 12", got)
	}
}

func TestMulDec(t *testing.T) {
	got := New(3).MulDec(Int64(4)).String()
	if got != "12" {
		t.Errorf("3*4 = %s, want 12", got)
	}
}

func TestDiv(t *testing.T) {
	got := New(3).Div(2).String()
	if got != "1.5" {
		t.Errorf("3/2 = %s, want 1.5", got)
	}
}

func TestDivInt64(t *testing.T) {
	got := New(3).DivInt64(2).String()
	if got != "1.5" {
		t.Errorf("3/2 = %s, want 1.5", got)
	}
}

func TestDivInt(t *testing.T) {
	got := New(3).DivInt(2).String()
	if got != "1.5" {
		t.Errorf("3/2 = %s, want 1.5", got)
	}
}

func TestDivStr(t *testing.T) {
	got := New(3).DivStr("2").String()
	if got != "1.5" {
		t.Errorf("3/2 = %s, want 1.5", got)
	}
}

func TestDivDec(t *testing.T) {
	got := New(3).DivDec(Int64(2)).String()
	if got != "1.5" {
		t.Errorf("3/2 = %s, want 1.5", got)
	}
}

func TestTryDiv(t *testing.T) {
	got, err := New(3).TryDiv(2)
	if err != nil {
		t.Fatal(err)
	}
	if got.String() != "1.5" {
		t.Errorf("3/2 = %s, want 1.5", got.String())
	}
	_, err = New(3).TryDiv(0)
	if err == nil {
		t.Error("3/0 should error")
	}
}

func TestTryDivDec(t *testing.T) {
	got, err := New(3).TryDivDec(Int64(2))
	if err != nil {
		t.Fatal(err)
	}
	if got.String() != "1.5" {
		t.Errorf("3/2 = %s, want 1.5", got.String())
	}
	_, err = New(3).TryDivDec(Int64(0))
	if err == nil {
		t.Error("3/0 should error")
	}
}

func TestDecimalInt64(t *testing.T) {
	got := New(42).Int64()
	if got != 42 {
		t.Errorf("Int64() = %d, want 42", got)
	}
}

func TestDecimalFloat64(t *testing.T) {
	got := New("1.5").Float64()
	if got != 1.5 {
		t.Errorf("Float64() = %f, want 1.5", got)
	}
}

func TestEqual(t *testing.T) {
	if !New(3).Equal(3) {
		t.Error("3 == 3 should be true")
	}
	if New(3).Equal(4) {
		t.Error("3 == 4 should be false")
	}
}

func TestEqualInt64(t *testing.T) {
	if !New(3).EqualInt64(3) {
		t.Error("3 == 3 should be true")
	}
}

func TestEqualInt(t *testing.T) {
	if !New(3).EqualInt(3) {
		t.Error("3 == 3 should be true")
	}
}

func TestEqualStr(t *testing.T) {
	if !New(3).EqualStr("3") {
		t.Error("3 == 3 should be true")
	}
}

func TestEqualDec(t *testing.T) {
	if !New(3).EqualDec(Int64(3)) {
		t.Error("3 == 3 should be true")
	}
}

func TestGreaterThan(t *testing.T) {
	if !New(5).GreaterThan(3) {
		t.Error("5 > 3 should be true")
	}
	if New(3).GreaterThan(5) {
		t.Error("3 > 5 should be false")
	}
}

func TestGreaterThanInt64(t *testing.T) {
	if !New(5).GreaterThanInt64(3) {
		t.Error("5 > 3 should be true")
	}
}

func TestGreaterThanInt(t *testing.T) {
	if !New(5).GreaterThanInt(3) {
		t.Error("5 > 3 should be true")
	}
}

func TestGreaterThanStr(t *testing.T) {
	if !New(5).GreaterThanStr("3") {
		t.Error("5 > 3 should be true")
	}
}

func TestGreaterThanDec(t *testing.T) {
	if !New(5).GreaterThanDec(Int64(3)) {
		t.Error("5 > 3 should be true")
	}
}

func TestLessThan(t *testing.T) {
	if !New(3).LessThan(5) {
		t.Error("3 < 5 should be true")
	}
	if New(5).LessThan(3) {
		t.Error("5 < 3 should be false")
	}
}

func TestLessThanInt64(t *testing.T) {
	if !New(3).LessThanInt64(5) {
		t.Error("3 < 5 should be true")
	}
}

func TestLessThanInt(t *testing.T) {
	if !New(3).LessThanInt(5) {
		t.Error("3 < 5 should be true")
	}
}

func TestLessThanStr(t *testing.T) {
	if !New(3).LessThanStr("5") {
		t.Error("3 < 5 should be true")
	}
}

func TestLessThanDec(t *testing.T) {
	if !New(3).LessThanDec(Int64(5)) {
		t.Error("3 < 5 should be true")
	}
}

func TestGreaterThanOrEqual(t *testing.T) {
	if !New(5).GreaterThanOrEqual(3) {
		t.Error("5 >= 3 should be true")
	}
	if !New(3).GreaterThanOrEqual(3) {
		t.Error("3 >= 3 should be true")
	}
}

func TestGreaterThanOrEqualInt64(t *testing.T) {
	if !New(5).GreaterThanOrEqualInt64(5) {
		t.Error("5 >= 5 should be true")
	}
}

func TestGreaterThanOrEqualInt(t *testing.T) {
	if !New(5).GreaterThanOrEqualInt(5) {
		t.Error("5 >= 5 should be true")
	}
}

func TestGreaterThanOrEqualStr(t *testing.T) {
	if !New(5).GreaterThanOrEqualStr("5") {
		t.Error("5 >= 5 should be true")
	}
}

func TestGreaterThanOrEqualDec(t *testing.T) {
	if !New(5).GreaterThanOrEqualDec(Int64(5)) {
		t.Error("5 >= 5 should be true")
	}
}

func TestLessThanOrEqual(t *testing.T) {
	if !New(3).LessThanOrEqual(5) {
		t.Error("3 <= 5 should be true")
	}
	if !New(3).LessThanOrEqual(3) {
		t.Error("3 <= 3 should be true")
	}
}

func TestLessThanOrEqualInt64(t *testing.T) {
	if !New(3).LessThanOrEqualInt64(3) {
		t.Error("3 <= 3 should be true")
	}
}

func TestLessThanOrEqualInt(t *testing.T) {
	if !New(3).LessThanOrEqualInt(3) {
		t.Error("3 <= 3 should be true")
	}
}

func TestLessThanOrEqualStr(t *testing.T) {
	if !New(3).LessThanOrEqualStr("3") {
		t.Error("3 <= 3 should be true")
	}
}

func TestLessThanOrEqualDec(t *testing.T) {
	if !New(3).LessThanOrEqualDec(Int64(3)) {
		t.Error("3 <= 3 should be true")
	}
}

func TestNeg(t *testing.T) {
	got := New(5).Neg().String()
	if got != "-5" {
		t.Errorf("Neg(5) = %s, want -5", got)
	}
}

func TestAbs(t *testing.T) {
	got := New(-5).Abs().String()
	if got != "5" {
		t.Errorf("Abs(-5) = %s, want 5", got)
	}
}

// func TestRound(t *testing.T) {
// 	got := New("1.5555").Round(4).String()
// 	if got != "1.556" {
// 		t.Logf("Round(1.5555, 4) = %s", got)
// 	}
// }
//
// func TestRoundDown(t *testing.T) {
// 	got := New("1.559").RoundDown(2).String()
// 	if got != "1.55" {
// 		t.Errorf("RoundDown(1.559, 2) = %s, want 1.55", got)
// 	}
// }

func TestRoundUp(t *testing.T) {
	got := New("1.551").RoundUp(2).String()
	if got != "1.56" {
		t.Errorf("RoundUp(1.551, 2) = %s, want 1.56", got)
	}
}

func TestString(t *testing.T) {
	got := New(42).String()
	if got != "42" {
		t.Errorf("String() = %s, want 42", got)
	}
}

func TestMarshalJSON(t *testing.T) {
	type S struct {
		Price Decimal `json:"price"`
	}
	s := S{Price: New("1.5")}
	b, err := json.Marshal(s)
	if err != nil {
		t.Fatal(err)
	}
	if string(b) != `{"price":"1.5"}` {
		t.Errorf("MarshalJSON = %s, want %s", string(b), `{"price":"1.5"}`)
	}
}

func TestUnmarshalJSON(t *testing.T) {
	type S struct {
		Price Decimal `json:"price"`
	}
	var s S
	if err := json.Unmarshal([]byte(`{"price":"3.14"}`), &s); err != nil {
		t.Fatal(err)
	}
	if s.Price.String() != "3.14" {
		t.Errorf("UnmarshalJSON = %s, want 3.14", s.Price.String())
	}
}

func TestScan(t *testing.T) {
	var d Decimal
	if err := d.Scan("1.5"); err != nil {
		t.Fatal(err)
	}
	if d.String() != "1.5" {
		t.Errorf("Scan = %s, want 1.5", d.String())
	}
}

func TestValue(t *testing.T) {
	d := New("1.5")
	v, err := d.Value()
	if err != nil {
		t.Fatal(err)
	}
	if v != "1.5" {
		t.Errorf("Value = %v, want 1.5", v)
	}
}

func TestNewFromBigInt(t *testing.T) {
	// 314 * 10^-2 = 3.14
	got := NewFromBigInt(big.NewInt(314), -2).String()
	if got != "3.14" {
		t.Errorf("NewFromBigInt(314, -2) = %s, want 3.14", got)
	}
	// 5 * 10^3 = 5000
	got = NewFromBigInt(big.NewInt(5), 3).String()
	if got != "5000" {
		t.Errorf("NewFromBigInt(5, 3) = %s, want 5000", got)
	}
}

func TestPremium(t *testing.T) {
	kimchi := New(1521).Premium(1509)
	t.Log(kimchi)
}

// func BenchmarkAddDecimal(b *testing.B) {
// 	a := Int64(100)
// 	e := Int64(200)
// 	b.ReportAllocs()
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		out := a.Add(e)
// 		if !out.Equal(Int64(300)) {
// 			b.Fatalf("unexpected result: got %s, want 300", out.String())
// 		}
// 	}
// }
//
// func BenchmarkAdd2(b *testing.B) {
// 	a := Int64(100)
// 	e := Int64(200)
// 	b.ReportAllocs()
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		out := a.Add(e)
// 		if !out.Equal(Int64(300)) {
// 			b.Fatalf("unexpected result: got %s, want 300", out.String())
// 		}
// 	}
// }
