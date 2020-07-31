### Docker command
```
docker build -t shippy-service-vessel .
docker run shippy-service-vessel
```

### rebuild proto buffer code
```
protoc --proto_path=. --go_out=. --micro_out=. \
	  proto/vessel/vessel.proto
```