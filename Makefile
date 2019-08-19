default: build

pre:
	export PATH=$PATH:$GOPATH/bin

fmt:
	go fmt ./...

run:
	GO111MODULE=on go run ./*.go

install:
	GO111MODULE=on go install ./

bin:
	mkdir -p bin

build: bin
	GO111MODULE=on go build -o ./bin/troncli ./*.go

releases:
	GOOS=linux GOARCH=386 CGO_ENABLED=0 go build -o ./bin/troncli.386.linux ./*.go
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./bin/troncli.amd64.linux ./*.go
	GOOS=linux GOARCH=arm CGO_ENABLED=0 go build -o ./bin/troncli.arm.linux ./*.go
	GOOS=windows GOARCH=386 CGO_ENABLED=0 go build -o ./bin/troncli.386.exe ./*.go
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o ./bin/troncli.amd64.exe ./*.go
	GOOS=windows GOARCH=arm CGO_ENABLED=0 go build -o ./bin/troncli.arm.exe ./*.go
	GOOS=darwin GOARCH=386 CGO_ENABLED=0 go build -o ./bin/troncli.386.mac ./*.go
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o ./bin/troncli.amd64.mac ./*.go 
		


clean:
	rm -rf bin