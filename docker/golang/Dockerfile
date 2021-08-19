FROM gcr.io/google.com/cloudsdktool/cloud-sdk:latest

SHELL ["/bin/bash", "-c"]

RUN apt-get update && apt-get install -y vim wget unzip

# golangをインストール
RUN wget https://golang.org/dl/go1.11.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.11.linux-amd64.tar.gz
RUN echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.bashrc
RUN source ~/.bashrc

COPY ./golang /app
WORKDIR /app
