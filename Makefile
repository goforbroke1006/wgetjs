SERVICE_NAME=wgetjs
ENTRY_POINT=./cmd/wgetjs.go

all:
	echo 'undefined task'

.PHONY: build
build:
	GOOS=linux   GOARCH=amd64 go build -o ./build/${SERVICE_NAME}        ${ENTRY_POINT}
	GOOS=windows GOARCH=386   go build -o ./build/${SERVICE_NAME}.exe    ${ENTRY_POINT}
	GOOS=windows GOARCH=amd64 go build -o ./build/${SERVICE_NAME}_64.exe ${ENTRY_POINT}