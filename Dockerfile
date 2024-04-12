# syntax=docker/dockerfile:1

FROM golang:1.21-alpine
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY *.go ./
RUN go build -o /main
EXPOSE 12318
CMD [ "/main" ]