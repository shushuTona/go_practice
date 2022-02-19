FROM golang:1.17.7-alpine

RUN apk update && apk add git

RUN apk --no-cache add tzdata && \
            cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
            apk del tzdata
