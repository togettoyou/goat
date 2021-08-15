FROM golang:1.15 AS builder
ENV GO111MODULE=on
ENV GOPROXY https://goproxy.cn,direct
COPY . /root/togettoyou/
WORKDIR /root/togettoyou/
RUN make docs

FROM scratch
COPY --from=builder /root/togettoyou/go-server /root/togettoyou/
COPY --from=builder /root/togettoyou/conf/ /root/togettoyou/conf/
WORKDIR /root/togettoyou/
EXPOSE 8888
ENTRYPOINT ["./go-server"]