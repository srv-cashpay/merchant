package reservation

import (
	"strconv"

	"github.com/labstack/echo/v4"
	dto "github.com/srv-cashpay/merchant/dto"
	res "github.com/srv-cashpay/util/s/response"
)

func (b *domainHandler) GetById(c echo.Context) error {
	var req dto.GetReservationByIdRequest

	// idStr := c.Param("id")

	idUint, err := res.IsNumber(c, "id")
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	idInt, err := strconv.Atoi(idUint)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	req.ID = uint(idInt)

	transaction, err := b.serviceReservation.GetById(req)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.NotFound, err).Send(c)

	}

	return res.SuccessResponse(transaction).Send(c)

}
