
GOPATHP?=$(PWD)/../../../..
GOPATHD?=$(HOME)/go

PB_HOME?=gpay
BIN_NAME?=gpay
IMG_NS?=docker.io/tangfeixiong
IMG_REPO?=gpay
IMG_TAG?=latest

GIT_COMMIT=$(shell date +%y%m%d%H%M)-git$(shell git rev-parse --short=7 HEAD)

DOCKER_FILE?=Dockerfile.busybox
DOCKER_REGISTRY?=172.17.4.59:5000/admin

all: protoc-grpc go-bindata-web docker-push

protoc-grpc:
	@if [[ ! -e $(GOPATHD) ]] ; \
	then \
		ln -sf /Users/fanhongling/go $(GOPATHD) ; \
	fi

	@protoc -I/usr/local/include -I. \
		-I${GOPATHP}/src \
		-I${GOPATHD}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I${GOPATHD}/src/github.com/gogo/protobuf \
		-I${GOPATHD}/src \
		--gogo_out=Mgoogle/api/annotations.proto=github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api,Mgogoproto/gogo.proto=github.com/gogo/protobuf/gogoproto,Mpb/datatype.proto=github.com/tangfeixiong/$(PB_HOME)/pb,plugins=grpc:. \
		pb/service.proto pb/datatype.proto
	@protoc -I/usr/local/include -I. \
		-I${GOPATHP}/src \
		-I${GOPATHD}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I${GOPATHD}/src/github.com/gogo/protobuf \
		-I${GOPATHD}/src \
		--grpc-gateway_out=logtostderr=true:. \
		pb/service.proto
	@protoc -I/usr/local/include -I. \
		-I${GOPATHP}/src \
		-I${GOPATHD}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I${GOPATHD}/src/github.com/gogo/protobuf \
		-I${GOPATHD}/src \
		--swagger_out=logtostderr=true:. \
		pb/service.proto

go-bindata-web:
	@pkg=webapp; src=static/...; output_file=pkg/ui/data/$${pkg}/datafile.go; \
		go-bindata -nocompress -o $${output_file} -prefix $${PWD} -pkg $${pkg} $${src}

go-bindata-swagger:
	#@pkg=artifact; src=template/...; output_file=pkg/spec/$${pkg}/artifacts.go; \
	#	go-bindata -nocompress -o $${output_file} -prefix $${PWD} -pkg $${pkg} $${src}
	@go generate ./pb
	@go-bindata -nocompress -o pkg/ui/data/swagger/datafile.go -prefix $${PWD} -pkg swagger third_party/swagger-ui/...

go-build:
	@CGO_ENABLED=0 go build -a -v -o ./bin/$${BIN_NAME} --installsuffix cgo -ldflags "-s" ./

go-install:
	@go install -v ./

docker-build:
	#@docker build --no-cache --force-rm -t $(IMG_NS)/$(IMG_REPO):$(IMG_TAG) ./
	@docker build --no-cache --force-rm -t $(IMG_NS)/$(IMG_REPO):$(IMG_TAG) -f $(DOCKER_FILE) ./

docker-push: docker-build
	@docker push $(IMG_NS)/$(IMG_REPO):$(IMG_TAG)

.PHONY: all protoc-grpc go-bindata-web go-bindata-spec go-build go-install docker-build docker-push
