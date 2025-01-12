FROM golang:1.23 AS builder

WORKDIR /src

COPY go.mod go.sum .

RUN go mod download

COPY . .

ENV GOCACHE=/root/.cache/go-build

RUN --mount=type=cache,target="/root/.cache/go-build" cd ./cmd/cron && env GOOS=linux CGO_ENABLED=0 go build -o /bin/cronApp ./main.go

FROM alpine:3.21.2 AS base

WORKDIR /src

COPY --from=builder /bin/cronApp /bin/cronApp

COPY .env .

CMD [ "/bin/cronApp" ]
