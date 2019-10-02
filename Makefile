build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/cron handlers/cron/main.go
	env GOOS=linux go build -tags=jsoniter -ldflags="-s -w" -o bin/app handlers/app/main.go

