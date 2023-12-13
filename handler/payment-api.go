package handler

import (
	"context"
	"fmt"

	"github.com/lenny-mo/payment-api/proto/paymentapi"
	"github.com/lenny-mo/payment/proto/payment"
)

// 实现下面的接口
// server api
//
//	type PaymentAPIHandler interface {
//		MakePayment(context.Context, *MakePaymentRequest, *MakePaymentResponse) error
//		GetPayment(context.Context, *GetPaymentRequest, *GetPaymentResponse) error
//		UpdatePayment(context.Context, *UpdatePaymentRequest, *UpdatePaymentResponse) error
//	}
type PaymentAPI struct {
	payment.PaymentService
}

func (p *PaymentAPI) MakePayment(ctx context.Context, req *paymentapi.MakePaymentRequest, res *paymentapi.MakePaymentResponse) error {

	// 调用领域层的方法
	payreq := payment.MakePaymentRequest{
		OrderId: req.OrderId,
	}

	// 目前只是处理paypal订单
	switch req.Method {
	case "paypal":

		payres, err := p.PaymentService.MakePayment(ctx, &payreq)
		if err != nil {
			fmt.Println(err)
			res.Msg = err.Error()
			return err
		}

		res.Msg = "success " + "paymentId: " + payres.PaymentID
		res.PaymentId = payres.PaymentID
	}

	return nil
}

func (p *PaymentAPI) GetPayment(ctx context.Context, req *paymentapi.GetPaymentRequest, res *paymentapi.GetPaymentResponse) error {
	data, err := p.PaymentService.GetPaymentStatus(context.TODO(), &payment.GetPaymentStatusRequest{
		PaymentId: req.PaymentId,
	})

	if err != nil {
		fmt.Println("err during get payment")
		return err
	}

	res.PaymentInfo = "payment_method: " + data.PaymentData.PaymentMethod + "," +
		"payment_id: " + data.PaymentData.TransactionId

	return nil
}

func (p *PaymentAPI) UpdatePayment(ctx context.Context, req *paymentapi.UpdatePaymentRequest, res *paymentapi.UpdatePaymentResponse) error {
	data, err := p.PaymentService.UpdatePayment(context.TODO(), &payment.UpdatePaymentRequest{
		PaymentData: &payment.Payment{
			TransactionId:     req.PaymentId,
			PaymentMethod:     req.PaymentMethod,
			TransactionStatus: req.PaymentStatus,
		},
	})
	if err != nil {
		fmt.Println(err)
		return err
	}

	res.Code = "success"
	res.Msg = "payment_id: " + data.PaymentId

	return nil
}
