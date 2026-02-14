BINARY_NAME=goweather

.PHONY: build test clean

build:
	CGO_LDFLAGS_ALLOW="-sectcreate" \
	go build -ldflags='-extldflags "-sectcreate __TEXT __info_plist ./Info.plist"' \
	-o $(BINARY_NAME) .

test:
	go test ./...

clean:
	rm -f $(BINARY_NAME)
	go clean
