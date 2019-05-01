FROM golang:1.12.4-alpine
LABEL  maintainer "Dongri Jin <dongrify@gmail.com>"

RUN apk add --update alpine-sdk

# ENV
ARG go_env
ENV GO_ENV $go_env

ADD . /go/src/github.com/dongri/candle
WORKDIR /go/src/github.com/dongri/candle
RUN go install github.com/dongri/candle

CMD ["/go/bin/candle"] 
EXPOSE 3001
