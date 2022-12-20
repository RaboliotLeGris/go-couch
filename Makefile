default:
	go run .

start_couchdb:
	ls

test: lint
	go test ./...

lint:
	golangci-lint run

fmt:
	gofmt -s -w .