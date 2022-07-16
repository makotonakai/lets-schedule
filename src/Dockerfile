FROM golang:1.18

COPY ./ /app

WORKDIR /app
RUN go install

EXPOSE 1323

CMD [ "go", "run", "server.go" ]



