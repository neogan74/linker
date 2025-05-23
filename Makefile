.PHONY: generate
generate: generate_proto generate_openapi

.PHONY: generate_proto
generate_proto:
	@if ! command -v 'easyp' &> /dev/null; then \
		echo "Please install easyp!"; exit 1; \
	fi;
	@easyp generate

.PHONY: generate_openapi
generate_openapi:
	@if ! command -v 'oapi-codegen' &> /dev/null; then \
		echo "Please install oapi-codegen!"; exit 1; \
	fi;
	@mkdir -p internal/api/openapi/v1
	@mkdir -p internal/api/openapi/v1/bot-api
	@oapi-codegen -package v1 \
		-generate server,types \
		api/openapi/v1/bot-api.yaml > internal/api/openapi/v1/bot-api/bot-api.gen.go
	@mkdir -p internal/api/openapi/v1/scrapper
	@oapi-codegen -package v1 \
		-generate server,types \
		api/openapi/v1/scrapper.yaml > internal/api/openapi/v1/scrapper/scrapper.gen.go

.PHONY: clean
clean:
	@rm -rf./bin