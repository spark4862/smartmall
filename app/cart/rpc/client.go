package rpc

import (
	"sync"

	"github.com/cloudwego/kitex/client"
	"github.com/spark4862/smartmall/app/cart/conf"
	cartutils "github.com/spark4862/smartmall/app/cart/utils"
	"github.com/spark4862/smartmall/common/clientsuite"
	"github.com/spark4862/smartmall/rpc_gen/kitex_gen/product/productcategoryservice"
)

var (
	ProductClient productcategoryservice.Client
	once          sync.Once
	ServiceName   = conf.GetConf().Kitex.Service
	RegisterAddr  = conf.GetConf().Registry.RegistryAddress[0]
	err           error
)

func Init() {
	once.Do(func() {
		initProductClient()
	})
}

func initProductClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegisterAddr,
		}),
	}

	ProductClient, err = productcategoryservice.NewClient("product", opts...)
	cartutils.MustHandleError(err)
}
