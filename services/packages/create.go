package packages

import (
	"fmt"
	"os"

	"github.com/midtrans/midtrans-go"

	"github.com/midtrans/midtrans-go/snap"
	"github.com/srv-cashpay/merchant/constant"

	dto "github.com/srv-cashpay/merchant/dto"
	util "github.com/srv-cashpay/util/s"
)

func (s *packagesService) Create(req dto.PackagesRequest) (dto.PackagesResponse, error) {
	orderID := util.GenerateRandomString()

	create := dto.PackagesRequest{
		ID:          util.GenerateRandomString(),
		UserID:      req.UserID,
		CreatedBy:   req.CreatedBy,
		OrderID:     orderID,
		GrossAmount: req.GrossAmount,
		Status:      constant.StatusInitiated,
	}

	created, err := s.Repo.Create(create)
	if err != nil {
		return dto.PackagesResponse{}, err
	}

	midtransClient := snap.Client{}
	midtransClient.New(os.Getenv("MidKeyProd"), midtrans.Production)

	transactionReq := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  create.OrderID,
			GrossAmt: int64(create.GrossAmount),
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: create.CreatedBy,
			LName: req.LastName,
			Email: req.Email,
			Phone: req.CreatedBy,
		},
	}

	snapTokenResp, err := midtransClient.CreateTransaction(transactionReq)
	if snapTokenResp == nil || snapTokenResp.Token == "" {
		_ = s.Repo.UpdateStatus(orderID, constant.StatusFailed)
		return dto.PackagesResponse{}, nil
	}

	_ = s.Repo.UpdateStatus(orderID, constant.StatusPending)

	response := dto.PackagesResponse{
		ID:          created.ID,
		OrderID:     orderID,
		GrossAmount: create.GrossAmount,
		Token:       snapTokenResp.Token,
		UserID:      created.UserID,
		CreatedBy:   created.CreatedBy,
		Status:      constant.StatusPending,
		RedirectURL: snapTokenResp.RedirectURL,
	}

	return response, nil
}

func (s *packagesService) UpdateStatus(orderID string, transactionStatus string) error {
	var status string

	switch transactionStatus {
	case "settlement", "capture":
		status = "paid"
	case "pending":
		status = "pending"
	case "deny", "cancel", "expire":
		status = "failed"
	case "refund":
		status = "refunded"
	case "partial_refund":
		status = "partial_refunded"
	default:
		return fmt.Errorf("unknown transaction status: %s", transactionStatus)
	}

	return s.Repo.UpdateStatus(orderID, status)
}
