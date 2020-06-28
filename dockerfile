FROM golang:1.14

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor
ENV APP_USER app
ENV APP_HOME /go/src/mockit

# setting working directory
WORKDIR /go/src/app

# installing dependencies
RUN go get github.com/sirupsen/logrus
RUN go get github.com/buaazp/fasthttprouter
RUN go get github.com/valyala/fasthttp

COPY / /go/src/app/
RUN go build -o mockit

EXPOSE 8010

CMD ["./mockit"]