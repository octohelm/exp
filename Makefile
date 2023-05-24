tidy:
	go mod tidy

test:
	CGO_ENABLED=0 go test -failfast ./...

test.race:
	CGO_ENABLED=1 go test -v -race ./...

lint:
	gofumpt -l -w .

dep:
	go get -u ./...
