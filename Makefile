DIR=./cmd/inliner

.PHONY: build

build:
	@go install ${DIR}

test:
	@go test ${DIR}

