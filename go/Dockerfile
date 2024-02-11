FROM golang:alpine3.19

WORKDIR /app

ARG SERVICE

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN rm -f .env

RUN go build -o app ./cmd/cli/main.go

RUN chmod +x ./app

EXPOSE 80

