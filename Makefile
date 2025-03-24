
build:
	-go build -o ./tmp/main cmd/main.go
dev:
	# - lsof -ti :8080 | xargs kill -9
	- make init
	- air -c .air.toml
run:
	go run cmd/main.go