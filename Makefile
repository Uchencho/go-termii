OUTPUT = main

.PHONY: test
test:
	go test ./...

.PHONY: clean
clean:
	rm -f $(OUTPUT)

build-local: 
	go build -o $(OUTPUT) ./client/default.go

run: build-local
	@echo ">> Running application ..."
	TERMII_API_KEY=TLhOpqZJBTAogePP6k7odCheznezm6KA4MLUb1FNMI4K7dlQ1gXgMqhfgw8Cwr \
	TERMII_URL=https://termii.com/ \
	TERMII_SENDER_ID=commands-stream-three-dev \
	./$(OUTPUT)
