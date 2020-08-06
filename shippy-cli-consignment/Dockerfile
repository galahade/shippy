FROM golang:alpine as builder

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io,direct

RUN apk update && apk upgrade && \
    apk add --no-cache git

WORKDIR /app
COPY go.mod go.sum /app/
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o cli *.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app
ADD consignment.json /app/consignment.json
COPY --from=builder /app/cli /app/cli

CMD ["./cli"]