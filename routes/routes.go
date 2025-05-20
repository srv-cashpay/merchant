package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/srv-cashpay/merchant/configs"
	h_pos "github.com/srv-cashpay/merchant/handlers/pos"
	r_pos "github.com/srv-cashpay/merchant/repositories/pos"
	s_pos "github.com/srv-cashpay/merchant/services/pos"

	h_merk "github.com/srv-cashpay/merchant/handlers/merk"
	r_merk "github.com/srv-cashpay/merchant/repositories/merk"
	s_merk "github.com/srv-cashpay/merchant/services/merk"

	h_category "github.com/srv-cashpay/merchant/handlers/category"
	r_category "github.com/srv-cashpay/merchant/repositories/category"
	s_category "github.com/srv-cashpay/merchant/services/category"

	h_unit "github.com/srv-cashpay/merchant/handlers/unit"
	r_unit "github.com/srv-cashpay/merchant/repositories/unit"
	s_unit "github.com/srv-cashpay/merchant/services/unit"

	h_dashboard "github.com/srv-cashpay/merchant/handlers/dashboard"
	r_dashboard "github.com/srv-cashpay/merchant/repositories/dashboard"
	s_dashboard "github.com/srv-cashpay/merchant/services/dashboard"

	h_permission "github.com/srv-cashpay/merchant/handlers/dashboard/permission"
	r_permission "github.com/srv-cashpay/merchant/repositories/dashboard/permission"
	s_permission "github.com/srv-cashpay/merchant/services/dashboard/permission"

	// r_role "github.com/srv-cashpay/merchant/repositories/dashboard/role"
	// s_role "github.com/srv-cashpay/merchant/services/dashboard/role"

	h_subscribe "github.com/srv-cashpay/merchant/handlers/subscribe"
	r_subscribe "github.com/srv-cashpay/merchant/repositories/subscribe"
	s_subscribe "github.com/srv-cashpay/merchant/services/subscribe"

	h_authenticator "github.com/srv-cashpay/merchant/handlers/authenticator_request"
	r_authenticator "github.com/srv-cashpay/merchant/repositories/authenticator_request"
	s_authenticator "github.com/srv-cashpay/merchant/services/authenticator_request"

	h_tax "github.com/srv-cashpay/merchant/handlers/tax"
	r_tax "github.com/srv-cashpay/merchant/repositories/tax"
	s_tax "github.com/srv-cashpay/merchant/services/tax"

	h_discount "github.com/srv-cashpay/merchant/handlers/discount"
	r_discount "github.com/srv-cashpay/merchant/repositories/discount"
	s_discount "github.com/srv-cashpay/merchant/services/discount"

	h_payment "github.com/srv-cashpay/merchant/handlers/payment"
	r_payment "github.com/srv-cashpay/merchant/repositories/payment"
	s_payment "github.com/srv-cashpay/merchant/services/payment"

	h_history "github.com/srv-cashpay/merchant/handlers/subscribe/history"
	r_history "github.com/srv-cashpay/merchant/repositories/subscribe/history"
	s_history "github.com/srv-cashpay/merchant/services/subscribe/history"

	h_user "github.com/srv-cashpay/merchant/handlers/user"
	r_user "github.com/srv-cashpay/merchant/repositories/user"
	s_user "github.com/srv-cashpay/merchant/services/user"

	h_product "github.com/srv-cashpay/merchant/handlers/product"
	r_product "github.com/srv-cashpay/merchant/repositories/product"
	s_product "github.com/srv-cashpay/merchant/services/product"

	h_getmerk "github.com/srv-cashpay/merchant/handlers/product/merk"
	r_getmerk "github.com/srv-cashpay/merchant/repositories/product/merk"
	s_getmerk "github.com/srv-cashpay/merchant/services/product/merk"

	h_getcategory "github.com/srv-cashpay/merchant/handlers/product/category"
	r_getcategory "github.com/srv-cashpay/merchant/repositories/product/category"
	s_getcategory "github.com/srv-cashpay/merchant/services/product/category"

	h_printer "github.com/srv-cashpay/merchant/handlers/printer"
	r_printer "github.com/srv-cashpay/merchant/repositories/printer"
	s_printer "github.com/srv-cashpay/merchant/services/printer"

	h_merchant "github.com/srv-cashpay/merchant/handlers/merchant"
	r_merchant "github.com/srv-cashpay/merchant/repositories/merchant"
	s_merchant "github.com/srv-cashpay/merchant/services/merchant"

	h_table "github.com/srv-cashpay/merchant/handlers/table"
	r_table "github.com/srv-cashpay/merchant/repositories/table"
	s_table "github.com/srv-cashpay/merchant/services/table"

	"github.com/srv-cashpay/middlewares/middlewares"
)

var (
	DB = configs.InitDB()

	JWT = middlewares.NewJWTService()

	authenticatorR = r_authenticator.NewAuthenticatorRepository(DB)
	authenticatorS = s_authenticator.NewAuthenticatorService(authenticatorR, JWT)
	authenticatorH = h_authenticator.NewAuthenticatorHandler(authenticatorS)

	merchantR = r_merchant.NewMerchantRepository(DB)
	merchantS = s_merchant.NewMerchantService(merchantR, JWT)
	merchantH = h_merchant.NewMerchantHandler(merchantS)

	printerR = r_printer.NewPrinterRepository(DB)
	printerS = s_printer.NewPrinterService(printerR, JWT)
	printerH = h_printer.NewPrinterHandler(printerS)

	subscribeR = r_subscribe.NewSubscribeRepository(DB)
	subscribeS = s_subscribe.NewSubscribeService(subscribeR, JWT)
	subscribeH = h_subscribe.NewSubscribeHandler(subscribeS)

	posR = r_pos.NewPosRepository(DB)
	posS = s_pos.NewPosService(posR, JWT)
	posH = h_pos.NewPosHandler(posS)

	merkR = r_merk.NewMerkRepository(DB)
	merkS = s_merk.NewMerkService(merkR, JWT)
	merkH = h_merk.NewMerkHandler(merkS)

	permissionR = r_permission.NewPermissionRepository(DB)
	permissionS = s_permission.NewPermissionService(permissionR, JWT)
	permissionH = h_permission.NewPermissionHandler(permissionS)

	categoryR = r_category.NewCategoryRepository(DB)
	categoryS = s_category.NewCategoryService(categoryR, JWT)
	categoryH = h_category.NewCategoryHandler(categoryS)

	unitR = r_unit.NewUnitRepository(DB)
	unitS = s_unit.NewUnitService(unitR, JWT)
	unitH = h_unit.NewUnitHandler(unitS)

	dashboardR = r_dashboard.NewDashboardRepository(DB)
	dashboardS = s_dashboard.NewDashboardService(dashboardR, JWT)
	dashboardH = h_dashboard.NewDashboardHandler(dashboardS)

	productR = r_product.NewProductRepository(DB)
	productS = s_product.NewProductService(productR, JWT)
	productH = h_product.NewProductHandler(productS)

	getmerkR = r_getmerk.NewGetMerkRepository(DB)
	getmerkS = s_getmerk.NewGetMerkService(getmerkR, JWT)
	getmerkH = h_getmerk.NewMerkHandler(getmerkS)

	discountR = r_discount.NewDiscountRepository(DB)
	discountS = s_discount.NewDiscountService(discountR, JWT)
	discountH = h_discount.NewDiscountHandler(discountS)

	paymentR = r_payment.NewPaymentRepository(DB)
	paymentS = s_payment.NewPaymentService(paymentR, JWT)
	paymentH = h_payment.NewPaymentHandler(paymentS)

	historyR = r_history.NewHistoryRepository(DB)
	historyS = s_history.NewHistoryService(historyR, JWT)
	historyH = h_history.NewHistoryHandler(historyS)

	tableR = r_table.NewTableRepository(DB)
	tableS = s_table.NewTableService(tableR, JWT)
	tableH = h_table.NewTableHandler(tableS)

	userR = r_user.NewUserRepository(DB)
	userS = s_user.NewUserService(userR, JWT)
	userH = h_user.NewUserHandler(userS)

	taxR = r_tax.NewTaxRepository(DB)
	taxS = s_tax.NewTaxService(taxR, JWT)
	taxH = h_tax.NewTaxHandler(taxS)

	getcategoryR = r_getcategory.NewGetCategoryRepository(DB)
	getcategoryS = s_getcategory.NewGetCategoryService(getcategoryR, JWT)
	getcategoryH = h_getcategory.NewCategoryHandler(getcategoryS)
)

func New() *echo.Echo {

	e := echo.New()
	sub := e.Group("sub", middlewares.AuthorizeJWT(JWT))
	{
		sub.GET("/transaction/:order_id/status", subscribeH.CheckTransactionStatus)
		sub.POST("/midtrans/callback", subscribeH.MidtransCallback)
		sub.POST("/charge-bni", subscribeH.ChargeBni)
		sub.POST("/charge-permata", subscribeH.ChargePermata)
		sub.POST("/charge-mandiri", subscribeH.ChargeMandiri)
		sub.POST("/charge-bri", subscribeH.ChargeBri)
		sub.POST("/charge-cimb", subscribeH.ChargeCimb)

		sub.POST("/charge-qris", subscribeH.ChargeQris)
		sub.POST("/charge-gopay", subscribeH.ChargeGopay)
		sub.POST("/charge-shopeepay", subscribeH.ChargeShopeePay)

		sub.POST("/charge-gpay", subscribeH.ChargeGpay)
		sub.GET("/tokenize", subscribeH.TokenizeCardHandler)
		sub.POST("/charge-card", subscribeH.CardPayment)
		sub.POST("/cancel/:order_id", subscribeH.CancelPay)
	}
	pos := e.Group("api/merchant", middlewares.AuthorizeJWT(JWT))
	{
		pos.POST("/pos/create", posH.Create)
	}
	packages := e.Group("api/merchant", middlewares.AuthorizeJWT(JWT))
	{
		packages.POST("/packages/create", subscribeH.Create)
	}
	merchant := e.Group("api/merchant", middlewares.AuthorizeJWT(JWT))
	{
		merchant.PUT("/update", merchantH.Update)
		merchant.GET("/get", merchantH.Get)
	}
	printer := e.Group("api/merchant", middlewares.AuthorizeJWT(JWT))
	{
		printer.PUT("/printer/update", printerH.Update)
		printer.GET("/printer/get", printerH.Get)
		printer.POST("/printer/create", printerH.Create)
		printer.DELETE("/printer/:id", printerH.Delete)
	}

	merk := e.Group("api/merchant", middlewares.AuthorizeJWT(JWT))
	{
		merk.POST("/merk/create", merkH.Create)
		merk.GET("/merk/pagination", merkH.Get)
		merk.GET("/merk/:id", merkH.GetById)
		merk.PUT("/merk/update/:id", merkH.Update)
		merk.DELETE("/merk/:id", merkH.Delete)
		merk.DELETE("/merk/bulk-delete", merkH.BulkDelete)
	}
	permission := e.Group("api/merchant", middlewares.AuthorizeJWT(JWT))
	{
		permission.POST("/permission/create", permissionH.Create)
		permission.GET("/permission", permissionH.Get)
		permission.GET("/permission/pagination", permissionH.Pagination)
		permission.PUT("/permission/update/:id", permissionH.Update)
		permission.DELETE("/permission/:id", permissionH.Delete)
		permission.DELETE("/permission/bulk-delete", permissionH.BulkDelete)
	}
	// role := e.Group("api/merchant", middlewares.AuthorizeJWT(JWT))
	{
		// role.POST("/role/create", roleH.Create)
		// role.GET("/role/pagination", roleH.Get)
		// role.PUT("/role/update/:id", roleH.Update)
		// role.DELETE("/role/:id", roleH.Delete)
		// role.DELETE("/role/bulk-delete", roleH.BulkDelete)
	}
	tax := e.Group("api/merchant", middlewares.AuthorizeJWT(JWT))
	{
		tax.POST("/tax/create", taxH.Create)
		tax.GET("/tax/pagination", taxH.Get)
		tax.PUT("/tax/update/:id", taxH.Update)
		tax.DELETE("/tax/:id", taxH.Delete)
		tax.DELETE("/tax/bulk-delete", taxH.BulkDelete)
	}

	table := e.Group("api/merchant", middlewares.AuthorizeJWT(JWT))
	{
		table.POST("/table/create", tableH.Create)
		table.GET("/table/pagination", tableH.Get)
		table.PUT("/table/update/:id", tableH.Update)
		table.DELETE("/table/:id", tableH.Delete)
		table.DELETE("/table/bulk-delete", tableH.BulkDelete)
	}

	payment := e.Group("api/merchant", middlewares.AuthorizeJWT(JWT))
	{
		payment.POST("/payment/create", paymentH.Create)
		payment.GET("/payment/pagination", paymentH.Get)
		payment.PUT("/payment/update/:id", paymentH.Update)
		payment.DELETE("/payment/:id", paymentH.Delete)
		payment.DELETE("/payment/bulk-delete", paymentH.BulkDelete)
	}
	history := e.Group("api/merchant", middlewares.AuthorizeJWT(JWT))
	{
		history.GET("/history/pagination", historyH.Get)
		history.GET("/history/:id", historyH.GetById)
	}

	discount := e.Group("api/merchant", middlewares.AuthorizeJWT(JWT))
	{
		discount.POST("/discount/create", discountH.Create)
		discount.GET("/discount/pagination", discountH.Get)
		discount.PUT("/discount/update/:id", discountH.Update)
		discount.DELETE("/discount/:id", discountH.Delete)
		discount.DELETE("/discount/bulk-delete", discountH.BulkDelete)
	}

	user := e.Group("api/merchant", middlewares.AuthorizeJWT(JWT))
	{
		user.POST("/user/create", userH.Create)
		user.GET("/user/pagination", userH.Get)
		user.PUT("/user/update/:id", userH.Update)
		user.DELETE("/user/:id", userH.Delete)
		user.DELETE("/user/bulk-delete", userH.BulkDelete)
	}

	category := e.Group("api/merchant", middlewares.AuthorizeJWT(JWT))
	{
		category.POST("/category/create", categoryH.Create)
		category.GET("/category/pagination", categoryH.Get)
		category.GET("/category/:id", categoryH.GetById)
		category.PUT("/category/update/:id", categoryH.Update)
		category.DELETE("/category/:id", categoryH.Delete)
		category.DELETE("/category/bulk-delete", categoryH.BulkDelete)
	}
	unit := e.Group("api/merchant", middlewares.AuthorizeJWT(JWT))
	{
		unit.POST("/unit/create", unitH.Create)
		unit.GET("/unit/pagination", unitH.Get)
		unit.PUT("/unit/:id", unitH.Update)
		unit.DELETE("/unit/:id", unitH.Delete)
		unit.DELETE("/unit/bulk-delete", unitH.BulkDelete)
	}
	authenticator := e.Group("api/merchant", middlewares.AuthorizeJWT(JWT))
	{
		authenticator.POST("/authenticator/request", authenticatorH.Create)
	}
	dashboard := e.Group("api/dashboard", middlewares.AuthorizeJWT(JWT))
	{
		dashboard.GET("/index", dashboardH.Get)
	}
	product := e.Group("/api/merchant", middlewares.AuthorizeJWT(JWT))
	{
		product.POST("/product/create", productH.Create)
		product.GET("/product/:id", productH.GetById)
		product.DELETE("/product/:id", productH.Delete)
		product.DELETE("/product/bulk-delete", productH.BulkDelete)
		product.PUT("/product/update/:id", productH.Update)
		product.GET("/product/pagination", productH.Get)
		product.GET("/product/merk", getmerkH.Get)
		product.GET("/product/category", getcategoryH.Get)
		product.PUT("/product/upload/:id", productH.UploadImage)
	}
	e.GET("/uploads/:file_name", productH.GetPicture)

	return e
}
