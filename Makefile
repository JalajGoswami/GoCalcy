.DEFAULT_GOAL = build
Name = Jalaj

build:
	CC=gcc.exe CGO_ENABLED=1 GOARCH=amd64 GOOS=windows go build -o ./bin/main.exe

run: buildLinux
	./bin/main

buildLinux:
	CGO_ENABLED=1 GOARCH=amd64 go build -o ./bin/main

greet:
	echo ${Name}
