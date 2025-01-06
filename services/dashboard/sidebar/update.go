package sidebar

import "github.com/srv-cashpay/merchant/dto"

func (b *sidebarService) Update(req dto.SidebarUpdateRequest) (dto.SidebarUpdateResponse, error) {
	request := dto.SidebarUpdateRequest{
		Label:     req.Label,
		Icon:      req.Icon,
		UpdatedBy: req.UpdatedBy,
		UserID:    req.UserID,
		To:        req.To,
	}

	sidebar, err := b.Repo.Update(req)
	if err != nil {
		return sidebar, err
	}

	response := dto.SidebarUpdateResponse{
		Label:     req.Label,
		Icon:      req.Icon,
		UpdatedBy: request.UpdatedBy,
		UserID:    request.UserID,
		To:        req.To,
	}

	return response, nil
}
