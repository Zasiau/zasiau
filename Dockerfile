FROM golang:1.12.4-alpine
LABEL  maintainer "Dongri Jin <dongrify@gmail.com>"

RUN apk add --update alpine-sdk

# ENV
ARG go_env
ENV GO_ENV $go_env

ADD . /go/src/github.com/dongri/gonion
WORKDIR /go/src/github.com/dongri/gonion
RUN go install github.com/dongri/gonion

CMD ["/go/bin/gonion"] 
EXPOSE 3001
