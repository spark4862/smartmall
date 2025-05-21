package rpc

import (
	"sync"

	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/spark4862/smartmall/app/frontend/conf"
	frontendUtils "github.com/spark4862/smartmall/app/frontend/utils"
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
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)
	UserClient, err = userservice.NewClient("user", client.WithResolver(r))
	frontendUtils.MustHandleError(err)
}

func initProductClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)
	ProductClient, err = productcategoryservice.NewClient("product", client.WithResolver(r))
	frontendUtils.MustHandleError(err)
}

func initCartClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)
	CartClient, err = cartservice.NewClient("cart", client.WithResolver(r))
	frontendUtils.MustHandleError(err)
}

func initCheckoutClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)
	CheckoutClient, err = checkoutservice.NewClient("checkout", client.WithResolver(r))
	frontendUtils.MustHandleError(err)
}

func initOrderClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)
	OrderClient, err = orderservice.NewClient("order", client.WithResolver(r))
	frontendUtils.MustHandleError(err)
}
