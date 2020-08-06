module github.com/galahade/shippy/shippy-service-vessel

go 1.14

replace github.com/galahade/shippy/shippy-service-vessel => ./

require (
	github.com/golang/protobuf v1.4.2
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/micro/v2 v2.9.3 // indirect
	github.com/tidwall/pretty v1.0.1 // indirect
	go.mongodb.org/mongo-driver v1.4.0
	golang.org/x/text v0.3.3 // indirect
)
