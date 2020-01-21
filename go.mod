module github.com/newmind/shippy-service-consignment

go 1.13

replace github.com/newmind/shippy-service-vessel => ../shippy-service-vessel

require (
	github.com/golang/protobuf v1.3.2
	github.com/micro/go-micro v1.18.0
	github.com/newmind/shippy-service-vessel v1.2.5
	github.com/pkg/errors v0.8.1
	github.com/sirupsen/logrus v1.4.2
	go.mongodb.org/mongo-driver v1.2.1
	golang.org/x/net v0.0.0-20200114155413-6afb5195e5aa
)
