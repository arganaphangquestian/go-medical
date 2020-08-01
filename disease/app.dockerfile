FROM golang:alpine AS go-build
WORKDIR /app
COPY . /app
RUN cd /app && go build -o disease

FROM alpine

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
WORKDIR /app
COPY --from=go-build /app/disease /app
EXPOSE 8080
ENTRYPOINT ./disease