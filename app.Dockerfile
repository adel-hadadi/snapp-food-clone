FROM alpine:latest

WORKDIR /src

COPY webApp .

COPY .env .

CMD [ "./webApp" ]
