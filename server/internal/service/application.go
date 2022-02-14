package service

import (
	"context"
	"github.com/Xanonymous-GitHub/carnival/db/pkg/ent/proto/entpb"
	"log"
)

func CreateApplication(ctx context.Context, client entpb.ApplicationServiceClient, applicationData *entpb.Application) error {
	created, err := client.Create(ctx, &entpb.CreateApplicationRequest{Application: applicationData})
	if err != nil {
		LogGrpcErr(err, client)
		return err
	}

	get, err := client.Get(ctx, &entpb.GetApplicationRequest{Id: created.Id})
	if err != nil {
		LogGrpcErr(err, client)
		return err
	}

	log.Printf("application created, id=%d: %v", get.Id, get)
	return nil
}
