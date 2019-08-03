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

releases: build
	GOOS=linux GOARCH=arm CGO_ENABLED=0 go build -o ./bin/troncli.arm.linux ./*.go 
	GOOS=linux GOARCH=386 CGO_ENABLED=0 go build -o ./bin/troncli.386.linux ./*.go 
	GOOS=windows GOARCH=386 CGO_ENABLED=0 go build -o ./bin/troncli.exe ./*.go 


clean:
	rm -rf bin