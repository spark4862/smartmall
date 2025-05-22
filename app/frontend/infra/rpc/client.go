package rpc

import (
	"sync"

	"github.com/cloudwego/kitex/client"
	"github.com/spark4862/smartmall/app/frontend/conf"
	frontendUtils "github.com/spark4862/smartmall/app/frontend/utils"
	"github.com/spark4862/smartmall/common/clientsuite"
	"github.com/spark4862/smartmall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/spark4862/smartmall/rpc_gen/kitex_gen/checkout/checkoutservice"
	"github.com/spark4862/smartmall/rpc_gen/kitex_gen/order/orderservice"
	"github.com/spark4862/smartmall/rpc_gen/kitex_gen/product/productcategoryservice"
	"github.com/spark4862/smartmall/rpc_gen/kitex_gen/user/userservice"
)

var (
	UserClient     userservice.Client
	ProductClient  productcategoryservice.Client
	CartClient     cartservice.Client
	CheckoutClient checkoutservice.Client
	OrderClient    orderservice.Client
	once           sync.Once
	ServiceName    = frontendUtils.ServiceName
	RegisterAddr   = conf.GetConf().Hertz.RegistryAddr
	err            error
)

func Init() {
	once.Do(func() {
		initUserClient()
		initProductClient()
		initCartClient()
		initCheckoutClient()
		initOrderClient()
	})
}

func initUserClient() {
	UserClient, err = userservice.NewClient(
		"user",
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegisterAddr,
		}),
	)
	frontendUtils.MustHandleError(err)
}

func initProductClient() {
	ProductClient, err = productcategoryservice.NewClient(
		"product",
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegisterAddr,
		}),
	)
	frontendUtils.MustHandleError(err)
}

func initCartClient() {
	CartClient, err = cartservice.NewClient(
		"cart",
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegisterAddr,
		}),
	)
	frontendUtils.MustHandleError(err)
}

func initCheckoutClient() {
	CheckoutClient, err = checkoutservice.NewClient(
		"checkout",
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegisterAddr,
		}),
	)
	frontendUtils.MustHandleError(err)
}

func initOrderClient() {
	OrderClient, err = orderservice.NewClient(
		"order",
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegisterAddr,
		}),
	)
	frontendUtils.MustHandleError(err)
}
