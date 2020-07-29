module github.com/galahade/shippy/shippy-service-vessel

go 1.14

replace github.com/galahade/shippy/shippy-service-vessel => ./

require (
	github.com/golang/protobuf v1.4.0
	github.com/micro/go-micro/v2 v2.9.1
	golang.org/x/text v0.3.3 // indirect
)
