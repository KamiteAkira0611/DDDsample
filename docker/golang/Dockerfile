FROM gcr.io/google.com/cloudsdktool/cloud-sdk:latest

SHELL ["/bin/bash", "-c"]

COPY ./golang /app
WORKDIR /app

RUN apt-get update && apt-get install -y vim wget unzip

# golangをインストール
RUN wget https://golang.org/dl/go1.11.13.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.11.13.linux-amd64.tar.gz
ENV goPATH /usr/local/go/bin
ENV GOPATH /go
ENV GOBIN /go/bin

# twipをインストール
# 参考 https://twitchtv.github.io/twirp/docs/install.html
RUN apt-get install -y protobuf-compiler
RUN /usr/local/go/bin/go install google.golang.org/protobuf/cmd/protoc-gen-go
RUN /usr/local/go/bin/go install github.com/twitchtv/twirp/protoc-gen-twirp

# パスを通す
RUN echo "export PATH=$GOBIN:$PATH:$goPATH" >> ~/.bashrc
RUN source ~/.bashrc
