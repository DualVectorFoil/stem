package client

import (
	"github.com/Diode222/etcd_service_discovery/etcdservice"
	"github.com/DualVectorFoil/stem/app/conf"
	"github.com/DualVectorFoil/stem/pb"
	"sync"
)

var accountClient pb.AccountServiceClient
var accountClientOnce sync.Once

func NewAccountClient() pb.AccountServiceClient {
	accountClientOnce.Do(func() {
		accountClient = etcdservice.NewServiceManager(conf.ETCD_ADDRESS).GetClient(conf.SOLAR_SERVICE_NAME, pb.NewAccountClientWrapper).(pb.AccountServiceClient)
	})
	return accountClient
}
