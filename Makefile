VERSION=$(shell cat VERSION)

format:
	gofmt -w -s .

test:
	go test -v ./...

push:
	git tag $(VERSION)
	git push origin $(VERSION)