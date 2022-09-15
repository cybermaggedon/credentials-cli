
GOFILES=credentials-cli

all: ${GOFILES}

SOURCES=credentials-cli.go mgmt.go

credentials-cli: ${SOURCES}
	GOPATH=$$(pwd)/go go build ${SOURCES}

credentials-cli.macos:${SOURCES}
	GOOS=darwin GOARCH=386 GOPATH=$$(pwd)/go go build -o $@ ${SOURCES}

upload:
	gsutil cp credentials-cli.macos \
	    gs://example.example.com/example-creds/credentials-cli.macos

clean:
	rm -rf ${GOFILES}
	rm -rf go

