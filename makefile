run:
	go run src/main.go

compile:
	sqlc generate
	mkdir -p dist/bin
	GOOS=linux GOARCH=arm go build -o bin/main-linux-arm src/main.go
	GOOS=linux GOARCH=arm64 go build -o bin/main-linux-arm64 src/main.go
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 src/main.go