package user

import (
	"fmt"
	"math"
	"strings"

	"github.com/srv-cashpay/auth/entity"
	dto "github.com/srv-cashpay/merchant/dto"
	mentity "github.com/srv-cashpay/merchant/entity"
	"github.com/srv-cashpay/merchant/helpers"
	util "github.com/srv-cashpay/util/s"
)

func (r *userRepository) Get(req *dto.Pagination) (dto.UserMerchantPaginationResponse, int) {
	var users []entity.AccessDoor
	var totalRows int64
	totalPages, fromRow, toRow := 0, 0, 0

	offset := (req.Page - 1) * req.Limit

	// --- Ambil data utama user merchant ---
	find := r.DB.
		Preload("Verified").
		Preload("Merchant").
		Where("merchant_id = ?", req.MerchantID)

	// --- Tambahkan filter search jika ada ---
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

	// --- Hitung total baris data ---
	if errCount := find.Model(&entity.AccessDoor{}).Count(&totalRows).Error; errCount != nil {
		return dto.UserMerchantPaginationResponse{}, 0
	}

	// --- Ambil data user merchant ---
	if err := find.
		Limit(req.Limit).
		Offset(offset).
		Order(req.Sort).
		Find(&users).Error; err != nil {
		return dto.UserMerchantPaginationResponse{}, 0
	}

	// --- Ambil semua role, buat map biar cepat lookup ---
	var roles []mentity.Role
	roleMap := make(map[string]string)
	if err := r.DB.Find(&roles).Error; err == nil {
		for _, role := range roles {
			roleMap[role.ID] = role.Role
		}
	}

	// --- Hitung total halaman ---
	totalPages = int(math.Ceil(float64(totalRows) / float64(req.Limit)))

	// --- Hitung posisi baris ---
	if req.Page == 1 {
		fromRow = 1
		toRow = req.Limit
	} else {
		fromRow = (req.Page-1)*req.Limit + 1
		toRow = req.Page * req.Limit
	}
	if toRow > int(totalRows) {
		toRow = int(totalRows)
	}

	// --- Mapping hasil data ---
	var userResponses []dto.GetUserMerchantResponse
	for _, u := range users {
		decryptedWa, err := util.Decrypt(u.Whatsapp)
		if err != nil {
			continue
		}
		decryptedEmail, err := util.Decrypt(u.Email)
		if err != nil {
			continue
		}

		verifiedStatus := "not verified"
		if u.Verified.Verified {
			verifiedStatus = "verified"
		}

		accountStatus := "inactive"
		if u.Verified.StatusAccount {
			accountStatus = "active"
		}

		userResponses = append(userResponses, dto.GetUserMerchantResponse{
			ID:       u.ID,
			FullName: helpers.TruncateString(u.FullName, 47),
			Whatsapp: decryptedWa,
			Email:    decryptedEmail,
			RoleName: roleMap[u.AccessRoleID], // ✅ ambil nama role dari map hasil query roles
			Verified: dto.UserMerchantVerified{
				ID:             u.Verified.ID,
				UserID:         u.Verified.UserID,
				Token:          u.Verified.Token,
				Verified:       verifiedStatus,
				StatusAccount:  accountStatus,
				AccountExpired: u.Verified.AccountExpired,
				Otp:            u.Verified.Otp,
				ExpiredAt:      u.Verified.ExpiredAt,
			},
			Merchant: dto.GetMerchantResponse{
				MerchantName: u.Merchant.MerchantName,
				Address:      u.Merchant.Address,
				Country:      u.Merchant.Country,
				City:         u.Merchant.City,
				Zip:          u.Merchant.Zip,
				CurrencyID:   u.Merchant.CurrencyID,
				Phone:        u.Merchant.Phone,
			},
		})
	}

	// --- Response pagination ---
	response := dto.UserMerchantPaginationResponse{
		Limit:      req.Limit,
		Page:       req.Page,
		Sort:       req.Sort,
		TotalRows:  int(totalRows),
		TotalPages: totalPages,
		FromRow:    fromRow,
		ToRow:      toRow,
		Data:       userResponses,
		Searchs:    req.Searchs,
	}

	return response, totalPages
}
