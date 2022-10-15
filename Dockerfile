FROM golang:1.18.2 AS build
WORKDIR /httpserver
COPY .  /httpserver
ENV CGO_ENABLED=0
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
RUN cd /httpserver && GOOS=linux go build -installsuffix cgo -o httpserver main.go

FROM ubuntu
COPY --from=build /httpserver/httpserver /httpserver/httpserver
EXPOSE 80
ENV ENV local
WORKDIR /httpserver/
ENTRYPOINT ["./httpserver"]