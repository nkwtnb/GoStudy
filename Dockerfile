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
# パッケージのインポート
# RUN go get -u golang.org/x/tools/cmd/goimports