FROM golang:1.13-alpine3.11 AS build
RUN apk update && apk --no-cache add gcc g++ git make ca-certificates
WORKDIR /go/src/github.com/arganaphangquestian/go-medical
COPY go.mod go.sum ./
COPY gateway gateway
RUN GO111MODULE=on go build -o /go/bin/gateway github.com/arganaphangquestian/go-medical/gateway/graph/*.go

FROM alpine:3.11
WORKDIR /usr/bin
COPY --from=build /go/bin .
EXPOSE 8080
CMD ["gateway"]
