FROM golang:1.12.4-alpine
LABEL  maintainer "Dongri Jin <dongrify@gmail.com>"

RUN apk add --no-cache git
RUN go get -u github.com/githubnemo/CompileDaemon

ADD . /go/src/github.com/dongri/gonion
WORKDIR /go/src/github.com/dongri/gonion

CMD PORT=3001 CompileDaemon -command="./gonion"
EXPOSE 3001
