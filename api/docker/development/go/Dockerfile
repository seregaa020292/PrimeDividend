FROM golang:1.18-alpine AS build

RUN apk add --no-cache git curl

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
RUN curl -sSfL https://raw.githubusercontent.com/pressly/goose/master/install.sh | GOOSE_INSTALL=$GOPATH sh
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
RUN go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
RUN go install github.com/google/wire/cmd/wire@latest

FROM golang:1.18-alpine

ARG USER_ID
ARG GROUP_ID
ENV BINPATH="/usr/local/bin"

RUN apk add --no-cache gcc g++ make bash coreutils unzip apache2-utils ca-certificates

COPY ./common/wait-for-it.sh $BINPATH/wait-for-it
COPY ./common/migrate.sh $BINPATH/migrate
RUN chmod 555 $BINPATH/wait-for-it $BINPATH/migrate

COPY --from=build $GOPATH/bin/air $BINPATH/air
COPY --from=build $GOPATH/bin/goose $BINPATH/goose
COPY --from=build $GOPATH/bin/golangci-lint $BINPATH/golangci-lint
COPY --from=build $GOPATH/bin/oapi-codegen $BINPATH/oapi-codegen
COPY --from=build $GOPATH/bin/wire $BINPATH/wire

RUN addgroup -g ${GROUP_ID} web && adduser -u ${USER_ID} -G web -s /bin/sh -D web

WORKDIR /app

USER web
