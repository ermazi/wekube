build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
	sudo -H docker build -t consignment-cli .
run:
	sudo -H docker run --rm -e MICRO_REGISTRY=mdns consignment-cli