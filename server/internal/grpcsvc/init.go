package grpcsvc

import (
	"github.com/Xanonymous-GitHub/carnival/db/pkg/ent/proto/entpb"
	. "github.com/Xanonymous-GitHub/carnival/server/pkg/vp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"strconv"
)

func init() {
	dbSvcAddr := ":" + strconv.Itoa(Vp.GetInt("dbServerPort"))

	// TODO(TU): remove insecure credentials when in production mode.
	devCredentials := insecure.NewCredentials()

	dial, err := grpc.Dial(dbSvcAddr, grpc.WithTransportCredentials(devCredentials))

	if err != nil {
		log.Fatalf("can not init grpc client!\n%v", err)
	}

	defer func(dial *grpc.ClientConn) {
		err := dial.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(dial)

	ApplicationSvcClient = entpb.NewApplicationServiceClient(dial)
}
