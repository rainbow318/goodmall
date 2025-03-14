package mtl

import (
	"net"
	"net/http"

	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/server"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var Registry *prometheus.Registry // 相当于prometheus sdk提供的一个注册中心，用来注册指标

// 初始化普罗米欧斯的方法
// metricsPort 是metrics server监听的地址
// registryAddr 是注册中心的地址
func InitMetric(serviceName, metricsPort, registryAddr string) (registry.Registry, *registry.Info) {
	Registry = prometheus.NewRegistry()
	Registry.MustRegister(collectors.NewGoCollector())                                       // 注册go运行时相关的指标
	Registry.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{})) // 注册进程相关的指标
	// 把我们的服务注册到consul注册中心，这样当我们的服务启动时，普罗米欧斯就知道我们有哪些服务，从而不需要yaml一个个去配置服务
	r, _ := consul.NewConsulRegister(registryAddr)    // 忽略错误，因为这里一般不会发生错误
	addr, _ := net.ResolveTCPAddr("tcp", metricsPort) // 把我们传入的metricsPort转化为一个tcp地址
	registryInfo := &registry.Info{                   // 把我们的指标服务统一注册到一个serviceName的普罗米欧斯的consul服务里
		ServiceName: "prometheus",
		Addr:        addr,
		Weight:      1,
		Tags:        map[string]string{"service": serviceName},
	}
	_ = r.Register(registryInfo)

	server.RegisterShutdownHook(func() {
		r.Deregister(registryInfo)
	})
	// 启动metric server
	// 只采集了默认的指标数据
	// 被监控的应用程序提供一个/metrics的http接口，该接口返回当前系统的监控指标数据
	http.Handle("/metrics", promhttp.HandlerFor(Registry, promhttp.HandlerOpts{}))
	// 异步启动一个http服务来让prometheus拉取指标（这个http服务只是数据暴露通道）
	go http.ListenAndServe(metricsPort, nil)

	return r, registryInfo
}
