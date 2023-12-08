FROM golang:latest
LABEL authors="jimmy"

ENV GOPROXY https://goproxy.cn,direct
WORKDIR /app
COPY . .
RUN go build -o main ./cmd/main.go
EXPOSE 4000
ENTRYPOINT ["./gmall"]