package dto

type MerkRequest struct {
	ID         string `json:"id"`
	UserID     string `json:"user_id"`
	MerchantID string `json:"merchant_id"`
	MerkName   string `json:"merk_name"`
	CreatedBy  string `json:"created_by"`
}

type MerkResponse struct {
	ID         string `json:"id"`
	UserID     string `json:"user_id"`
	MerchantID string `json:"merchant_id"`
	MerkName   string `json:"merk_name"`
	CreatedBy  string `json:"created_by"`
}

type GetByIdRequest struct {
	ID string `param:"id" validate:"required"`
}

type DeleteRequest struct {
	ID        string `param:"id" validate:"required"`
	DeletedBy string `json:"deleted_by"`
}

type DeleteResponse struct {
	ID        string `param:"id" validate:"required"`
	DeletedBy string `json:"deleted_by"`
}

type MerkUpdateRequest struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	MerkName  string `json:"merk_name"`
	UpdatedBy string `json:"updated_by"`
}

type MerkUpdateResponse struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	MerkName  string `json:"merk_name"`
	UpdatedBy string `json:"updated_by"`
}

type BulkDeleteRequest struct {
	ID        []string `json:"id"`
	DeletedBy string   `json:"deleted_by"`
}

type BulkDeleteResponse struct {
	ID        []string `json:"id"`
	DeletedBy string   `json:"deleted_by"`
	Count     int      `json:"count"`
}
