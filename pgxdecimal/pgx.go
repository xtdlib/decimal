package pgxdecimal

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/xtdlib/decimal"
)

type Decimal struct {
	decimal.Decimal
}

func (d *Decimal) ScanNumeric(v pgtype.Numeric) error {
	if !v.Valid {
		return fmt.Errorf("cannot scan NULL into *decimal.Decimal")
	}

	if v.NaN {
		return fmt.Errorf("cannot scan NaN into *decimal.Decimal")
	}

	if v.InfinityModifier != pgtype.Finite {
		return fmt.Errorf("cannot scan %v into *decimal.Decimal", v.InfinityModifier)
	}

	bi := v.Int
	if bi == nil {
		bi = big.NewInt(0)
	}

	d.Decimal = decimal.NewFromBigInt(bi, v.Exp)
	return nil
}

func (d Decimal) NumericValue() (pgtype.Numeric, error) {
	neg, hi, lo, prec, ok := d.Decimal.ToHiLo()
	if ok && hi == 0 {
		coef := new(big.Int).SetUint64(lo)
		if neg {
			coef.Neg(coef)
		}
		return pgtype.Numeric{Int: coef, Exp: -int32(prec), Valid: true}, nil
	}

	// Fallback for large numbers
	s := d.String()
	coef := new(big.Int)
	var exp int32
	if idx := strings.IndexByte(s, '.'); idx >= 0 {
		digits := s[:idx] + s[idx+1:]
		coef.SetString(digits, 10)
		exp = -int32(len(s) - idx - 1)
	} else {
		coef.SetString(s, 10)
	}
	return pgtype.Numeric{Int: coef, Exp: exp, Valid: true}, nil
}
