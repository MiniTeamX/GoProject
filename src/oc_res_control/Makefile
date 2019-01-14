SRCS=$(wilcard *.go)
PROGRAM=xsafe.filter_service
VERSION=$(shell git describe --always --long --dirty)

# make
# make all

.PHONY: all
all: vendor $(PROGRAM)

$(PROGRAM): $(SRCS)
	@go version
	go build -o $@ -gcflags "-N -l" -ldflags="-X main._VERSION_=$(VERSION)" *.go

# make target
.PHONY: target
target: $(PROGRAM)

# make vendor
vendor: glide.yaml
	glide install

# make clean
.PHONY: clean
clean:
	rm -rf $(PROGRAM)

.PHONY: upload
upload:
	scp $(PROGRAM) root1@192.168.26.193:/var/run/xsafe/filter_service/bin/
