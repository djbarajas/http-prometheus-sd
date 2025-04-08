DOCKER := $(shell if which podman >/dev/null 2>/dev/null; then echo podman; else echo docker; fi)

.PHONY: binary
binary:
	CGO_ENABLED=0 \
	GOOS=linux \
	GOARCH=amd64 \
	go build -o ./build/http-service-discovery

.PHONY: build
build: binary
	$(DOCKER) build --platform linux/amd64  -t http-service-discovery .