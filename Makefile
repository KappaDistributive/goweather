BINARY_NAME=goweather

build:
	go build -o ${BINARY_NAME} v1/*.go

clean:
	go clean
	rm ${BINARY_NAME}
