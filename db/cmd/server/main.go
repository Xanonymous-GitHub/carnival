package main

import (
	"context"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	"fmt"
	"github.com/Xanonymous-GitHub/carnival/db/pkg/ent"
	"github.com/Xanonymous-GitHub/carnival/db/pkg/ent/migrate"
	"github.com/Xanonymous-GitHub/carnival/db/pkg/ent/proto/entpb"
	"github.com/Xanonymous-GitHub/carnival/db/pkg/vp"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)

var cvp *viper.Viper

func init() {
	cvp = vp.InitViper()
}

func main() {
	const sqlType = dialect.MySQL

	var (
		dbUserName = cvp.GetString("dbUserName")
		dbPassword = cvp.GetString("dbPassword")
		dbHost     = cvp.GetString("dbHost")
		dbPort     = cvp.GetInt("dbPort")
		dbName     = cvp.GetString("dbName")
	)

	serverAddress := ":" + strconv.Itoa(cvp.GetInt("dbServerPort"))
	dataSourceName := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?parseTime=true",
		dbUserName,
		dbPassword,
		dbHost,
		dbPort, dbName,
	)

	// Initialize an ent client.
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

	// Run the migration tool.
	ctx := context.Background()
	err = client.Schema.Create(
		ctx,
		schema.WithAtlas(true),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		migrate.WithGlobalUniqueID(true),
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// Initialize the generated services.
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

	// Register service with server.
	entpb.RegisterApplicationServiceServer(server, applicationSvc)
	entpb.RegisterApplicationAssignmentHistoryServiceServer(server, applicationAssignmentHistorySvc)
	entpb.RegisterApplicationStatusHistoryServiceServer(server, applicationStatusHistorySvc)
	entpb.RegisterAttachmentServiceServer(server, attachmentSvc)
	entpb.RegisterReviewerServiceServer(server, reviewerSvc)
	entpb.RegisterTicketServiceServer(server, ticketSvc)

	// Open port for listening to traffic.
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
