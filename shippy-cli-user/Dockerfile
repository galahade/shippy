FROM golang:alpine as builder

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io,direct
# Use this flow to make download packages as little as possible.

WORKDIR /app
COPY go.mod go.sum /app/
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o cli *.go

# Run container
FROM alpine:latest

#RUN apk --no-cache add ca-certificates
COPY --from=builder /app/cli /cli

ENTRYPOINT ["/cli"]