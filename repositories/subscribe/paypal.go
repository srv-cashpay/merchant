package subscribe

import (
	"context"

	paypal "github.com/plutov/paypal/v4"
)

func (r *subscribeRepository) CreatePaypalOrder(amount, currency string) (*paypal.Order, error) {
	orderIntent := paypal.OrderIntentCapture

	purchaseUnit := paypal.PurchaseUnitRequest{
		Amount: &paypal.PurchaseUnitAmount{
			Currency: currency,
			Value:    amount,
		},
	}

	appContext := &paypal.ApplicationContext{
		ReturnURL: "https://cashpay.my.id/return",
		CancelURL: "https://cashpay.my.id/cancel",
	}

	return r.paypalClient.CreateOrder(
		context.Background(),
		orderIntent,
		[]paypal.PurchaseUnitRequest{purchaseUnit},
		nil,        // payment source (gunakan nil jika tidak spesifik)
		appContext, // context aplikasi (return/cancel URL)
	)
}

func (r *subscribeRepository) CapturePaypalOrder(orderID string) (*paypal.CaptureOrderResponse, error) {
	return r.paypalClient.CaptureOrder(context.Background(), orderID, paypal.CaptureOrderRequest{})
}
