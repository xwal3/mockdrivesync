.PHONY: run test test-ingest tidy clean

run:
	go run ./cmd/worker

test:
	go test ./... -v

test-ingest:
	go test ./internal/ingest -v

tidy:
	go mod tidy

clean:
	go clean -cache -testcache -modcache
