module github.com/galahade/shippy/shippy-cli-consignment

go 1.14

replace github.com/galahade/shippy/shippy-service-consignment => ../shippy-service-consignment

require (
	github.com/galahade/shippy/shippy-service-consignment v0.0.0-00010101000000-000000000000
	github.com/micro/go-micro/v2 v2.9.1
)
