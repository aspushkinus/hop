FROM golang

ENV UPDATED_ON 2016-04-27

RUN apt-get update

RUN apt-get install -y \
  netcat \
  git

RUN go get github.com/sclevine/agouti
RUN go get github.com/onsi/ginkgo/ginkgo
RUN go get github.com/onsi/gomega
