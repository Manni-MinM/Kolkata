.PHONY: run
run:
	go run ./cmd/kolkata

.PHONY: resolve
resolve:
	go mod vendor

.PHONY: test
test:
	go test ./...

.PHONY: build-image
build-image:
	docker build -f build/package/Dockerfile -t kolkata:v1.0.0 .

