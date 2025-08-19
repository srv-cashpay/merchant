package dashboard

type FCMToken struct {
	UserID string `gorm:"primaryKey"`
	Token  string
}

func (r *dashboardRepository) SaveToken(userID, token string) error {
	err := r.DB.Exec(`
		INSERT INTO fcm_tokens (user_id, token)
		VALUES (?, ?)
		ON CONFLICT (user_id) DO UPDATE SET token = EXCLUDED.token
	`, userID, token).Error

	return err
}

func (r *dashboardRepository) GetAllTokens() ([]string, error) {
	var tokens []string

	// Query pakai Raw + Scan
	err := r.DB.Raw("SELECT token FROM fcm_tokens").Scan(&tokens).Error
	if err != nil {
		return nil, err
	}

	return tokens, nil
}
