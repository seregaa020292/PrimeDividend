#!/bin/bash
set -e

readonly service="$1"
readonly path="../../internal/handlers/http"

mkdir -p "$path/$service"

oapi-codegen -old-config-style \
  -generate "types,chi-server" \
  -include-tags "$service" \
  -package "$service" \
  -o "$path/$service/openapi.gen.go" \
  "./swagger.yml"
