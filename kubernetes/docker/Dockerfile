FROM golang:1.19-alpine3.16

RUN mkdir /app
WORKDIR /app

COPY go-app/ .

RUN go mod download
RUN go mod verify
RUN go build -o docker

EXPOSE 8080

CMD ["./docker"]