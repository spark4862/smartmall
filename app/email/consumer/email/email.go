package email

import (
	"context"
	"fmt"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"github.com/nats-io/nats.go"
	"github.com/spark4862/smartmall/app/email/intra/mq"
	"github.com/spark4862/smartmall/app/email/intra/notify"
	emailRpc "github.com/spark4862/smartmall/rpc_gen/kitex_gen/email"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"google.golang.org/protobuf/proto"
)

func ConsumerInit() {
	tracer := otel.Tracer("shop-nats-consumer")
	sub, err := mq.Nc.Subscribe("email", func(m *nats.Msg) {
		var req emailRpc.EmailReq
		err := proto.Unmarshal(m.Data, &req)
		if err != nil {
			klog.Error(err)
			return
		}

		ctx := context.Background()
		ctx = otel.GetTextMapPropagator().Extract(ctx, propagation.HeaderCarrier(m.Header))
		_, span := tracer.Start(ctx, "shop-email-consumer")

		defer span.End()

		noopEmail := notify.NewNoopEmail()
		err = noopEmail.Send(&req)
		if err != nil {
			klog.Error(err)
			fmt.Println(err)
			return
		}
		fmt.Println(req.Content)
	})
	if err != nil {
		klog.Fatal(err)
		fmt.Println(err)
		panic(err)
	}

	server.RegisterShutdownHook(func() {
		_ = sub.Unsubscribe()
		mq.Nc.Close()
	})
}
