BINARY=app
BUILD_DIR=build
IMAGE_NAME=grpc-go-server
IMAGE_URI=lechuckroh/$(IMAGE_NAME)

.PHONY: dep
dep:
	go get && go mod vendor

.PHONY: build build-static
build:
	go build -o $(BINARY) *.go

build-linux64-static:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -mod=vendor -o $(BUILD_DIR)/$(BINARY) *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: image
image:
	docker build --rm -t $(IMAGE_NAME):latest .

.PHONY: push
push:
	docker tag $(IMAGE_NAME):latest $(IMAGE_URI):latest
	docker push $(IMAGE_URI):latest
	docker rmi $(IMAGE_URI):latest
