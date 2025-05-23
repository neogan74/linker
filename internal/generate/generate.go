package generate

//go:generate oapi-codegen -config bot_server_config.yaml ../../api/openapi/v1/bot-api.yaml
//go:generate oapi-codegen -config scrapper_server_config.yaml ../../api/openapi/v1/scrapper-api.yaml
