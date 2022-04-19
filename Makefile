all: docker m1 install
	
m1:
	env GOOS=darwin GOARCH=arm64 go build -o ./bin
	
docker:
	env GOOS=linux GOARCH=amd64 go build -o ./bin/docker/logsink

install:
	cp ./bin/logsink ~/go/bin
	chmod +x ~/go/bin/logsink

sign:
	codesign -s - ~/go/bin/logsink
