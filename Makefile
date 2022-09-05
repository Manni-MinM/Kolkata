.PHONY: run
run:
	go run ./cmd/kolkata

.PHONY: resolve
resolve:
	go mod vendor

.PHONY: test
test:
	go test ./...

.PHONY: image
image:
	docker build -f build/package/Dockerfile -t kolkata:v1.0.0 .

.PHONY: compose-up
compose-up:
	docker-compose -f deployments/docker-compose.yaml up

.PHONY: compose-down
compose-down:
	docker-compose -f deployments/docker-compose.yaml down --rmi=local

