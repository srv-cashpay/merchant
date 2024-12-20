package packages

import (
	"github.com/labstack/echo/v4"
	res "github.com/srv-cashpay/util/s/response"
)

type MidtransNotification struct {
	OrderID           string `json:"order_id"`
	TransactionStatus string `json:"transaction_status"`
	GrossAmount       string `json:"gross_amount"`
	SignatureKey      string `json:"signature_key"`
}

func (h *domainHandler) MidtransCallback(c echo.Context) error {
	var notificationPayload MidtransNotification

	// Parsing request body dari Midtrans
	err := c.Bind(&notificationPayload)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	// Ambil OrderID dan status dari payload
	orderID := notificationPayload.OrderID
	transactionStatus := notificationPayload.TransactionStatus

	// Update status pembayaran di database berdasarkan OrderID
	err = h.servicePackages.UpdateStatus(orderID, transactionStatus)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err).Send(c)
	}

	return res.SuccessResponse("Payment status updated successfully").Send(c)
}
