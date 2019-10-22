FROM golang:1.9

MAINTAINER Eranga Bandara (erangaeb@gmail.com)

# install packages
RUN apt-get update
RUN apt-get install unzip

# install protoc
env PROTOC_ZIP=protoc-3.7.1-linux-x86_64.zip
RUN curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.7.1/$PROTOC_ZIP
RUN unzip -o $PROTOC_ZIP -d /usr/local bin/protoc
RUN unzip -o $PROTOC_ZIP -d /usr/local 'include/*'
RUN rm -f $PROTOC_ZIP

# install go dependencies
RUN	go get github.com/gorilla/mux
RUN go get github.com/gorilla/handlers
RUN go get -a github.com/golang/protobuf/protoc-gen-go
RUN go get github.com/streadway/amqp

# copy app
ADD . /app
WORKDIR /app

# add protobuf spec
RUN git clone https://gitlab.com/rahasak-labs/rabbit.git
RUN protoc rabbit/spec/*.proto --go_out=:$GOPATH/src
RUN rm -rf rabbit

# build
RUN go build -o build/gateway src/*.go

# server running port
EXPOSE 7654

ENTRYPOINT ["/app/docker-entrypoint.sh"]
