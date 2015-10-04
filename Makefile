say:
	@go run ./cmd/histoire-pour-enfant-generator/main.go | tee /tmp/a
	@cat /tmp/a | say
