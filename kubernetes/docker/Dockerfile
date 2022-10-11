# The first stage will be the build-stage which uses the golang:latest image as a base.
FROM golang:1.19-alpine3.16 as builder
RUN mkdir /app
WORKDIR /app
COPY go-app/ .
RUN go mod download
RUN go mod verify
RUN go build -o docker

# The Second and final stage should be based on the apline:latest image.
FROM alpine:3.16.2
COPY --from=builder /app/docker .
EXPOSE 8080
CMD ["./docker"]