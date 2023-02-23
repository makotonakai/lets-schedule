FROM golang:1.18

WORKDIR /app
COPY ./ /app
RUN go install

EXPOSE 1323

CMD [ "go", "run", "server.go" ]



