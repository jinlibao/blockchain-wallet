all: build test

build: clean
	( cd server/main ; go build )
	( cd client ; go build )

.PHONY: test
test:
	( cd server/addr ; go test )
	( cd server/block ; go test )
	( cd server/config ; go test )
	( cd server/hash ; go test )
	( cd server/lib ; go test )
	( cd server/merkle ; go test )
	( cd server/mine ; go test )
	( cd server/main ; make test )

.PHONY: run_server
run_server:
	( cd server/main ; ./run_server.sh)

.PHONY: run_client
run_client:
	( cd client ; ./run_client.sh)

clean:
	(cd server/main ; go clean ; rm -rf data out)
	(cd client ; go clean ; rm -rf wallet-data)
