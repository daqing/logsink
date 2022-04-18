all: docker
	go build -o ./bin

docker:
	env GOOS=linux GOARCH=amd64 go build -o ./bin/docker/logsink
