.PHONY: cmd test

test:
	go test ./...

cmd:
	go build -o build/brato ./cmd/brato
