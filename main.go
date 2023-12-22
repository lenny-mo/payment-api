package main

import (
	"fmt"

	"github.com/lenny-mo/emall-utils/tracer"
	"github.com/lenny-mo/payment-api/circuit"
	"github.com/lenny-mo/payment-api/handler"
	"github.com/lenny-mo/payment-api/proto/paymentapi"
	"github.com/lenny-mo/payment/proto/payment"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
)

func main() {
	// 注册中心
	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	serviceName := "go.micro.api.payment-api"
	// 开启链路追踪
	err := tracer.InitTracer(serviceName, "127.0.0.1:6831")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer tracer.Closer.Close()
	opentracing.SetGlobalTracer(tracer.Tracer)

	// 创建服务
	service := micro.NewService(
		micro.Name(serviceName),
		micro.Version("latest"),
		micro.Address("127.0.0.1:8091"),
		micro.Registry(consulRegistry),
		// 上下游添加链路追踪
		micro.WrapClient(opentracing2.NewClientWrapper(opentracing.GlobalTracer())),
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
		// 添加请求微服务时候的熔断机制
		micro.WrapClient(circuit.NewClientWrapper()),
	)

	service.Init()

	paymentService := payment.NewPaymentService("go.micro.service.payment", service.Client())
	// 注册服务
	if err := paymentapi.RegisterPaymentAPIHandler(service.Server(), &handler.PaymentAPI{
		PaymentService: paymentService,
	}); err != nil {
		fmt.Println(err)
	}

	// 启动服务
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}

}
