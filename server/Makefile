
all: build test

build:
	( cd main ; go build )

.PHONY: test
test:
	( cd addr ; go test )
	( cd block ; go test )
	( cd config ; go test )
	( cd hash ; go test )
	( cd lib ; go test )
	( cd merkle ; go test )
	( cd mine ; go test )
	( cd main ; make test )

