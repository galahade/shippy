module google.golang.org

go 1.14

replace (
	github.com/galahade/shippy/shippy-service-consignment => ./
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)

require (
	github.com/galahade/shippy/shippy-service-consignment v0.0.0-00010101000000-000000000000
	github.com/golang/protobuf v1.4.2
	github.com/micro/go-micro/v2 v2.9.1
	golang.org/x/text v0.3.3 // indirect
)
