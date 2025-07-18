package dto

import (
	"mime/multipart"
	"time"
)

type ProductRequest struct {
	ID           string `json:"id"`
	UserID       string `json:"user_id"`
	Barcode      string `json:"barcode"`
	MerchantID   string `json:"merchant_id"`
	MerkID       string `json:"merk_id"`
	CategoryID   string `json:"category_id"`
	ProductName  string `json:"product_name"`
	Description  string `json:"description"`
	Stock        int    `json:"stock"`
	MinimalStock int    `json:"minimal_stock"`
	Price        int    `json:"price"`
	Status       int    `json:"status"`
	CreatedBy    string `json:"created_by"`
}

type ProductResponse struct {
	ID           string    `json:"id"`
	UserID       string    `json:"user_id"`
	MerchantID   string    `json:"merchant_id"`
	Barcode      string    `json:"barcode"`
	MerkID       string    `json:"merk_id"`
	CategoryID   string    `json:"category_id"`
	ProductName  string    `json:"product_name"`
	Description  string    `json:"description"`
	Stock        int       `json:"stock"`
	MinimalStock int       `json:"minimal_stock"`
	Price        int       `json:"price"`
	Status       string    `json:"status"`
	CreatedBy    string    `json:"created_by"`
	CreatedAt    Timestamp `json:"created_at"`
}

type ProductGetByIdRequest struct {
	ID string `param:"id" validate:"required"`
}

type ProductDeleteRequest struct {
	ID        string `param:"id" validate:"required"`
	DeletedBy string `json:"deleted_by"`
}

type ProductDeleteResponse struct {
	ID        string `json:"id"`
	DeletedBy string `json:"deleted_by"`
}

type ProductBulkDeleteRequest struct {
	ID        []string `json:"id"`
	DeletedBy string   `json:"deleted_by"`
}

type ProductBulkDeleteResponse struct {
	ID        []string `json:"id"`
	DeletedBy string   `json:"deleted_by"`
	Count     int      `json:"count"`
}

type ProductUpdateRequest struct {
	ID           string `json:"id"`
	Barcode      string `json:"barcode"`
	ProductName  string `json:"product_name"`
	Stock        int    `json:"stock"`
	MinimalStock int    `json:"minimal_stock"`
	Price        int    `json:"price"`
	Status       int    `json:"status"`
	UpdatedBy    string `json:"updated_by"`
	UserID       string `json:"user_id"`
	MerchantID   string `json:"merchant_id"`
	Description  string `json:"description"`
}

type ProductUploadRequest struct {
	ID          string `json:"id"`
	File        *multipart.FileHeader
	CreatedBy   string `json:"created_by"`
	UpdatedBy   string `json:"updated_by"`
	ProductID   string `json:"product_id"`
	UserID      string `json:"user_id"`
	MerchantID  string `json:"merchant_id"`
	Destination string `json:"destination"`
}

type ProductUploadResponse struct {
	FilePath string `json:"file_path"`
}

type GetProductUploadRequest struct {
	FileName string `param:"file_name" validate:"required"`
}

type GetProductUploadResponse struct {
	FileName string `json:"file_name"`
	FilePath string `json:"file_path"`
}

type ProductUpdateResponse struct {
	ID           string `json:"id"`
	Barcode      string `json:"barcode"`
	ProductName  string `json:"product_name"`
	Stock        int    `json:"stock"`
	MinimalStock int    `json:"minimal_stock"`
	Price        int    `json:"price"`
	Status       int    `json:"status"`
	UpdatedBy    string `json:"updated_by"`
	UserID       string `json:"user_id"`
	MerchantID   string `json:"merchant_id"`
	Description  string `json:"description"`
}

// Formatter untuk timestamp dengan nama bulan
type Timestamp time.Time

// Format waktu: 19 December 2024, 09:28:17
const timeFormat = "02 January 2006, 15:04:05"

// MarshalJSON untuk memformat waktu
func (t Timestamp) MarshalJSON() ([]byte, error) {
	// Konversi waktu ke zona waktu lokal
	localTime := time.Time(t).Local()
	formattedTime := "\"" + localTime.Format(timeFormat) + "\""
	return []byte(formattedTime), nil
}

// UnmarshalJSON untuk parsing waktu dari JSON (opsional)
func (t *Timestamp) UnmarshalJSON(data []byte) error {
	parsedTime, err := time.Parse("\""+timeFormat+"\"", string(data))
	if err != nil {
		return err
	}
	*t = Timestamp(parsedTime)
	return nil
}

// ToTime untuk mengonversi kembali ke time.Time
func (t Timestamp) ToTime() time.Time {
	return time.Time(t)
}
