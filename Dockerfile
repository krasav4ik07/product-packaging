FROM golang:1.18

WORKDIR ./src/product-packaging
COPY ./ ./

RUN go mod download
RUN go build -o product-packaging ./cmd/main.go

ENTRYPOINT ["./product-packaging","-path","./config/config.json"]

