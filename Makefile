run:
	docker run --rm --env GOPATH=/go -v ${PWD}:/go/src/app -w /go/src/app golang:1.13.5-buster go run main.go

test:
	docker run --rm --env GOPATH=/go -v ${PWD}:/go/src/app -w /go/src/app golang:1.13.5-buster go test -v