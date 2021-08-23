FROM golang:1.13.15-alpine as builder

RUN apk update && apk upgrade && \
    apk add --no-cache git

RUN mkdir /app
WORKDIR /app

ENV GO111MODULE=on

COPY . .
#RUN go get github.com/golang/protobuf/proto@v1.5.2
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o vessel-service

# Run container
FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/vessel-service .

CMD ["./vessel-service"]