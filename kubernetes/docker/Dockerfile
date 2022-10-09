FROM golang:1.19-alpine3.16

RUN mkdir /app
WORKDIR /app

RUN export GO111MODULE=on
RUN go install github.com/anynines/tmp-homework-sh/kubernetes/docker/go-app@latest
RUN cd /app && https://github.com/anynines/tmp-homework-sh.git

RUN cd /app/tmp-homework-sh/kubernetes/docker/go-app && go build

EXPOSE 8080

ENTRYPOINT ["/app/tmp-homework-sh/kubernetes/docker/go-app/main"]