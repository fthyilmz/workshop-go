FROM golang:1.15 as builder

RUN apt-get update  -y && apt-get install -y --no-install-recommends apt-utils bzip2

WORKDIR /app
COPY ./ /app

RUN go mod download
#RUN go run seeder/seeder.go
RUN go build -o main .

#FROM scratch

#WORKDIR /app

#COPY --from=builder . .

EXPOSE 8080

ENTRYPOINT ["tail -f /dev/null"]