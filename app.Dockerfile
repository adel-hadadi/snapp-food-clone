FROM golang:1.23 AS build

WORKDIR /src

COPY go.mod go.sum .

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=go.sum,target=go.sum \
    --mount=type=bind,source=go.mod,target=go.mod \
    go mod download -x

COPY . .

ENV GOCACHE=/root/.cache/go-build

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=cache,target="/root/.cache/go-build" \
    --mount=type=bind,target=. \
    cd ./cmd/api && env GOOS=linux CGO_ENABLED=0 go build -o /bin/webApp ./main.go

FROM alpine:3.21.2 AS base

WORKDIR /src

COPY --from=build /bin/webApp /bin/webApp

COPY .env .

CMD [ "/bin/webApp" ]
