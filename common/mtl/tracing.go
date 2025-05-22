package mtl

import (
	"github.com/kitex-contrib/obs-opentelemetry/provider"
)

func InitTracing(serviceName string) provider.OtelProvider {
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(serviceName),
		// 标识span的服务
		provider.WithInsecure(),
		provider.WithEnableMetrics(false),
		// provider.WithExportEndpoint(),
	)
	return p
}
