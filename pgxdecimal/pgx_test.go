package pgxdecimal

import (
	"context"
	"math/big"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/xtdlib/decimal"
)

func TestScanNumeric(t *testing.T) {
	// 3.14 = Int=314, Exp=-2
	n := pgtype.Numeric{Int: big.NewInt(314), Exp: -2, Valid: true}
	var d Decimal
	if err := d.ScanNumeric(n); err != nil {
		t.Fatal(err)
	}
	if d.String() != "3.14" {
		t.Errorf("ScanNumeric = %s, want 3.14", d.String())
	}
}

func TestScanNumericZero(t *testing.T) {
	n := pgtype.Numeric{Int: big.NewInt(0), Exp: 0, Valid: true}
	var d Decimal
	if err := d.ScanNumeric(n); err != nil {
		t.Fatal(err)
	}
	if d.String() != "0" {
		t.Errorf("ScanNumeric = %s, want 0", d.String())
	}
}

func TestScanNumericNegative(t *testing.T) {
	n := pgtype.Numeric{Int: big.NewInt(-15), Exp: -1, Valid: true}
	var d Decimal
	if err := d.ScanNumeric(n); err != nil {
		t.Fatal(err)
	}
	if d.String() != "-1.5" {
		t.Errorf("ScanNumeric = %s, want -1.5", d.String())
	}
}

func TestScanNumericNull(t *testing.T) {
	n := pgtype.Numeric{Valid: false}
	var d Decimal
	if err := d.ScanNumeric(n); err == nil {
		t.Error("ScanNumeric(NULL) should error")
	}
}

func TestScanNumericNaN(t *testing.T) {
	n := pgtype.Numeric{Valid: true, NaN: true}
	var d Decimal
	if err := d.ScanNumeric(n); err == nil {
		t.Error("ScanNumeric(NaN) should error")
	}
}

func TestPostgresRoundTrip(t *testing.T) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgres://postgres:postgres@oci-aca-001:5432/postgres")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close(ctx)

	conn.TypeMap().RegisterType(&pgtype.Type{
		Name:  "numeric",
		OID:   pgtype.NumericOID,
		Codec: pgtype.NumericCodec{},
	})

	// Write
	want := Decimal{Decimal: decimal.Str("3.14")}
	var got Decimal
	err = conn.QueryRow(ctx, "SELECT $1::numeric", want).Scan(&got)
	if err != nil {
		t.Fatal(err)
	}
	if got.String() != "3.14" {
		t.Errorf("got %s, want 3.14", got.String())
	}
}

func TestPostgresTable(t *testing.T) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgres://postgres:postgres@oci-aca-001:5432/postgres")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close(ctx)

	_, err = conn.Exec(ctx, "DROP TABLE IF EXISTS test_decimal")
	if err != nil {
		t.Fatal(err)
	}
	_, err = conn.Exec(ctx, "CREATE TABLE test_decimal (price numeric NOT NULL)")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Exec(ctx, "DROP TABLE test_decimal")

	want := Decimal{Decimal: decimal.Str("123.45")}
	_, err = conn.Exec(ctx, "INSERT INTO test_decimal (price) VALUES ($1)", want)
	if err != nil {
		t.Fatal(err)
	}

	var got Decimal
	err = conn.QueryRow(ctx, "SELECT price FROM test_decimal").Scan(&got)
	if err != nil {
		t.Fatal(err)
	}
	if got.String() != "123.45" {
		t.Errorf("got %s, want 123.45", got.String())
	}
}

func TestNumericValue(t *testing.T) {
	d := Decimal{Decimal: decimal.Str("3.14")}
	n, err := d.NumericValue()
	if err != nil {
		t.Fatal(err)
	}
	if !n.Valid {
		t.Fatal("NumericValue not valid")
	}
	var d2 Decimal
	if err := d2.ScanNumeric(n); err != nil {
		t.Fatal(err)
	}
	if d2.String() != "3.14" {
		t.Errorf("round-trip = %s, want 3.14", d2.String())
	}
}
