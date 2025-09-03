package order

import (
	"github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (r *orderRepository) SaveToken(req dto.TokenRequest) error {
	return r.DB.Exec(`
        INSERT INTO fcm_tokens (user_id, token)
        VALUES (?, ?)
        ON CONFLICT (user_id) DO UPDATE SET token = EXCLUDED.token
    `, req.UserID, req.Token).Error
}

func (r *orderRepository) GetAllTokens() ([]string, error) {
	var tokens []string
	err := r.DB.Raw("SELECT token FROM fcm_tokens").Scan(&tokens).Error
	if err != nil {
		return nil, err
	}
	return tokens, nil
}

func (r *orderRepository) DeleteToken(token string) error {
	return r.DB.Where("token = ?", token).Delete(&entity.FCMToken{}).Error
}

func (r *orderRepository) SaveOrder(req dto.FCMRequest) (dto.FCMResponse, error) {

	create := entity.Order{
		Title:      req.Title,
		Body:       req.Body,
		ID:         req.ID,
		OrderName:  req.OrderName,
		MerchantID: req.MerchantID,
		UserID:     req.UserID,
		CreatedBy:  req.CreatedBy,
		Product:    req.ProductJSON,
	}

	if err := r.DB.Save(&create).Error; err != nil {
		return dto.FCMResponse{}, err
	}

	response := dto.FCMResponse{
		Name: req.Name,
	}

	return response, nil

}
