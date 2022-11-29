# goバージョン
FROM golang:1.19.1-alpine
# アップデートとgitのインストール
# alpineイメージはスリムだが、gitなどのツールも入っていないので入れる
RUN apk update && apk add git
# boiler-plateディレクトリの作成
RUN mkdir -p /go/src/github.com/boiler-plate
# ワーキングディレクトリの設定
WORKDIR /go/src/github.com/boiler-plate
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /go/src/github.com/boiler-plate
COPY .env /go/src/github.com/boiler-plate
# Databaseマイグレーション用ツールインストール
RUN go install github.com/pressly/goose/v3/cmd/goose@latest