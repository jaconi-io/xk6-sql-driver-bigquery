all: test build example

test: *.go testdata/*.js
	go test -count 1 ./...

build: k6

k6: *.go go.mod go.sum
	xk6 build --with github.com/grafana/xk6-sql@latest --with github.com/jaconi-io/xk6-sql-driver-bigquery=.

example: k6
	./k6 run examples/example.js

.PHONY: test all example
