
.PHONY: binary
binary:
	CGO_ENABLED=0 \
	GOOS=linux \
	GOARCH=amd64 \
	go build -o ./build/go-service


build: binary
	podman build --platform linux/amd64  -t go-service .