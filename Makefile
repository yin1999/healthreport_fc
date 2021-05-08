# flags
LDFLAGS=-ldflags "-s -w"

# make
all: dep build pack

dep:
	@echo Downloading dependencies...
	@go mod download

build: export GOOS=linux
build: export GOARCH=amd64
build:
	@echo Building...
	@go build ${LDFLAGS} -o bootstrap

pack:
	@echo packing...
	@zip -r bootstrap.zip bootstrap
	@rm -rf bootstrap