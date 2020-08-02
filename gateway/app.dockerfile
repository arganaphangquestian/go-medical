FROM golang:1.13-alpine3.11 AS build
RUN apk --no-cache add gcc g++ make ca-certificates
WORKDIR /go/src/github.com/arganaphangquestian/go-medical
COPY go.mod go.sum ./
COPY vendor vendor
COPY blood blood
COPY disease disease
COPY role role
COPY gender gender
COPY user user
COPY history history
COPY gateway gateway
RUN GO111MODULE=on go build -mod vendor -o /go/bin/app ./gateway/graph

FROM alpine:3.11
WORKDIR /usr/bin
COPY --from=build /go/bin .
EXPOSE 8080
CMD ["app"]
