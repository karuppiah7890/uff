build:
	go get -v
	go mod tidy -v
	go mod vendor -v
	go build -mod=vendor -v -o out/uff

compile:
	go build -mod=vendor -v -o out/uff

install:
	go install -mod=vendor -v
