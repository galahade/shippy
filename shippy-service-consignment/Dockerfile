FROM golang:alpine as builder

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io,direct

RUN apk update && apk upgrade && \
    apk add --no-cache git

WORKDIR /app
COPY go.mod go.sum /app/
RUN go mod download

COPY . .

#RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o shippy-service-consignment
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o service *.go
# Run container
FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/service .

CMD ["./service"]
