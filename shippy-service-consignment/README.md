
### rebuild go proto buffer code
```
protoc -I. --go_out=plugins=grpc:. \
	  proto/consignment/consignment.proto
```

### Docker build command

```
docker build -t shippy-service-consignment .
```

### Docker container command
```
docker run -p 50051:50051 shippy-service-consignment
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