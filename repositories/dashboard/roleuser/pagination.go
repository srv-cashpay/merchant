package roleuser

import (
	"fmt"
	"math"
	"strings"

	dto "github.com/srv-cashpay/merchant/dto"
)

func (r *RoleUserRepository) Pagination(req *dto.Pagination) (RepositoryResult, int) {
	var roleusers []dto.RoleUserResponse
	var totalRows int64

	offset := (req.Page - 1) * req.Limit

	find := r.DB.Table("role_users AS ru").
		Select(`
			ru.id,
			roles.role AS role_id,
			access_doors.full_name AS user_id,
			permissions.label AS permission_id,   
			ru.created_at
		`).
		Joins("JOIN roles ON roles.id = ru.role_id").
		Joins("JOIN permissions ON permissions.id = ru.permission_id"). // <=== JOIN BY ID
		Joins("JOIN access_doors ON access_doors.id = ru.user_id").
		Limit(req.Limit).
		Offset(offset).
		Order(req.Sort)

	// Search filter
	if req.Searchs != nil {
		for _, s := range req.Searchs {
			switch s.Action {
			case "equals":
				find = find.Where(fmt.Sprintf("%s = ?", s.Column), s.Query)
			case "contains":
				find = find.Where(fmt.Sprintf("%s LIKE ?", s.Column), "%"+s.Query+"%")
			case "in":
				find = find.Where(fmt.Sprintf("%s IN (?)", s.Column), strings.Split(s.Query, ","))
			}
		}
	}

	if err := find.Scan(&roleusers).Error; err != nil {
		return RepositoryResult{Error: err}, 0
	}

	req.Rows = roleusers

	// Count
	if err := r.DB.Table("role_users AS ru").
		Joins("JOIN roles ON roles.id = ru.role_id").
		Joins("JOIN permissions ON permissions.id = ru.permission_id").
		Joins("JOIN access_doors ON access_doors.id = ru.user_id").
		Count(&totalRows).Error; err != nil {
		return RepositoryResult{Error: err}, 0
	}

	req.TotalRows = int(totalRows)
	totalPages := int(math.Ceil(float64(totalRows) / float64(req.Limit)))
	req.TotalPages = totalPages

	// fromRow / toRow
	if req.Page == 1 {
		req.FromRow = 1
		req.ToRow = len(roleusers)
	} else {
		req.FromRow = (req.Page-1)*req.Limit + 1
		req.ToRow = req.FromRow + len(roleusers) - 1
	}

	return RepositoryResult{Result: req}, totalPages
}
