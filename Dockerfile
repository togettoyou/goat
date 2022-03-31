FROM golang:1.16 AS builder
ENV GO111MODULE=on
ENV GOPROXY https://goproxy.cn,direct
WORKDIR /build
ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
RUN make

FROM scratch
WORKDIR /app
COPY --from=builder /build/go-server .
COPY --from=builder /build/conf/ ./conf/
ENTRYPOINT ["./go-server"]