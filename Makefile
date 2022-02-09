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
	TERMII_API_KEY=FILL-ME \
	TERMII_URL=FILL-ME \
	TERMII_SENDER_ID=FILL-ME \
	./$(OUTPUT)
