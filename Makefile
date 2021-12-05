# flags
targets?=aliyun tencent # default to all
LDFLAGS=-ldflags "-s -w -buildid="

# make
all: dep build

dep:
	@echo Downloading dependencies...
	@go mod download

build:
	@echo Building...
	@for target in ${targets}; do \
	GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -tags $${target} -trimpath -o bootstrap; \
	zip -rq $${target}-serverless.zip bootstrap; \
	done
	@rm -f bootstrap
