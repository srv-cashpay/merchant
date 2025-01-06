package packages

import (
	"fmt"
	"os"

	"github.com/midtrans/midtrans-go"

	"github.com/midtrans/midtrans-go/snap"

	dto "github.com/srv-cashpay/merchant/dto"
	util "github.com/srv-cashpay/util/s"
)

func (s *packagesService) Create(req dto.PackagesRequest) (dto.PackagesResponse, error) {
	// Buat record awal dengan status default sementara
	create := dto.PackagesRequest{
		ID:          util.GenerateRandomString(),
		UserID:      req.UserID,
		CreatedBy:   req.CreatedBy,
		OrderID:     util.GenerateRandomString(), // Generate Order ID
		GrossAmount: req.GrossAmount,
		Status:      "initiated", // Status awal sebelum transaksi ke Midtrans
	}

	// Simpan data awal package ke repository
	created, err := s.Repo.Create(create)
	if err != nil {
		return dto.PackagesResponse{}, err
	}

	// Inisialisasi client Midtrans
	midtransClient := snap.Client{}
	midtransClient.New(os.Getenv("MidKeyProd"), midtrans.Production)

	// Setup data transaksi Midtrans
	transactionReq := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  create.OrderID,
			GrossAmt: int64(create.GrossAmount),
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: req.FirstName,
			LName: req.LastName,
			Email: req.Email,
			Phone: req.CreatedBy,
		},
	}

	// Buat Snap token dan ambil respons dari Midtrans
	snapTokenResp, _ := midtransClient.CreateTransaction(transactionReq)

	// Update status berdasarkan respons Midtrans
	status := "pending" // Status default berdasarkan skenario pembayaran Midtrans
	if snapTokenResp.Token == "" {
		status = "failed" // Jika token tidak berhasil dibuat
	}

	// Simpan status transaksi terbaru ke database
	err = s.Repo.UpdateStatus(create.OrderID, status)
	if err != nil {
		return dto.PackagesResponse{}, err
	}

	// Persiapkan response
	response := dto.PackagesResponse{
		ID:          created.ID,
		OrderID:     create.OrderID,
		GrossAmount: create.GrossAmount,
		Token:       snapTokenResp.Token,
		UserID:      created.UserID,
		CreatedBy:   created.CreatedBy,
		Status:      status,
		RedirectURL: snapTokenResp.RedirectURL,
	}

	return response, nil
}

func (s *packagesService) UpdateStatus(orderID string, transactionStatus string) error {
	var status string

	// Mapping status dari Midtrans ke status internal
	switch transactionStatus {
	case "capture", "settlement":
		status = "paid"
	case "deny", "cancel", "expire":
		status = "failed"
	case "pending":
		status = "pending"
	default:
		return fmt.Errorf("unknown transaction status: %s", transactionStatus)
	}

	// Update status di repository
	err := s.Repo.UpdateStatus(orderID, status)
	if err != nil {
		return err
	}

	return nil
}
