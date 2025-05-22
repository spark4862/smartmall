package mtl

import (
	"fmt"
	"net"
	"net/http"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/server"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var Registry *prometheus.Registry

func InitMetric(serviceName, metricsPort, registryAddr string) (registry.Registry, *registry.Info) {
	Registry = prometheus.NewRegistry()
	// 建一个新的 Prometheus Registry，它是存储和管理所有指标的容器
	Registry.MustRegister(collectors.NewGoCollector())
	// 注册一个 Go 运行时的指标收集器，它会收集如 goroutine 数量、内存使用情况、CPU 使用情况等有关 Go 运行时的信息。
	Registry.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))
	// 注册一个进程级的指标收集器，收集当前进程的相关指标，如 进程 ID、内存、CPU 使用情况等。
	r, err := consul.NewConsulRegister(registryAddr)
	if err != nil {
		klog.Fatal(err)
		fmt.Println(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", metricsPort)
	if err != nil {
		klog.Fatal(err)
		fmt.Println(err)
	}
	registryInfo := &registry.Info{
		ServiceName: "prometheus",
		Addr:        addr,
		Weight:      1,
		Tags:        map[string]string{"service": serviceName},
	}

	err = r.Register(registryInfo)
	if err != nil {
		klog.Fatal(err)
		fmt.Println(err)
		panic(err)
	}
	server.RegisterShutdownHook(func() {
		r.Deregister(registryInfo)
	})

	http.Handle("/metrics", promhttp.HandlerFor(Registry, promhttp.HandlerOpts{}))
	go http.ListenAndServe(metricsPort, nil)
	return r, registryInfo
	// 用于非kitex框架无法加钩子的情况
}
