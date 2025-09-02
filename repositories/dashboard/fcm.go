package dashboard

import (
	"github.com/srv-cashpay/merchant/entity"
)

func (r *dashboardRepository) SaveToken(userID, token string) error {
	return r.DB.Exec(`
        INSERT INTO fcm_tokens (user_id, token)
        VALUES (?, ?)
        ON CONFLICT (user_id) DO UPDATE SET token = EXCLUDED.token
    `, userID, token).Error
}

func (r *dashboardRepository) GetAllTokens() ([]string, error) {
	var tokens []string
	err := r.DB.Raw("SELECT token FROM fcm_tokens").Scan(&tokens).Error
	if err != nil {
		return nil, err
	}
	return tokens, nil
}

func (r *dashboardRepository) DeleteToken(token string) error {
	return r.DB.Where("token = ?", token).Delete(&entity.FCMToken{}).Error
}
