module github.com/xtdlib/decimal/bench

go 1.26.1

require (
	github.com/cockroachdb/apd v1.1.0
	github.com/xtdlib/decimal v0.0.0-00010101000000-000000000000
)

require (
	github.com/lib/pq v1.12.3 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/quagmt/udecimal v1.10.0 // indirect
)

replace github.com/xtdlib/decimal => ..
