##########
default: build
##########

build: deps lint test
	gox -output=".bin/{{.OS}}-{{.Arch}}/turing-machine" -verbose

clean:
	rm -fr ./.bin

deps:
	godep go install -v -x ./...

errcheck:
	errcheck github.com/jonathanmarvens/turing-machine

fmt:
	gofmt -e -s -w .

lint: fmt errcheck

test: lint
	# TODO(@jonathanmarvens): Add some damn tests to fix this.

.PHONY: default
.PHONY: build clean deps errcheck fmt lint test
