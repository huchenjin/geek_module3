package service

import (
	"log"
	"net"
	"net/http"
	"os"
)

func Header(w http.ResponseWriter, r *http.Request) {
	// 接收客户端 request，并将 request 中带的 header 写入 response header
	for key := range r.Header {
		val := r.Header.Get(key)
		w.Header().Add(key, val)
	}
	w.Write([]byte("Header"))
}

func Version(w http.ResponseWriter, r *http.Request) {
	// 设置环境值的值
	err := os.Setenv("VERSION", "Geek Test Value")
	if err != nil {
		return
	}
	// 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	version := os.Getenv("VERSION")
	w.Header().Set("Version", version)

	w.Write([]byte("Version:" + version))
}

func ClientIP(w http.ResponseWriter, r *http.Request) {
	// 获取Client IP
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		log.Println("get ip err:", err)
	}
	//Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
	if net.ParseIP(ip) == nil {
		return
	}
	w.Write([]byte("client ip is :" + ip))
}

func Healthz(w http.ResponseWriter, r *http.Request) {
	// 当访问 localhost/healthz 时，应返回 200 默认就是200
	w.Write([]byte("当前请求：Healthz"))
}
