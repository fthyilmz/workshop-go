FROM golang:1.15 as builder

RUN apt-get update  -y && apt-get install -y --no-install-recommends apt-utils bzip2

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o main .

EXPOSE 8080

ENTRYPOINT ["/app/bin/start.sh"]