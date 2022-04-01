FROM golang:latest

LABEL maintainer=""

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

ENV PORT 8081

RUN go build

CMD ["./Go-task"]