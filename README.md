# geek_module3

1. 进入服务器目录，执行 docker image build -t hugh-httpserver:v1 . 构建本地镜像
```shell
root@iZ2ze06j2vnul9uxekp3lvZ:/home/geek_module3# docker image build -t hugh-httpserver:v1 .
Sending build context to Docker daemon  88.06kB
Step 1/13 : FROM golang:1.18.2 AS build
1.18.2: Pulling from library/golang
e756f3fdd6a3: Pull complete
bf168a674899: Pull complete
e604223835cc: Pull complete
6d5c91c4cd86: Pull complete
93c221c34e03: Pull complete
b7b49b07aaad: Pull complete
26ce77466dae: Pull complete
Digest: sha256:04fab5aaf4fc18c40379924674491d988af3d9e97487472e674d0b5fd837dfac
Status: Downloaded newer image for golang:1.18.2
 ---> 738636f06010
Step 2/13 : WORKDIR /httpserver
 ---> Running in 9869cb976fc9
Removing intermediate container 9869cb976fc9
 ---> 0fb340bcf842
Step 3/13 : COPY .  /httpserver
 ---> 674edc1866f9
Step 4/13 : ENV CGO_ENABLED=0
 ---> Running in 230332b49172
Removing intermediate container 230332b49172
 ---> 94c99504ef41
Step 5/13 : ENV GO111MODULE=on
 ---> Running in 6ef4844640e4
Removing intermediate container 6ef4844640e4
 ---> f5471ca469a2
Step 6/13 : ENV GOPROXY=https://goproxy.cn,direct
 ---> Running in 0a9be05eca61
Removing intermediate container 0a9be05eca61
 ---> 5f76dc4ec00b
Step 7/13 : RUN cd /httpserver && GOOS=linux go build -installsuffix cgo -o httpserver main.go
 ---> Running in 776fb176af1d
Removing intermediate container 776fb176af1d
 ---> 8851266d083e
Step 8/13 : FROM ubuntu
 ---> 216c552ea5ba
Step 9/13 : COPY --from=build /httpserver/httpserver /httpserver/httpserver
 ---> 1864e17a1b95
Step 10/13 : EXPOSE 80
 ---> Running in 20a14570f733
Removing intermediate container 20a14570f733
 ---> 89fe976124c8
Step 11/13 : ENV ENV local
 ---> Running in db891ea7eced
Removing intermediate container db891ea7eced
 ---> 6a99f617174b
Step 12/13 : WORKDIR /httpserver/
 ---> Running in cd2345f2bcfb
Removing intermediate container cd2345f2bcfb
 ---> fed731b97230
Step 13/13 : ENTRYPOINT ["./httpserver"]
 ---> Running in 1df62c687e93
Removing intermediate container 1df62c687e93
 ---> ca2745690769
Successfully built ca2745690769
Successfully tagged hugh-httpserver:v1
```

查看当前镜像列表
```shell
root@iZ2ze06j2vnul9uxekp3lvZ:/home/geek_module3# docker image ls
REPOSITORY        TAG       IMAGE ID       CREATED          SIZE
hugh-httpserver   v1        ca2745690769   26 seconds ago   84.1MB
```

将镜像推送至 docker 官方镜像仓库 hub.docker.com
```shell
root@iZ2ze06j2vnul9uxekp3lvZ:/home/geek_module3# docker tag hugh-httpserver:v1 huchenjin/hugh-httpserver:v1
root@iZ2ze06j2vnul9uxekp3lvZ:/home/geek_module3# docker push huchenjin/hugh-httpserver:v1
The push refers to repository [docker.io/huchenjin/hugh-httpserver]
3e1f07147ee0: Pushed
17f623af01e2: Mounted from library/ubuntu
v1: digest: sha256:b673f52bf132ccefc8a18c33bc62bcb4d74a543ce2254bf427ef1b8f4602bcf3 size: 740
```

通过 docker 命令本地启动 httpserver
```shell
root@iZ2ze06j2vnul9uxekp3lvZ:~# docker run -d --name httpserver -p 8080:80 huchenjin/hugh-httpserver:v1
be5f2241b0aaba4311c29a2f352132fd5bc0d69988cd7b64a03e717ccde60edb
root@iZ2ze06j2vnul9uxekp3lvZ:~#
root@iZ2ze06j2vnul9uxekp3lvZ:~#
root@iZ2ze06j2vnul9uxekp3lvZ:~# docker ps -a
CONTAINER ID   IMAGE                          COMMAND          CREATED         STATUS         PORTS                                   NAMES
be5f2241b0aa   huchenjin/hugh-httpserver:v1   "./httpserver"   7 seconds ago   Up 6 seconds   0.0.0.0:8080->80/tcp, :::8080->80/tcp   httpserver
root@iZ2ze06j2vnul9uxekp3lvZ:~#
```

通过 nsenter 进入容器查看 IP 配置
```shell
root@iZ2ze06j2vnul9uxekp3lvZ:~# docker inspect -f {{.State.Pid}} be5f2241b0aa
30922
root@iZ2ze06j2vnul9uxekp3lvZ:~# nsenter -n -t 30922
root@iZ2ze06j2vnul9uxekp3lvZ:~# ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
16: eth0@if17: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default
    link/ether 02:42:ac:12:00:02 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet 172.18.0.2/16 brd 172.18.255.255 scope global eth0
       valid_lft forever preferred_lft forever
root@iZ2ze06j2vnul9uxekp3lvZ:~#
```
