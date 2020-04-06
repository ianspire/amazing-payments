INCLUDES := -I. -I/go/src

all: build mocks

build:
	docker build -t ianspire/amazing-payments/proto .
	docker run --rm -v $(CURDIR):/go/src/github.com/ianspire/amazing-payments/proto -w /go/src/github.com/ianspire/amazing-payments/proto ianspire/proto clean generate

clean: proto-clean mock-clean swagger-clean

generate: proto-generate swagger-generate

proto-generate:
	protoc ${INCLUDES} --go_out=plugins=grpc:/go/src /go/src/github.com/ianspire/amazing-payments/proto/*.proto
	protoc ${INCLUDES} --grpc-gateway_out=logtostderr=true:/go/src /go/src/github.com/ianspire/amazing-payments/proto/*.proto

swagger-generate:
	protoc ${INCLUDES} --swagger_out=logtostderr=true:/go/src /go/src/github.com/ianspire/amazing-payments/proto/*.proto

proto-clean:
	rm -f *.pb.go *.pb.gw.go

swagger-clean:
	rm -f *.swagger.json

mocks:
	GO111MODULE=on go get github.com/golang/mock/mockgen@latest
	mkdir ./mocks
	mockgen github.com/ianspire/amazing-payments/proto PaymentServiceClient,PaymentServiceServer > ./mocks/mockPayment.go

mock-clean:
	rm -rf ./mocks
