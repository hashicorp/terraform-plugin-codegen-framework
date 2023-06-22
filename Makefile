build:
	go build ./cmd/terraform-plugin-codegen-framework

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
	go run ./cmd/terraform-plugin-codegen-framework generate all \
		--input ./internal/cmd/testdata/custom_and_external/ir.json \
		--output ./internal/cmd/testdata/custom_and_external/all_output

	go run ./cmd/terraform-plugin-codegen-framework generate resources \
		--input ./internal/cmd/testdata/custom_and_external/ir.json \
		--output ./internal/cmd/testdata/custom_and_external/resources_output

	go run ./cmd/terraform-plugin-codegen-framework generate data-sources \
		--input ./internal/cmd/testdata/custom_and_external/ir.json \
		--output ./internal/cmd/testdata/custom_and_external/data_sources_output

	go run ./cmd/terraform-plugin-codegen-framework generate provider \
		--input ./internal/cmd/testdata/custom_and_external/ir.json \
		--output ./internal/cmd/testdata/custom_and_external/provider_output

.PHONY: lint fmt test
