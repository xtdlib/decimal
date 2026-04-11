package decimal

import (
	"database/sql/driver"
	"fmt"
	"math"
	"math/big"
	"strings"

	"github.com/quagmt/udecimal"
)

var defaultPrec uint8 = 19

type Decimal struct {
	udecimal.Decimal
}

func (d Decimal) Trunc(prec uint8) Decimal {
	return Decimal{
		Decimal: d.Decimal.Trunc(prec),
	}
}

func (d Decimal) RoundUp(prec uint8) Decimal {
	return Decimal{
		Decimal: d.RoundAwayFromZero(prec),
	}
}

func (d Decimal) String() string {
	return d.Decimal.String()
}

// StringFixed returns a string representation of the decimal with a fixed number of decimal places.
func (d Decimal) StringFixed(prec uint8) string {
	return d.Decimal.Trunc(prec).StringFixed(prec)
}

func (d Decimal) MarshalJSON() ([]byte, error) {
	return d.Decimal.MarshalJSON()
}

func (d *Decimal) UnmarshalJSON(data []byte) error {
	return d.Decimal.UnmarshalJSON(data)
}

func (d Decimal) MarshalText() ([]byte, error) {
	return d.Decimal.MarshalText()
}

func (d *Decimal) UnmarshalText(data []byte) error {
	return d.Decimal.UnmarshalText(data)
}

func (d *Decimal) Scan(src any) error {
	return d.Decimal.Scan(src)
}

func (d Decimal) Value() (driver.Value, error) {
	return d.Decimal.Value()
}

func (d Decimal) ToHiLo() (neg bool, hi uint64, lo uint64, prec uint8, ok bool) {
	return d.Decimal.ToHiLo()
}

// BigIntExpToString converts a big.Int coefficient and exponent to a decimal string.
func BigIntExpToString(coef *big.Int, exp int32) string {
	s := coef.String()
	e := int(exp)

	if e >= 0 {
		return s + strings.Repeat("0", e)
	}

	neg := ""
	digits := s
	if digits[0] == '-' {
		neg = "-"
		digits = digits[1:]
	}
	pos := len(digits) + e
	if pos <= 0 {
		return neg + "0." + strings.Repeat("0", -pos) + digits
	}
	return neg + digits[:pos] + "." + digits[pos:]
}

func (d Decimal) Add(e any) Decimal {
	return Decimal{
		Decimal: d.Decimal.Add(New(e).Decimal),
	}
}

func (d Decimal) AddInt64(e int64) Decimal {
	return Decimal{
		Decimal: d.Decimal.Add(Int64(e).Decimal),
	}
}

func (d Decimal) AddInt(e int) Decimal {
	return Decimal{
		Decimal: d.Decimal.Add(Int64(int64(e)).Decimal),
	}
}

func (d Decimal) AddStr(e string) Decimal {
	return Decimal{
		Decimal: d.Decimal.Add(Str(e).Decimal),
	}
}

func (d Decimal) AddDec(e Decimal) Decimal {
	return Decimal{
		Decimal: d.Decimal.Add(e.Decimal),
	}
}

func (d Decimal) Sub(e any) Decimal {
	return Decimal{
		Decimal: d.Decimal.Sub(New(e).Decimal),
	}
}

func (d Decimal) SubInt64(e int64) Decimal {
	return Decimal{
		Decimal: d.Decimal.Sub(Int64(e).Decimal),
	}
}

func (d Decimal) SubInt(e int) Decimal {
	return Decimal{
		Decimal: d.Decimal.Sub(Int64(int64(e)).Decimal),
	}
}

func (d Decimal) SubStr(e string) Decimal {
	return Decimal{
		Decimal: d.Decimal.Sub(Str(e).Decimal),
	}
}

func (d Decimal) SubDec(e Decimal) Decimal {
	return Decimal{
		Decimal: d.Decimal.Sub(e.Decimal),
	}
}

func (d Decimal) Mul(e any) Decimal {
	return Decimal{
		Decimal: d.Decimal.Mul(New(e).Decimal),
	}
}

func (d Decimal) MulInt64(e int64) Decimal {
	return Decimal{
		Decimal: d.Decimal.Mul(Int64(e).Decimal),
	}
}

func (d Decimal) MulInt(e int) Decimal {
	return Decimal{
		Decimal: d.Decimal.Mul(Int64(int64(e)).Decimal),
	}
}

func (d Decimal) MulStr(e string) Decimal {
	return Decimal{
		Decimal: d.Decimal.Mul(Str(e).Decimal),
	}
}

func (d Decimal) MulDec(e Decimal) Decimal {
	return Decimal{
		Decimal: d.Decimal.Mul(e.Decimal),
	}
}

func (d Decimal) Div(e any) Decimal {
	return Decimal{
		Decimal: d.Decimal.MustDiv(New(e).Decimal),
	}
}

func (d Decimal) DivInt64(e int64) Decimal {
	return Decimal{
		Decimal: d.Decimal.MustDiv(Int64(e).Decimal),
	}
}

func (d Decimal) DivInt(e int) Decimal {
	return Decimal{
		Decimal: d.Decimal.MustDiv(Int64(int64(e)).Decimal),
	}
}

func (d Decimal) DivStr(e string) Decimal {
	return Decimal{
		Decimal: d.Decimal.MustDiv(Str(e).Decimal),
	}
}

func (d Decimal) DivDec(e Decimal) Decimal {
	return Decimal{
		Decimal: d.Decimal.MustDiv(e.Decimal),
	}
}

func (d Decimal) TryDiv(e any) (Decimal, error) {
	v, err := d.Decimal.Div(New(e).Decimal)
	if err != nil {
		return Decimal{}, err
	}
	return Decimal{
		Decimal: v,
	}, nil
}

func (d Decimal) TryDivDec(e Decimal) (Decimal, error) {
	v, err := d.Decimal.Div(e.Decimal)
	if err != nil {
		return Decimal{}, err
	}
	return Decimal{
		Decimal: v,
	}, nil
}

func (d Decimal) Int64() int64 {
	v, err := d.Decimal.Int64()
	if err != nil {
		panic(err)
	}
	return v
}

func (d Decimal) Float64() float64 {
	return d.Decimal.InexactFloat64()
}

func (d Decimal) Equal(e any) bool {
	return d.Decimal.Equal(New(e).Decimal)
}

func (d Decimal) EqualInt64(e int64) bool {
	return d.Decimal.Equal(Int64(e).Decimal)
}

func (d Decimal) EqualInt(e int) bool {
	return d.Decimal.Equal(Int64(int64(e)).Decimal)
}

func (d Decimal) EqualStr(e string) bool {
	return d.Decimal.Equal(Str(e).Decimal)
}

func (d Decimal) EqualDec(e Decimal) bool {
	return d.Decimal.Equal(e.Decimal)
}

func (d Decimal) GreaterThan(e any) bool {
	return d.Decimal.GreaterThan(New(e).Decimal)
}

func (d Decimal) GreaterThanInt64(e int64) bool {
	return d.Decimal.GreaterThan(Int64(e).Decimal)
}

func (d Decimal) GreaterThanInt(e int) bool {
	return d.Decimal.GreaterThan(Int64(int64(e)).Decimal)
}

func (d Decimal) GreaterThanStr(e string) bool {
	return d.Decimal.GreaterThan(Str(e).Decimal)
}

func (d Decimal) GreaterThanDec(e Decimal) bool {
	return d.Decimal.GreaterThan(e.Decimal)
}

func (d Decimal) LessThan(e any) bool {
	return d.Decimal.LessThan(New(e).Decimal)
}

func (d Decimal) LessThanInt64(e int64) bool {
	return d.Decimal.LessThan(Int64(e).Decimal)
}

func (d Decimal) LessThanInt(e int) bool {
	return d.Decimal.LessThan(Int64(int64(e)).Decimal)
}

func (d Decimal) LessThanStr(e string) bool {
	return d.Decimal.LessThan(Str(e).Decimal)
}

func (d Decimal) LessThanDec(e Decimal) bool {
	return d.Decimal.LessThan(e.Decimal)
}

func (d Decimal) GreaterThanOrEqual(e any) bool {
	return d.Decimal.GreaterThanOrEqual(New(e).Decimal)
}

func (d Decimal) GreaterThanOrEqualInt64(e int64) bool {
	return d.Decimal.GreaterThanOrEqual(Int64(e).Decimal)
}

func (d Decimal) GreaterThanOrEqualInt(e int) bool {
	return d.Decimal.GreaterThanOrEqual(Int64(int64(e)).Decimal)
}

func (d Decimal) GreaterThanOrEqualStr(e string) bool {
	return d.Decimal.GreaterThanOrEqual(Str(e).Decimal)
}

func (d Decimal) GreaterThanOrEqualDec(e Decimal) bool {
	return d.Decimal.GreaterThanOrEqual(e.Decimal)
}

func (d Decimal) LessThanOrEqual(e any) bool {
	return d.Decimal.LessThanOrEqual(New(e).Decimal)
}

func (d Decimal) LessThanOrEqualInt64(e int64) bool {
	return d.Decimal.LessThanOrEqual(Int64(e).Decimal)
}

func (d Decimal) LessThanOrEqualInt(e int) bool {
	return d.Decimal.LessThanOrEqual(Int64(int64(e)).Decimal)
}

func (d Decimal) LessThanOrEqualStr(e string) bool {
	return d.Decimal.LessThanOrEqual(Str(e).Decimal)
}

func (d Decimal) LessThanOrEqualDec(e Decimal) bool {
	return d.Decimal.LessThanOrEqual(e.Decimal)
}

func (d Decimal) Neg() Decimal {
	return Decimal{
		Decimal: d.Decimal.Neg(),
	}
}

func (d Decimal) Abs() Decimal {
	return Decimal{
		Decimal: d.Decimal.Abs(),
	}
}

// Premium returns (d - base) / base.
func (d Decimal) Premium(base any) Decimal {
	b := New(base)
	return d.SubDec(b).DivDec(b).MulInt64(100)
}

var Int64 = func(i int64) Decimal {
	return Decimal{
		Decimal: udecimal.MustFromInt64(i, 0),
	}
}

var Str = func(s string) Decimal {
	return Decimal{
		Decimal: udecimal.MustParse(s),
	}
}

var Float64 = func(f float64) Decimal {
	if d, err := udecimal.NewFromFloat64(f); err == nil {
		return Decimal{Decimal: d}
	}
	// Fallback: truncate float to default precision via int64 math.
	scale := math.Pow10(int(defaultPrec))
	scaled := math.Trunc(f * scale)
	if math.IsNaN(scaled) || math.IsInf(scaled, 0) || scaled > math.MaxInt64 || scaled < math.MinInt64 {
		panic(fmt.Sprintf("decimal: cannot convert float64 %v to Decimal", f))
	}
	d, err := udecimal.NewFromInt64(int64(scaled), defaultPrec)
	if err != nil {
		panic(fmt.Sprintf("decimal: cannot convert float64 %v to Decimal: %v", f, err))
	}
	return Decimal{Decimal: d}
}

func SetDefaultPrecision(p uint8) {
	udecimal.SetDefaultPrecision(p)
	defaultPrec = p
}

func NewFromBigInt(value *big.Int, exp int32) Decimal {
	if exp >= 0 && value.IsInt64() {
		// Multiply coefficient by 10^exp to get the integer value
		v := value.Int64()
		if exp == 0 {
			return Decimal{Decimal: udecimal.MustFromInt64(v, 0)}
		}
		// Check if v * 10^exp fits in int64
		mul := new(big.Int).Mul(value, new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(exp)), nil))
		if mul.IsInt64() {
			return Decimal{Decimal: udecimal.MustFromInt64(mul.Int64(), 0)}
		}
	} else if exp < 0 {
		prec := uint8(-exp)
		if value.IsInt64() && prec <= 19 {
			return Decimal{Decimal: udecimal.MustFromInt64(value.Int64(), prec)}
		}
	}

	// Fallback: udecimal has no big.Int constructor
	return Decimal{
		Decimal: udecimal.MustParse(BigIntExpToString(value, exp)),
	}
}

func New(v any) Decimal {
	switch v := v.(type) {
	case int:
		return Int64(int64(v))
	case int32:
		return Int64(int64(v))
	case int16:
		return Int64(int64(v))
	case int64:
		return Int64(v)
	case float32:
		return Float64(float64(v))
	case float64:
		return Float64(v)
	case string:
		return Str(v)
	case Decimal:
		return v
	default:
		panic("unsupported type")
	}
}
