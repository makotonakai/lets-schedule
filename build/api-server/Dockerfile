FROM golang:1.18

RUN mkdir /go/src/app
COPY ./src/ /go/src/app/

WORKDIR /go/src/app
RUN go install

EXPOSE 1323



