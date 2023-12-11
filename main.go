package main

import (
	"fmt"

	"github.com/lenny-mo/payment-api/handler"
	"github.com/lenny-mo/payment-api/proto/paymentapi"
	"github.com/lenny-mo/payment/proto/payment"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
)

func main() {
	// 注册中心
	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	// 创建服务
	service := micro.NewService(
		micro.Name("go.micro.api.payment-api"),
		micro.Version("latest"),
		micro.Address("127.0.0.1:8091"),
		micro.Registry(consulRegistry),
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
	go func() {
		if err := service.Run(); err != nil {
			fmt.Println(err)
		}
	}()
}
