package voucher

import (
	dto "github.com/srv-cashpay/merchant/dto"
)

func (b *voucherService) GetById(req dto.GetByIdRequest) (*dto.VoucherResponse, error) {
	transaction, err := b.Repo.GetById(req)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}
