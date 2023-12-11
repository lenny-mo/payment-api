package handler

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/lenny-mo/payment-api/proto/paymentapi"
	"github.com/lenny-mo/payment/proto/payment"
)

// 实现下面的接口
// type PaymentAPIHandler interface {
// 	MakePayment(context.Context, *MakePaymentRequest, *MakePaymentResponse) error
// 	GetPayment(context.Context, *GetPaymentRequest, *GetPaymentResponse) error
// }

type PaymentAPI struct {
	payment.PaymentService
}

func (p *PaymentAPI) MakePayment(ctx context.Context, req *paymentapi.MakePaymentRequest, res *paymentapi.MakePaymentResponse) error {
	if _, ok := req.Get["order_id"]; !ok {
		// todo: 添加zap日志
		fmt.Println("order_id is not exist")
		return errors.New("order_id is not exist")
	}
	orderIdstr := req.Get["order_id"].Values[0]
	orderId, err := strconv.ParseInt(orderIdstr, 10, 64)
	if err != nil {
		fmt.Println(err)
		return err
	}

	payreq := payment.MakePaymentRequest{
		OrderId: orderId,
	}

	payres, err := p.PaymentService.MakePayment(context.TODO(), &payreq)
	if err != nil {
		fmt.Println(err)
		res.Msg = err.Error()
		return err
	}

	res.Msg = "success"
	res.Body = "paymentId:" + payres.PaymentID

	return nil
}

func (p *PaymentAPI) GetPayment(ctx context.Context, req *paymentapi.GetPaymentRequest, res *paymentapi.GetPaymentResponse) error {
	data, err := p.PaymentService.GetPaymentStatus(context.TODO(), &payment.GetPaymentStatusRequest{
		PaymentId: "12",
	})

	if err != nil {
		fmt.Println("err during get payment")
		return err
	}

	res.PaymentInfo = data.PaymentData.PaymentMethod + data.PaymentData.TransactionId

	return nil
}
