package delivery

import (
	delivery_grpc "github.com/godpeny/gos/delivery/grpc"
	"github.com/godpeny/gos/ent"
)

type GosServer struct {
	Server delivery_grpc.DeliveryServer
	DB     *ent.Client
}
