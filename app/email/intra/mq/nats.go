package mq

import (
	"fmt"

	"github.com/cloudwego/kitex/pkg/klog"
	nats "github.com/nats-io/nats.go"
)

var (
	Nc  *nats.Conn
	err error
)

func Init() {
	Nc, err = nats.Connect(nats.DefaultURL)
	if err != nil {
		klog.Fatal(err)
		fmt.Println(err)
		panic(err)
	}
}
