# 使用最新版 golang 作为基础镜像
FROM golang:latest AS builder
#设置工作目录，没有则自动新建
WORKDIR /go/src/
#拷贝代码到当前
COPY ../ /go/src/
## 编译
RUN cd /go/src
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .


##基于builder构建go_server
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
## 设置镜像环境变量（打印golang的GC日志）
ENV GODEBUG=gctrace=1
## 将二进制文件拷贝进当前镜像
COPY --from=builder /go/src/ .
## 暴露6379端口
EXPOSE 6379
#运行服务
CMD ["./main"]