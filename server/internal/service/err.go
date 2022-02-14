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
		log.Printf("can not create %s!\nstatus=%s message=%s", name, se.Code(), se.Message())
	} else {
		log.Printf("can not create %s!", name)
	}
}
