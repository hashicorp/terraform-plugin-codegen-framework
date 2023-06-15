lint:
	golangci-lint run

fmt:
	gofmt -s -w -e .

test:
	go test $$(go list ./... | grep -v /output) -v -cover -timeout=120s -parallel=4

# Generate copywrite headers
generate:
	cd tools; go generate ./...

# Regenerate testdata folder
testdata:
	go run . all \
		-input ./internal/cmd/testdata/basic/ir.json \
		-output ./internal/cmd/testdata/basic/

.PHONY: lint fmt test
