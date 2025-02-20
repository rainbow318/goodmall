// Code generated by hertz generator.

package main

import (
	"context"
	"os"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/cors"
	"github.com/hertz-contrib/gzip"
	"github.com/hertz-contrib/logger/accesslog"
	hertzlogrus "github.com/hertz-contrib/logger/logrus"
	hertzprom "github.com/hertz-contrib/monitor-prometheus"
	hertzobslogrus "github.com/hertz-contrib/obs-opentelemetry/logging/logrus"
	hertztracing "github.com/hertz-contrib/obs-opentelemetry/tracing"
	"github.com/hertz-contrib/pprof"
	"github.com/hertz-contrib/sessions"
	"github.com/hertz-contrib/sessions/redis"
	"github.com/joho/godotenv"
	"github.com/suutest/app/frontend/biz/router"
	"github.com/suutest/app/frontend/conf"
	"github.com/suutest/app/frontend/infra/rpc"
	"github.com/suutest/app/frontend/middleware"
	frontendutils "github.com/suutest/app/frontend/utils"
	"github.com/suutest/common/mtl"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	ServiceName  = frontendutils.ServiceName
	MetricsPort  = conf.GetConf().Hertz.MetricsPort
	RegistryAddr = conf.GetConf().Hertz.RegistryAddr
)

func main() {
	_ = godotenv.Load()
	p := mtl.InitTracing(ServiceName)
	defer p.Shutdown(context.Background()) // 在服务关闭前，将剩余的链路数据都上传完
	consul, registryInfo := mtl.InitMetric(ServiceName, MetricsPort, RegistryAddr)
	defer consul.Deregister(registryInfo) // 这样hertz停止服务时就可以撤掉prometheus上的实例
	// init dal
	// dal.Init()
	rpc.Init()
	address := conf.GetConf().Hertz.Address // 从配置中获取监听地址

	tracer, cfg := hertztracing.NewServerTracer()
	h := server.New(server.WithHostPorts(address), // 创建服务器实例
		server.WithTracer(hertzprom.NewServerTracer("", "", hertzprom.WithDisableServer(true),
			hertzprom.WithRegistry(mtl.Registry),
		)),
		tracer,
	)
	h.Use(hertztracing.ServerMiddleware(cfg)) // 集成hertz的opentelementry中间件

	registerMiddleware(h) // 注册中间件

	// add a ping route to test
	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{"ping": "pong"})
	})

	router.GeneratedRegister(h) // 根路径"/"的处理逻辑注册到服务器上，因此当浏览器访问时会触发该逻辑并返回响应
	h.LoadHTMLGlob("template/*")

	h.Static("/static", "./")

	h.GET("/about", func(c context.Context, ctx *app.RequestContext) {
		hlog.CtxInfof(c, "GoodMall about page")
		ctx.HTML(consts.StatusOK, "about", utils.H{"Title": "About"})
	})

	// 定义一个GET请求路由/sign-in，当用户访问这个路径时，服务器渲染一个名为sign-in的模版，并传递一个Title参数给模版（参数值为"Sign In"）
	h.GET("/sign-in", func(c context.Context, ctx *app.RequestContext) {
		data := utils.H{
			"Title": "Sign In",
			"Next":  ctx.Query("next"),
		}
		ctx.HTML(consts.StatusOK, "sign-in", data) // 渲染tmpl文件
	})
	h.GET("/sign-up", func(c context.Context, ctx *app.RequestContext) {
		ctx.HTML(consts.StatusOK, "sign-up", utils.H{"Title": "Sign Up"})
	})

	h.Spin() // 启动服务器并开始监听请求
}

func registerMiddleware(h *server.Hertz) {
	store, _ := redis.NewStore(10, "tcp", conf.GetConf().Redis.Address, "", []byte(os.Getenv("SESSION_SECRET"))) // 依赖redies的session管理
	h.Use(sessions.New("GoodMall-shop", store))

	// log
	logger := hertzobslogrus.NewLogger(hertzobslogrus.WithLogger(hertzlogrus.NewLogger().Logger()))
	hlog.SetLogger(logger)
	hlog.SetLevel(conf.LogLevel())
	var flushInterval time.Duration
	if os.Getenv("GO_ENV") == "online" {
		flushInterval = time.Minute
	} else {
		flushInterval = time.Second
	}
	asyncWriter := &zapcore.BufferedWriteSyncer{ // 默认使用异步发盘方式
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.GetConf().Hertz.LogFileName,
			MaxSize:    conf.GetConf().Hertz.LogMaxSize,
			MaxBackups: conf.GetConf().Hertz.LogMaxBackups,
			MaxAge:     conf.GetConf().Hertz.LogMaxAge,
		}),
		FlushInterval: flushInterval,
	}
	hlog.SetOutput(asyncWriter)
	h.OnShutdown = append(h.OnShutdown, func(ctx context.Context) {
		asyncWriter.Sync()
	})

	// pprof
	if conf.GetConf().Hertz.EnablePprof {
		pprof.Register(h)
	}

	// gzip
	if conf.GetConf().Hertz.EnableGzip {
		h.Use(gzip.Gzip(gzip.DefaultCompression))
	}

	// access log
	if conf.GetConf().Hertz.EnableAccessLog {
		h.Use(accesslog.New())
	}

	// recovery
	h.Use(recovery.Recovery())

	// cores
	h.Use(cors.Default())

	middleware.Register(h)
}
