package rpc

import (
	"sync"

	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/spark4862/smartmall/app/cart/conf"
	cartutils "github.com/spark4862/smartmall/app/cart/utils"
	"github.com/spark4862/smartmall/rpc_gen/kitex_gen/product/productcategoryservice"
)

var (
	ProductClient productcategoryservice.Client
	once          sync.Once
)

func Init() {
	once.Do(func() {
		initProductClient()
	})
}

func initProductClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	cartutils.MustHandleError(err)
	ProductClient, err = productcategoryservice.NewClient("product", client.WithResolver(r))
	cartutils.MustHandleError(err)
}
