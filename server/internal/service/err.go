package service

import (
	"google.golang.org/grpc/status"
	"log"
	"reflect"
)

func LogGrpcErr(err error, svc interface{}) {
	se, ok := status.FromError(err)
	name := reflect.TypeOf(svc).String()
	if ok {
		log.Printf("%s error!\nstatus=%s message=%s", name, se.Code(), se.Message())
	} else {
		log.Printf("%s error!", name)
	}
}
