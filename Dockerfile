FROM golang:latest

WORKDIR /app

COPY ./VF-Server /app

RUN go build -ldflags="-s" Auth-Server

EXPOSE 8050

CMD ["./Auth-Server"]
