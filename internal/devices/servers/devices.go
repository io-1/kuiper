package servers

import (
	"github.com/io-1/kuiper/internal/devices/persistence"
	"github.com/io-1/kuiper/internal/devices/pubsub/publisher"

	devices_pb "github.com/io-1/kuiper/internal/pb/devices"
)

type DevicesServer struct {
	persistence persistence.Persistence
	devices_pb.UnimplementedDevicesServiceServer
	publisher publisher.Publisher
}

func NewDevicesServer(persistence persistence.Persistence, publisher publisher.Publisher) *DevicesServer {
	return &DevicesServer{
		persistence: persistence,
		publisher:   publisher,
	}
}
