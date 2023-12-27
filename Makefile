build:
	go build ./cmd/tfplugingen-framework

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
	go run ./cmd/tfplugingen-framework generate all \
		--input ./internal/cmd/testdata/custom_and_external/ir.json \
		--package specified \
		--output ./internal/cmd/testdata/custom_and_external/all_output/specified_pkg_name

	go run ./cmd/tfplugingen-framework generate all \
		--input ./internal/cmd/testdata/custom_and_external/ir.json \
		--output ./internal/cmd/testdata/custom_and_external/all_output/default_pkg_name

	go run ./cmd/tfplugingen-framework generate resources \
		--input ./internal/cmd/testdata/custom_and_external/ir.json \
		--package generated \
		--output ./internal/cmd/testdata/custom_and_external/resources_output

	go run ./cmd/tfplugingen-framework generate data-sources \
		--input ./internal/cmd/testdata/custom_and_external/ir.json \
		--package generated \
		--output ./internal/cmd/testdata/custom_and_external/data_sources_output

	go run ./cmd/tfplugingen-framework generate provider \
		--input ./internal/cmd/testdata/custom_and_external/ir.json \
		--package generated \
		--output ./internal/cmd/testdata/custom_and_external/provider_output

	go run ./cmd/tfplugingen-framework generate all \
		--input ./internal/cmd/testdata/template_tests/ir.json \
		--output ./internal/cmd/testdata/template_tests/output \
		--templates ./internal/cmd/testdata/template_tests/codegen_templates

	go run ./cmd/tfplugingen-framework scaffold resource \
		--name thing \
		--force \
		--package scaffold \
		--output-dir ./internal/cmd/testdata/scaffold/resource

	go run ./cmd/tfplugingen-framework scaffold data-source \
		--name thing \
		--force \
		--package scaffold \
		--output-dir ./internal/cmd/testdata/scaffold/data_source

	go run ./cmd/tfplugingen-framework scaffold provider \
		--name examplecloud \
		--force \
		--package scaffold \
		--output-dir ./internal/cmd/testdata/scaffold/provider

.PHONY: lint fmt test
