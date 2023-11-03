build:
	go build -o bin/gobp cmd/main/main.go

script:
	go build -o bin/gobp cmd/main/main.go
	cp bin/gobp /usr/local/bin/

run:
	go run cmd/main/main.go
