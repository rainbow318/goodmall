package mtl

import (
	"github.com/kitex-contrib/obs-opentelemetry/provider"
)

func InitTracing(serviceName string) provider.OtelProvider {
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(serviceName), // 配置service Name
		// provider.WithExportEndpoint("localhost:4317"), // 配置上报的端点。默认就是这个，所以这里就注释掉这句了
		provider.WithInsecure(),           // 这里我们是用自架的服务，使用非http进行上报
		provider.WithEnableMetrics(false), // 因为我们前面用了prometheus的指标，所以这里就不再用opentelementry的指标了
	)
	return p
}
