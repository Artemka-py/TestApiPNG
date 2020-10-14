FROM golang:latest AS buildContainer
WORKDIR /go/src/app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -v -mod mod -ldflags "-s -w" -o restapi .

FROM alpine:latest
WORKDIR /app
COPY --from=buildContainer /go/src/app/restapi .

ENV GIN_MODE release

ENV HOST 0.0.0.0
ENV PORT 8080

EXPOSE 8080

ENTRYPOINT ["./restapi"]
