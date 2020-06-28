FROM golang:1.11.10-alpine3.9 as builder

# installing git
RUN apk update && apk upgrade && \
    apk add --no-cache git

# setting working directory
WORKDIR /go/src/app

# installing dependencies
RUN go get github.com/sirupsen/logrus
RUN go get github.com/buaazp/fasthttprouter
RUN go get github.com/valyala/fasthttp

COPY / /go/src/app/
RUN go build -o mockit

FROM alpine:3.9

WORKDIR /go/src/app
COPY --from=builder /go/src/app/mockit /go/src/app/mockit

COPY *.json /app/response/
COPY *.xml /app/response/

EXPOSE 8010

CMD ["./mockit"]