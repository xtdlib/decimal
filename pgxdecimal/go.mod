module github.com/xtdlib/decimal/pgxdecimal

go 1.26.1

replace github.com/xtdlib/decimal => ..

require (
	github.com/jackc/pgx/v5 v5.9.1
	github.com/xtdlib/decimal v0.0.0-00010101000000-000000000000
)

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/quagmt/udecimal v1.10.0 // indirect
	golang.org/x/text v0.29.0 // indirect
)
