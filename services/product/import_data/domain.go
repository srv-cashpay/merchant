package import_data

import (
	"context"
	"mime/multipart"

	m "github.com/srv-cashpay/middlewares/middlewares"

	r "github.com/srv-cashpay/merchant/repositories/product/import_data"
)

type ImportService interface {
	GenerateTemplate() ([]byte, string, error)
	ImportProducts(ctx context.Context, fileHeader *multipart.FileHeader) (map[string]interface{}, error)
}

type importService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewImportService(Repo r.DomainRepository, jwtS m.JWTService) ImportService {
	return &importService{
		Repo: Repo,
		jwt:  jwtS,
	}
}
