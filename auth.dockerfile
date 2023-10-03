FROM golang:latest

WORKDIR /app

COPY ./VF-Server /app

RUN go build -o Auth-Server

EXPOSE 8080

CMD ["./Auth-Server"]
