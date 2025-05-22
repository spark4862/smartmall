package rpc

import (
	"sync"

	"github.com/cloudwego/kitex/client"
	"github.com/spark4862/smartmall/app/checkout/conf"
	"github.com/spark4862/smartmall/common/clientsuite"
	"github.com/spark4862/smartmall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/spark4862/smartmall/rpc_gen/kitex_gen/order/orderservice"
	"github.com/spark4862/smartmall/rpc_gen/kitex_gen/payment/paymentservice"
	"github.com/spark4862/smartmall/rpc_gen/kitex_gen/product/productcategoryservice"
)

var (
	CartClient    cartservice.Client
	ProductClient productcategoryservice.Client
	PaymentClient paymentservice.Client
	OrderClient   orderservice.Client
	once          sync.Once
	err           error
	ServiceName   = conf.GetConf().Kitex.Service
	RegisterAddr  = conf.GetConf().Registry.RegistryAddress[0]
)

func Init() {
	once.Do(func() {
		initCartClient()
		initProductClient()
		initPaymentClient()
		initOrderClient()
	})
}

func initCartClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegisterAddr,
		}),
	}

	CartClient, err = cartservice.NewClient("cart", opts...)
	if err != nil {
		panic(err)
	}
}

func initProductClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegisterAddr,
		}),
	}

	ProductClient, err = productcategoryservice.NewClient("product", opts...)
	if err != nil {
		panic(err)
	}
}

func initPaymentClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegisterAddr,
		}),
	}

	PaymentClient, err = paymentservice.NewClient("payment", opts...)
	if err != nil {
		panic(err)
	}
}

func initOrderClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegisterAddr,
		}),
	}

	OrderClient, err = orderservice.NewClient("order", opts...)
	if err != nil {
		panic(err)
	}
}
