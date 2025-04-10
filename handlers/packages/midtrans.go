package packages

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/srv-cashpay/merchant/constant"
	res "github.com/srv-cashpay/util/s/response"
)

type MidtransNotification struct {
	OrderID           string `json:"order_id"`
	TransactionStatus string `json:"transaction_status"`
	StatusCode        string `json:"status_code"` // <- tambahkan ini
	GrossAmount       string `json:"gross_amount"`
	SignatureKey      string `json:"signature_key"`
	PaymentType       string `json:"payment_type"`
	FraudStatus       string `json:"fraud_status"` // penting untuk kartu kredit
}

func (h *domainHandler) MidtransCallback(c echo.Context) error {
	var notificationPayload MidtransNotification

	if err := c.Bind(&notificationPayload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	orderID := notificationPayload.OrderID
	midtransStatus := notificationPayload.TransactionStatus

	internalStatus, err := MapMidtransStatusToInternal(midtransStatus)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	err = h.servicePackages.UpdateStatus(orderID, internalStatus)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err).Send(c)
	}

	return res.SuccessResponse("Payment status updated successfully").Send(c)
}

func MapMidtransStatusToInternal(midtransStatus string) (string, error) {
	switch midtransStatus {
	case "capture", "settlement":
		return constant.StatusPaid, nil
	case "deny", "cancel", "expire":
		return constant.StatusFailed, nil
	case "pending":
		return constant.StatusPending, nil
	default:
		return "", fmt.Errorf("unknown midtrans status: %s", midtransStatus)
	}
}
