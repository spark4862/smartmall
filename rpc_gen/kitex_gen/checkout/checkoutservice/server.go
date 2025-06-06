// Code generated by Kitex v0.9.1. DO NOT EDIT.
package checkoutservice

import (
	server "github.com/cloudwego/kitex/server"
	checkout "github.com/spark4862/smartmall/rpc_gen/kitex_gen/checkout"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler checkout.CheckoutService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}

func RegisterService(svr server.Server, handler checkout.CheckoutService, opts ...server.RegisterOption) error {
	return svr.RegisterService(serviceInfo(), handler, opts...)
}
