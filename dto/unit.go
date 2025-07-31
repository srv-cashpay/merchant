package dto

type UnitRequest struct {
	ID         string `json:"id"`
	UserID     string `json:"user_id"`
	MerchantID string `json:"merchant_id"`
	UnitName   string `json:"unit_name"`
	Status     int    `json:"status"`
	CreatedBy  string `json:"created_by"`
}

type UnitResponse struct {
	ID         string `json:"id"`
	UserID     string `json:"user_id"`
	MerchantID string `json:"merchant_id"`
	UnitName   string `json:"unit_name"`
	Status     int    `json:"status"`
	CreatedBy  string `json:"created_by"`
}

type UnitGetByIdRequest struct {
	ID string `param:"id" validate:"required"`
}

type UnitDeleteRequest struct {
	ID        string `param:"id" validate:"required"`
	DeletedBy string `json:"deleted_by"`
}

type UnitDeleteResponse struct {
	ID        string `param:"id" validate:"required"`
	DeletedBy string `json:"deleted_by"`
}

type UnitUpdateRequest struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	UnitName  string `json:"unit_name"`
	Status    int    `json:"status"`
	UpdatedBy string `json:"updated_by"`
}

type UnitUpdateResponse struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	UnitName  string `json:"unit_name"`
	Status    int    `json:"status"`
	UpdatedBy string `json:"updated_by"`
}

type UnitBulkDeleteRequest struct {
	ID        []string `json:"id"`
	DeletedBy string   `json:"deleted_by"`
}

type UnitBulkDeleteResponse struct {
	ID        []string `json:"id"`
	DeletedBy string   `json:"deleted_by"`
	Count     int      `json:"count"`
}
