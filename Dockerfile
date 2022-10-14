FROM golang:go1.18.2 AS build
WORKDIR /httpserver/
COPY . .
ENV CGO_ENABLED=0
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
RUN GOOS=linux go build -installsuffix cgo -o httpserver main/main.go

FROM busybox
COPY --from=build /httpserver/httpserver /httpserver/httpserver
EXPOSE 80
ENV ENV local
WORKDIR /httpserver/
ENTRYPOINT ["./httpserver"]