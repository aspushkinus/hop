FROM golang

ENV UPDATED_ON 2016-04-27

RUN apt-get update

RUN apt-get install -y \
  netcat \
  git

# RUN apt-get install -y \
#   xvfb xauth chromium chromium-l10n --no-install-recommends

# RUN apt-get install -y chromium

RUN go get github.com/sclevine/agouti
RUN go get github.com/onsi/ginkgo/ginkgo
RUN go get github.com/onsi/gomega
