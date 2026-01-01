package qris

import (
	"fmt"
	"strconv"

	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
	util "github.com/srv-cashpay/util/s"
)

func (r *qrisRepository) Create(req dto.CoQrisRequest) (dto.CoQrisResponse, error) {

	// Create the new qris entry
	create := entity.Qris{
		ID:         util.GenerateRandomString(),
		QrisName:   req.QrisName,
		Link:       req.Link,
		Status:     req.Status,
		UserID:     req.UserID,
		MerchantID: req.MerchantID,
		CreatedBy:  req.CreatedBy,
	}

	// Save the new qris to the database
	if err := r.DB.Save(&create).Error; err != nil {
		return dto.CoQrisResponse{}, err
	}

	// Map the status from integer to string
	statusMap := map[int]string{
		1: "active",
		2: "inactive",
	}

	createdStatus, err := strconv.Atoi(fmt.Sprintf("%v", create.Status))
	if err != nil {
		return dto.CoQrisResponse{}, fmt.Errorf("invalid status value: %v", create.Status)
	}

	statusString, ok := statusMap[createdStatus]
	if !ok {
		return dto.CoQrisResponse{}, fmt.Errorf("invalid status value in database")
	}

	// Build the response for the created qris
	response := dto.CoQrisResponse{
		ID:         create.ID,
		QrisName:   create.QrisName,
		Link:       create.Link,
		UserID:     create.UserID,
		MerchantID: create.MerchantID,
		Status:     statusString,
		CreatedBy:  create.CreatedBy,
	}

	return response, nil
}

// Function to generate the qris ID
