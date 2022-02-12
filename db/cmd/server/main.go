package main

import (
	"context"
	"entgo.io/ent/dialect"
	"fmt"
	"github.com/Xanonymous-GitHub/carnival/db/pkg/ent"
	"github.com/Xanonymous-GitHub/carnival/db/pkg/ent/proto/entpb"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	const sqlType = dialect.MySQL

	// TODO(TU): use viper.
	const (
		dbUserName = "root"
		dbPassword = "root"
		dbHost     = "localhost"
		dbPort     = "3306"
	)

	const serverAddress = ":5431"

	dataSourceName := fmt.Sprintf("%v:%v@tcp(%v:%v)/carnival", dbUserName, dbPassword, dbHost, dbPort)

	// Initialize an ent client.
	// TODO(TU): specify the dataSourceName.
	client, err := ent.Open(sqlType, dataSourceName)
	if err != nil {
		log.Fatalf("failed opening connection to %v: %v", sqlType, err)
		return
	}

	defer func(client *ent.Client) {
		err := client.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(client)

	// Run the migration tool (creating tables, etc).
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// Initialize the generated User service.
	var (
		applicationSvc                  = entpb.NewApplicationService(client)
		applicationAssignmentHistorySvc = entpb.NewApplicationAssignmentHistoryService(client)
		applicationStatusHistorySvc     = entpb.NewApplicationStatusHistoryService(client)
		attachmentSvc                   = entpb.NewAttachmentService(client)
		reviewerSvc                     = entpb.NewReviewerService(client)
		ticketSvc                       = entpb.NewTicketService(client)
	)

	// Create a new gRPC server.
	server := grpc.NewServer()

	// Register the User service with the server.
	entpb.RegisterApplicationServiceServer(server, applicationSvc)
	entpb.RegisterApplicationAssignmentHistoryServiceServer(server, applicationAssignmentHistorySvc)
	entpb.RegisterApplicationStatusHistoryServiceServer(server, applicationStatusHistorySvc)
	entpb.RegisterAttachmentServiceServer(server, attachmentSvc)
	entpb.RegisterReviewerServiceServer(server, reviewerSvc)
	entpb.RegisterTicketServiceServer(server, ticketSvc)

	// Open port 5000 for listening to traffic.
	lis, err := net.Listen("tcp", serverAddress)
	if err != nil {
		log.Fatalf("failed listening: %s", err)
	}

	defer func(lis net.Listener) {
		err := lis.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(lis)

	// Listen for traffic indefinitely.
	if err := server.Serve(lis); err != nil {
		log.Fatalf("server ended: %s", err)
	}

	server.GracefulStop()
}
