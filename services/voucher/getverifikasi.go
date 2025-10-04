package voucher

import (
	dto "github.com/srv-cashpay/merchant/dto"
)

func (b *voucherService) GetVerifikasi(req dto.GetVerifikasi) (*dto.VoucherResponse, error) {
	transaction, err := b.Repo.GetVerifikasi(req)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}
