package main

import (
	"testing"

	"github.com/cockroachdb/apd"
	"github.com/xtdlib/decimal"
)

// xtdlib/decimal benchmarks

func BenchmarkXtdAdd(b *testing.B) {
	x := decimal.New("1.23")
	y := decimal.New("4.56")
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = x.AddDec(y)
	}
}

func BenchmarkXtdMul(b *testing.B) {
	x := decimal.New("1.23")
	y := decimal.New("4.56")
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = x.MulDec(y)
	}
}

func BenchmarkXtdDiv(b *testing.B) {
	x := decimal.New("1.23")
	y := decimal.New("4.56")
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = x.DivDec(y)
	}
}

// cockroachdb/apd benchmarks

func BenchmarkApdAdd(b *testing.B) {
	x, _, _ := apd.NewFromString("1.23")
	y, _, _ := apd.NewFromString("4.56")
	ctx := apd.BaseContext
	out := new(apd.Decimal)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ctx.Add(out, x, y)
	}
}

func BenchmarkApdMul(b *testing.B) {
	x, _, _ := apd.NewFromString("1.23")
	y, _, _ := apd.NewFromString("4.56")
	ctx := apd.BaseContext
	out := new(apd.Decimal)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ctx.Mul(out, x, y)
	}
}

func BenchmarkApdDiv(b *testing.B) {
	x, _, _ := apd.NewFromString("1.23")
	y, _, _ := apd.NewFromString("4.56")
	ctx := apd.BaseContext
	out := new(apd.Decimal)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ctx.Quo(out, x, y)
	}
}

// int -> decimal conversion benchmarks

func BenchmarkXtdFromIntNew(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = decimal.New(12345)
	}
}

func BenchmarkXtdFromIntInt64(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = decimal.Int64(12345)
	}
}

func BenchmarkApdFromInt(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = apd.New(12345, 0)
	}
}
