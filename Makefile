.PHONY: check_install
check_install:
	which swagger || GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger

.PHONY: swagger
swagger: check_install
	GO111MODULE=off swagger generate spec -o ./swagger.yaml --scan-models

.PHONY: run
run:
	go run main.go

.PHONY: build
build:
	go build