FROM golang:1.16.0-alpine3.13 AS builder
LABEL MAINTAINER=github.com/dezhiShen
RUN go env -w GO111MODULE=auto \
  && go env -w GOPROXY=https://goproxy.cn,direct 
WORKDIR /build
COPY ./ .
RUN cd /build && go build -tags netgo -o miraigo cmd/main.go

FROM alpine:latest
COPY --from=builder /build/miraigo /usr/bin/miraigo
RUN chmod +x /usr/bin/miraigo
WORKDIR /data
ENTRYPOINT ["/usr/bin/miraigo"]