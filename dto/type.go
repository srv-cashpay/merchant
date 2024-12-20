package dto

type TypeRequest struct {
	ID         string `json:"id"`
	UserID     string `json:"user_id"`
	MerchantID string `json:"merchant_id"`
	TypeName   string `json:"type_name"`
	CreatedBy  string `json:"created_by"`
}
