FROM golang:1.23 AS build

WORKDIR /src

COPY . .

RUN cd cmd/api && env GOOS=linux CGO_ENABLED=0 go build -o /bin/app ./main.go

FROM alpine:latest

COPY --from=build /bin/app /bin/app

CMD ["/bin/app"]


