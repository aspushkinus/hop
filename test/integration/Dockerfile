FROM nimmis/alpine-golang

ENV UPDATED_ON 2016-04-27

RUN apk update

RUN apk add \
  bash \
  netcat-openbsd \
  git

RUN go get github.com/sclevine/agouti
RUN go get github.com/onsi/ginkgo/ginkgo
RUN go get github.com/onsi/gomega
