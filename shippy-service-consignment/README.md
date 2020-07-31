### Docker build command

```
docker build -t shippy-service-consignment .
```

### Docker container command
```
docker run -p 50051:50051 shippy-service-consignment

docker run -p 50051:50051 \
      -e MICRO_SERVER_ADDRESS=:50051 \
      shippy-service-consignment
```
### Install dependency

```
go get github.com/micro/micro/v2
go get github.com/micro/micro/v2/cmd/protoc-gen-micro@master

```
### Protoc buffer generate code command
```
protoc --proto_path=. --go_out=. --micro_out=. \
		proto/consignment/consignment.proto
```