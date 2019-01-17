test:
	go test -v ./...

ci:
	go test -race -coverprofile=coverage.txt -covermode=atomic -v ./...

lint:
	 gometalinter --vendor --fast --enable-gc --tests --aggregate --disable=gotype ./...

fmt:
	go fmt -w .