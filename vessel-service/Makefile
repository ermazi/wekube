build:
	@echo ${PATH}
	protoc -I. --go_out=plugins=micro:/Users/wangshubin/go/src/github.com/wekube/vessel-service \
	proto/vessel/vessel.proto
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
	docker build -t vessel-service .
run:
	docker run --rm -p 50052:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns vessel-service
gen:
	protoc -I. --go_out=plugins=micro:/Users/wangshubin/go/src/github.com/wekube/vessel-service \
	proto/vessel/vessel.proto