FROM golang:1.19-alpine AS build

RUN apk add --no-cache git curl build-base

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
RUN go install github.com/google/wire/cmd/wire@latest

FROM golang:1.19-alpine

ARG USER_ID
ARG GROUP_ID
ENV GOBINPATH="$GOPATH/bin"
ENV BINPATH="/usr/local/bin"

RUN apk add --no-cache gcc g++

COPY --from=build $GOBINPATH/air $BINPATH/air
COPY --from=build $GOBINPATH/golangci-lint $BINPATH/golangci-lint
COPY --from=build $GOBINPATH/wire $BINPATH/wire

RUN addgroup -g ${GROUP_ID} web && adduser -u ${USER_ID} -G web -s /bin/sh -D web

WORKDIR /app

USER web
