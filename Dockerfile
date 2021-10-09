# ベースにつかうイメージの指定
FROM golang:1.17 AS builder

WORKDIR /opt

COPY go.mod .
COPY go.sum .
COPY main.go .

RUN go build && cp go-count-webapp /usr/local/bin/go-count-webapp

# マルチステージビルド
# 軽量Linuxイメージ上で、できあがったバイナリだけをうごかす
FROM gcr.io/distroless/cc-debian10 AS runtime

COPY --from=builder /usr/local/bin/go-count-webapp /usr/local/bin/go-count-webapp

ENTRYPOINT ["/usr/local/bin/go-count-webapp"]
