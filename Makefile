.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -o chatgpt-lucy ./main.go

.PHONY: docker
docker:
	docker build . -t chatgpt-lucy:latest
