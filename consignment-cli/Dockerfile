FROM golang:1.13.15-alpine as builder

RUN apk update && apk upgrade && \
    apk add --no-cache git

RUN mkdir /app
WORKDIR /app

ENV GO111MODULE=on

COPY . .
#RUN go get github.com/golang/protobuf/proto@v1.5.2
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o consignment-cli


FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
ADD consignment.json /app/consignment.json
COPY --from=builder /app/consignment-cli .

CMD ["./consignment-cli"]