default: build tearup test teardown
run: build tearup

build:
	GOPATH=$$(pwd) CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./src/gateway/server ./src/gateway && \
	GOPATH=$$(pwd) CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./src/driver/server ./src/driver
tearup:
	docker-compose up --build --force-recreate --detach
teardown:
	docker-compose rm --force --stop
test: 
	GOPATH=$$(pwd) go test -timeout 30s ./...
