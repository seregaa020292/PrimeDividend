FROM golang:1.19-alpine AS build

RUN apk add --no-cache git curl build-base

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
RUN curl -sSfL https://raw.githubusercontent.com/pressly/goose/master/install.sh | GOOSE_INSTALL=$GOPATH sh
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
RUN go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
RUN go install github.com/google/wire/cmd/wire@latest
RUN go install github.com/go-jet/jet/v2/cmd/jet@latest

FROM golang:1.19-alpine

ARG USER_ID
ARG GROUP_ID
ENV GOBINPATH="$GOPATH/bin"
ENV BINPATH="/usr/local/bin"

RUN apk add --no-cache gcc g++ make bash coreutils unzip apache2-utils ca-certificates

COPY ./common/wait-for-it.sh $BINPATH/wait-for-it
COPY ./common/migrate.sh $BINPATH/migrate
COPY ./common/gojet.sh $BINPATH/gojet
RUN chmod 555 $BINPATH/wait-for-it $BINPATH/migrate $BINPATH/gojet

COPY --from=build $GOBINPATH/air $BINPATH/air
COPY --from=build $GOBINPATH/goose $BINPATH/goose
COPY --from=build $GOBINPATH/golangci-lint $BINPATH/golangci-lint
COPY --from=build $GOBINPATH/oapi-codegen $BINPATH/oapi-codegen
COPY --from=build $GOBINPATH/wire $BINPATH/wire
COPY --from=build $GOBINPATH/jet $BINPATH/jet

RUN addgroup -g ${GROUP_ID} web && adduser -u ${USER_ID} -G web -s /bin/sh -D web

WORKDIR /app

USER web
