package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/srv-cashpay/merchant/configs"
	h_pos "github.com/srv-cashpay/merchant/handlers/pos"
	r_pos "github.com/srv-cashpay/merchant/repositories/pos"
	s_pos "github.com/srv-cashpay/merchant/services/pos"

	h_order "github.com/srv-cashpay/merchant/handlers/order"
	r_order "github.com/srv-cashpay/merchant/repositories/order"
	s_order "github.com/srv-cashpay/merchant/services/order"

	h_merk "github.com/srv-cashpay/merchant/handlers/merk"
	r_merk "github.com/srv-cashpay/merchant/repositories/merk"
	s_merk "github.com/srv-cashpay/merchant/services/merk"

	h_voucher "github.com/srv-cashpay/merchant/handlers/voucher"
	r_voucher "github.com/srv-cashpay/merchant/repositories/voucher"
	s_voucher "github.com/srv-cashpay/merchant/services/voucher"

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

	h_role "github.com/srv-cashpay/merchant/handlers/dashboard/role"
	r_role "github.com/srv-cashpay/merchant/repositories/dashboard/role"
	s_role "github.com/srv-cashpay/merchant/services/dashboard/role"

	h_role_user "github.com/srv-cashpay/merchant/handlers/dashboard/roleuser"
	r_role_user "github.com/srv-cashpay/merchant/repositories/dashboard/roleuser"
	s_role_user "github.com/srv-cashpay/merchant/services/dashboard/roleuser"

	h_role_user_permission "github.com/srv-cashpay/merchant/handlers/dashboard/roleuserpermission"
	r_role_user_permission "github.com/srv-cashpay/merchant/repositories/dashboard/roleuserpermission"
	s_role_user_permission "github.com/srv-cashpay/merchant/services/dashboard/roleuserpermission"

	h_subscribe "github.com/srv-cashpay/merchant/handlers/subscribe"
	r_subscribe "github.com/srv-cashpay/merchant/repositories/subscribe"
	s_subscribe "github.com/srv-cashpay/merchant/services/subscribe"

	h_topup "github.com/srv-cashpay/merchant/handlers/topup"
	r_topup "github.com/srv-cashpay/merchant/repositories/topup"
	s_topup "github.com/srv-cashpay/merchant/services/topup"

	h_authenticator "github.com/srv-cashpay/merchant/handlers/authenticator_request"
	r_authenticator "github.com/srv-cashpay/merchant/repositories/authenticator_request"
	s_authenticator "github.com/srv-cashpay/merchant/services/authenticator_request"

	h_tax "github.com/srv-cashpay/merchant/handlers/tax"
	r_tax "github.com/srv-cashpay/merchant/repositories/tax"
	s_tax "github.com/srv-cashpay/merchant/services/tax"

	h_discount "github.com/srv-cashpay/merchant/handlers/discount"
	r_discount "github.com/srv-cashpay/merchant/repositories/discount"
	s_discount "github.com/srv-cashpay/merchant/services/discount"

	h_paymentmethod "github.com/srv-cashpay/merchant/handlers/subscribe/paymentmethod"
	r_paymentmethod "github.com/srv-cashpay/merchant/repositories/subscribe/paymentmethod"
	s_paymentmethod "github.com/srv-cashpay/merchant/services/subscribe/paymentmethod"

	h_transactionmethode "github.com/srv-cashpay/merchant/handlers/transactionmethode/qris"
	r_transactionmethode "github.com/srv-cashpay/merchant/repositories/transactionmethode/qris"
	s_transactionmethode "github.com/srv-cashpay/merchant/services/transactionmethode/qris"

	h_history "github.com/srv-cashpay/merchant/handlers/subscribe/history"
	r_history "github.com/srv-cashpay/merchant/repositories/subscribe/history"
	s_history "github.com/srv-cashpay/merchant/services/subscribe/history"

	h_user "github.com/srv-cashpay/merchant/handlers/user"
	r_user "github.com/srv-cashpay/merchant/repositories/user"
	s_user "github.com/srv-cashpay/merchant/services/user"

	h_usermerchant "github.com/srv-cashpay/merchant/handlers/usermerchant"
	r_usermerchant "github.com/srv-cashpay/merchant/repositories/usermerchant"
	s_usermerchant "github.com/srv-cashpay/merchant/services/usermerchant"

	h_product "github.com/srv-cashpay/merchant/handlers/product"
	r_product "github.com/srv-cashpay/merchant/repositories/product"
	s_product "github.com/srv-cashpay/merchant/services/product"

	h_importproduct "github.com/srv-cashpay/merchant/handlers/product/import_data"
	r_importproduct "github.com/srv-cashpay/merchant/repositories/product/import_data"
	s_importproduct "github.com/srv-cashpay/merchant/services/product/import_data"

	h_exportproduct "github.com/srv-cashpay/merchant/handlers/product/export_data"
	r_exportproduct "github.com/srv-cashpay/merchant/repositories/product/export_data"
	s_exportproduct "github.com/srv-cashpay/merchant/services/product/export_data"

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

	h_contentsetting "github.com/srv-cashpay/merchant/handlers/dashboard/contentsetting"
	r_contentsetting "github.com/srv-cashpay/merchant/repositories/dashboard/contentsetting"
	s_contentsetting "github.com/srv-cashpay/merchant/services/dashboard/contentsetting"

	h_table "github.com/srv-cashpay/merchant/handlers/table"
	r_table "github.com/srv-cashpay/merchant/repositories/table"
	s_table "github.com/srv-cashpay/merchant/services/table"

	h_deleteaccount "github.com/srv-cashpay/merchant/handlers/deleteaccount"
	r_deleteaccount "github.com/srv-cashpay/merchant/repositories/deleteaccount"
	s_deleteaccount "github.com/srv-cashpay/merchant/services/deleteaccount"

	h_reservation "github.com/srv-cashpay/merchant/handlers/reservation"
	r_reservation "github.com/srv-cashpay/merchant/repositories/reservation"
	s_reservation "github.com/srv-cashpay/merchant/services/reservation"

	h_pin "github.com/srv-cashpay/merchant/handlers/pin"
	r_pin "github.com/srv-cashpay/merchant/repositories/pin"
	s_pin "github.com/srv-cashpay/merchant/services/pin"

	"github.com/srv-cashpay/middlewares/middlewares"
)

var (
	DB = configs.InitDB()

	pp = configs.InitApp()

	JWT = middlewares.NewJWTService()

	authenticatorR = r_authenticator.NewAuthenticatorRepository(DB)
	authenticatorS = s_authenticator.NewAuthenticatorService(authenticatorR, JWT)
	authenticatorH = h_authenticator.NewAuthenticatorHandler(authenticatorS)

	merchantR = r_merchant.NewMerchantRepository(DB)
	merchantS = s_merchant.NewMerchantService(merchantR, JWT)
	merchantH = h_merchant.NewMerchantHandler(merchantS)

	contentsettingR = r_contentsetting.NewContentSettingRepository(DB)
	contentsettingS = s_contentsetting.NewContentSettingService(contentsettingR, JWT)
	contentsettingH = h_contentsetting.NewContentSettingHandler(contentsettingS)

	printerR = r_printer.NewPrinterRepository(DB)
	printerS = s_printer.NewPrinterService(printerR, JWT)
	printerH = h_printer.NewPrinterHandler(printerS)

	subscribeR = r_subscribe.NewSubscribeRepository(DB, pp)
	subscribeS = s_subscribe.NewSubscribeService(subscribeR, JWT)
	subscribeH = h_subscribe.NewSubscribeHandler(subscribeS)

	topupR = r_topup.NewSubscribeRepository(DB, pp)
	topupS = s_topup.NewSubscribeService(topupR, JWT)
	topupH = h_topup.NewSubscribeHandler(topupS)

	posR = r_pos.NewPosRepository(DB)
	posS = s_pos.NewPosService(posR, JWT)
	posH = h_pos.NewPosHandler(posS)

	merkR = r_merk.NewMerkRepository(DB)
	merkS = s_merk.NewMerkService(merkR, JWT)
	merkH = h_merk.NewMerkHandler(merkS)

	voucherR = r_voucher.NewVoucherRepository(DB)
	voucherS = s_voucher.NewVoucherService(voucherR, JWT)
	voucherH = h_voucher.NewVoucherHandler(voucherS)

	orderR = r_order.NewOrderRepository(DB)
	orderS = s_order.NewOrderService(orderR, JWT)
	orderH = h_order.NewOrderHandler(orderS)

	permissionR = r_permission.NewPermissionRepository(DB)
	permissionS = s_permission.NewPermissionService(permissionR, JWT)
	permissionH = h_permission.NewPermissionHandler(permissionS)

	roleR = r_role.NewRoleRepository(DB)
	roleS = s_role.NewRoleService(roleR, JWT)
	roleH = h_role.NewRoleHandler(roleS)

	roleuserR = r_role_user.NewRoleUserRepository(DB)
	roleuserS = s_role_user.NewRoleUserService(roleuserR, JWT)
	roleuserH = h_role_user.NewRoleUserHandler(roleuserS)

	roleuserpermissionR = r_role_user_permission.NewRoleUserPermissionRepository(DB)
	roleuserpermissionS = s_role_user_permission.NewRoleUserPermissionService(roleuserpermissionR, JWT)
	roleuserpermissionH = h_role_user_permission.NewRoleUserPermissionHandler(roleuserpermissionS)

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

	importproductR = r_importproduct.NewImportRepository(DB)
	importproductS = s_importproduct.NewImportService(importproductR, JWT)
	importproductH = h_importproduct.NewImportHandler(importproductS)

	exportproductR = r_exportproduct.NewExportRepository(DB)
	exportproductS = s_exportproduct.NewExportService(exportproductR, JWT)
	exportproductH = h_exportproduct.NewExportHandler(exportproductS)

	getmerkR = r_getmerk.NewGetMerkRepository(DB)
	getmerkS = s_getmerk.NewGetMerkService(getmerkR, JWT)
	getmerkH = h_getmerk.NewMerkHandler(getmerkS)

	discountR = r_discount.NewDiscountRepository(DB)
	discountS = s_discount.NewDiscountService(discountR, JWT)
	discountH = h_discount.NewDiscountHandler(discountS)

	paymentmethodR = r_paymentmethod.NewPaymentRepository(DB)
	paymentmethodS = s_paymentmethod.NewPaymentMethodService(paymentmethodR, JWT)
	paymentmethodH = h_paymentmethod.NewPaymentHandler(paymentmethodS)

	transactionmethodeR = r_transactionmethode.NewQrisRepository(DB)
	transactionmethodeS = s_transactionmethode.NewQrisService(transactionmethodeR, JWT)
	transactionmethodeH = h_transactionmethode.NewQrisHandler(transactionmethodeS)

	historyR = r_history.NewHistoryRepository(DB)
	historyS = s_history.NewHistoryService(historyR, JWT)
	historyH = h_history.NewHistoryHandler(historyS)

	tableR = r_table.NewTableRepository(DB)
	tableS = s_table.NewTableService(tableR, JWT)
	tableH = h_table.NewTableHandler(tableS)

	deleteaccountR = r_deleteaccount.NewDeleteAccountRepository(DB)
	deleteaccountS = s_deleteaccount.NewDeleteAccountService(deleteaccountR, JWT)
	deleteaccountH = h_deleteaccount.NewRequestDeleteHandler(deleteaccountS)

	reservationR = r_reservation.NewReservationRepository(DB)
	reservationS = s_reservation.NewReservationService(reservationR, JWT)
	reservationH = h_reservation.NewReservationHandler(reservationS)

	userR = r_user.NewUserRepository(DB)
	userS = s_user.NewUserService(userR, JWT)
	userH = h_user.NewUserHandler(userS)

	usermerchantR = r_usermerchant.NewUserMerchantRepository(DB)
	usermerchantS = s_usermerchant.NewUserMerchantService(usermerchantR, JWT)
	usermerchantH = h_usermerchant.NewUserMerchantHandler(usermerchantS)

	taxR = r_tax.NewTaxRepository(DB)
	taxS = s_tax.NewTaxService(taxR, JWT)
	taxH = h_tax.NewTaxHandler(taxS)

	getcategoryR = r_getcategory.NewGetCategoryRepository(DB)
	getcategoryS = s_getcategory.NewGetCategoryService(getcategoryR, JWT)
	getcategoryH = h_getcategory.NewCategoryHandler(getcategoryS)

	pinR = r_pin.NewPinRepository(DB)
	pinS = s_pin.NewPinService(pinR, JWT)
	pinH = h_pin.NewPinHandler(pinS)
)

func New() *echo.Echo {

	e := echo.New()
	// e.POST("/menu/order", orderH.Order)
	e.GET("/api/merchant/ws", orderH.HandleWebSocket)
	e.POST("/api/merchant/fcm/save-token", orderH.SaveToken)
	e.POST("/api/merchant/menu/order", orderH.SendBroadcast)
	e.GET("/api/merchant/voucher-verification/:id/:merchant_id", voucherH.GetVerifikasi)
	e.PUT("/api/merchant/voucher-verification/:id/:merchant_id", voucherH.Update)

	template := e.Group("/api/merchant", middlewares.AuthorizeJWT(JWT))
	{
		template.GET("/template", importproductH.DownloadTemplate)
		template.POST("/upload", importproductH.UploadProducts)
	}

	export := e.Group("/api/merchant", middlewares.AuthorizeJWT(JWT))
	{
		export.POST("/export/excel", exportproductH.ExportExcel)
	}

	sub := e.Group("api/merchant", middlewares.AuthorizeJWT(JWT))
	{
		sub.GET("/subscribe/transaction/:order_id/status", subscribeH.CheckTransactionStatus)
		sub.POST("/subscribe/midtrans/callback", subscribeH.MidtransCallback)
		sub.POST("/subscribe/charge-bni", subscribeH.ChargeBni)
		sub.POST("/subscribe/charge-permata", subscribeH.ChargePermata)
		sub.POST("/subscribe/charge-mandiri", subscribeH.ChargeMandiri)
		sub.POST("/subscribe/charge-bri", subscribeH.ChargeBri)
		sub.POST("/subscribe/charge-cimb", subscribeH.ChargeCimb)
		sub.POST("/subscribe/charge-qris", subscribeH.ChargeQris)
		sub.POST("/subscribe/charge-gopay", subscribeH.ChargeGopay)
		sub.POST("/subscribe/charge-shopeepay", subscribeH.ChargeShopeePay)
		sub.POST("/subscribe/charge-gpay", subscribeH.ChargeGpay)
		sub.GET("/subscribe/tokenize", subscribeH.TokenizeCardHandler)
		sub.POST("/subscribe/charge-card", subscribeH.CardPayment)
		sub.POST("/subscribe/cancel/:order_id", subscribeH.CancelPay)
		sub.POST("/subscribe/paypal", subscribeH.PayPal)
		sub.GET("/subscribe/paypal/capture/:order_id", subscribeH.CapturePaypalOrder)
	}

	topup := e.Group("api/merchant")
	{
		topup.GET("/topup/transaction/:order_id/status", topupH.CheckTransactionStatus)
		topup.POST("/topup/midtrans/callback", topupH.MidtransCallback)
		topup.POST("/topup/charge-bni", topupH.ChargeBni)
		topup.POST("/topup/charge-permata", topupH.ChargePermata)
		topup.POST("/topup/charge-mandiri", topupH.ChargeMandiri)
		topup.POST("/topup/charge-bri", topupH.ChargeBri)
		topup.POST("/topup/charge-cimb", topupH.ChargeCimb)
		topup.POST("/topup/charge-qris", topupH.ChargeQris)
		topup.POST("/topup/charge-gopay", topupH.ChargeGopay)
		topup.POST("/topup/charge-shopeepay", topupH.ChargeShopeePay)
		topup.POST("/topup/charge-gpay", topupH.ChargeGpay)
		topup.GET("/topup/tokenize", topupH.TokenizeCardHandler)
		topup.POST("/topup/charge-card", topupH.CardPayment)
		topup.POST("/topup/cancel/:order_id", topupH.CancelPay)
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
	contentsetting := e.Group("api/merchant", middlewares.AuthorizeJWT(JWT))
	{
		contentsetting.PUT("/contentsetting/update", contentsettingH.Update)
		contentsetting.GET("/contentsetting/get", contentsettingH.Get)
	}
	web := e.Group("api/merchant")
	{
		web.GET("/web/get/content", contentsettingH.Get)
		web.PUT("/web/update/content", contentsettingH.Update)
	}
	printer := e.Group("api/merchant", middlewares.AuthorizeJWT(JWT))
	{
		printer.PUT("/printer/update", printerH.Update)
		printer.GET("/printer/get", printerH.Get)
		printer.POST("/printer/create", printerH.Create)
		printer.DELETE("/printer/:id", printerH.Delete)
	}

	methode := e.Group("api/merchant", middlewares.AuthorizeJWT(JWT))
	{
		methode.POST("/methode-pay/qris", transactionmethodeH.Create)
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

	voucher := e.Group("api/merchant", middlewares.AuthorizeJWT(JWT))
	{
		voucher.POST("/voucher/create", voucherH.Create)
		voucher.GET("/voucher/pagination", voucherH.Get)
		voucher.GET("/voucher/:id", voucherH.GetById)
		voucher.DELETE("/voucher/:id", voucherH.Delete)
		voucher.DELETE("/voucher/bulk-delete", voucherH.BulkDelete)
	}

	order := e.Group("api/merchant", middlewares.AuthorizeJWT(JWT))
	{
		order.POST("/order/create", orderH.Create)
		order.GET("/order/pagination", orderH.Get)
		order.GET("/order/:id", orderH.GetById)
		order.PUT("/order/update/:id", orderH.Update)
		order.DELETE("/order/:id", orderH.Delete)
		order.DELETE("/order/bulk-delete", orderH.BulkDelete)
	}

	pin := e.Group("api/merchant", middlewares.AuthorizeJWT(JWT))
	{
		pin.POST("/pin/create", pinH.Create)
		pin.GET("/pin/pagination", pinH.Get)
		pin.GET("/pin/:id", pinH.GetById)
		pin.PUT("/pin/update/:id", pinH.Update)
		pin.DELETE("/pin/:id", pinH.Delete)
		pin.DELETE("/pin/bulk-delete", pinH.BulkDelete)
		pin.POST("/verify-pin", pinH.VerifyPIN)
		pin.GET("/pin/status", pinH.GetPinStatus)

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
	role := e.Group("api/merchant", middlewares.AuthorizeJWT(JWT))
	{
		role.POST("/role/create", roleH.Create)
		permission.GET("/role", permissionH.Get)
		role.GET("/role/pagination", roleH.Pagination)
		role.GET("/role_user", roleH.RoleUser)
		role.PUT("/role/update/:id", roleH.Update)
		role.DELETE("/role/:id", roleH.Delete)
		role.DELETE("/role/bulk-delete", roleH.BulkDelete)
	}
	roleuser := e.Group("api/merchant", middlewares.AuthorizeJWT(JWT))
	{
		roleuser.POST("/roleuser/create", roleuserH.Create)
		roleuser.GET("/roleuser", roleuserH.Get)
		roleuser.GET("/roleuser/pagination", roleuserH.Pagination)
		roleuser.PUT("/roleuser/update/:id", roleuserH.Update)
		roleuser.DELETE("/roleuser/:id", roleuserH.Delete)
		roleuser.DELETE("/roleuser/bulk-delete", roleuserH.BulkDelete)
	}
	roleuserpermission := e.Group("api/merchant", middlewares.AuthorizeJWT(JWT))
	{
		roleuserpermission.POST("/roleuserpermission/create", roleuserpermissionH.Create)
		roleuserpermission.GET("/roleuserpermission", roleuserpermissionH.Get)
		roleuserpermission.GET("/roleuserpermission/pagination", roleuserpermissionH.Pagination)
		roleuserpermission.PUT("/roleuserpermission/update/:id", roleuserpermissionH.Update)
		roleuserpermission.DELETE("/roleuserpermission/:id", roleuserpermissionH.Delete)
		roleuserpermission.DELETE("/roleuserpermission/bulk-delete", roleuserpermissionH.BulkDelete)
	}

	tax := e.Group("api/merchant", middlewares.AuthorizeJWT(JWT))
	{
		tax.POST("/tax/create", taxH.Create)
		tax.GET("/tax/pagination", taxH.Get)
		tax.GET("/tax/:id", taxH.GetById)
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
	reservation := e.Group("api/merchant", middlewares.AuthorizeJWT(JWT))
	{
		reservation.POST("/reservation/create", reservationH.Create)
		reservation.GET("/reservation/pagination", reservationH.Get)
		reservation.PUT("/reservation/update/:id", reservationH.Update)
		reservation.DELETE("/reservation/:id", reservationH.Delete)
		reservation.DELETE("/reservation/bulk-delete", reservationH.BulkDelete)
	}

	paymentmethod := e.Group("api/merchant", middlewares.AuthorizeJWT(JWT))
	{
		paymentmethod.POST("/payment-method/create", paymentmethodH.Create)
		paymentmethod.GET("/payment-method", paymentmethodH.Get)
		paymentmethod.PUT("/payment-method/update/:id", paymentmethodH.Update)
		paymentmethod.DELETE("/payment-method/:id", paymentmethodH.Delete)
		paymentmethod.DELETE("/payment-method/bulk-delete", paymentmethodH.BulkDelete)

	}
	e.GET("/api/merchant/uploads/:file_name", productH.GetPicture)
	e.GET("/api/merchant/payment-method/uploads/:file_name", paymentmethodH.GetPicture)

	history := e.Group("api/merchant", middlewares.AuthorizeJWT(JWT))
	{
		history.GET("/history/pagination", historyH.Get)
		history.GET("/history/:id", historyH.GetById)
		history.PUT("/history/expire/:order_id", historyH.CheckExpire)
	}

	discount := e.Group("api/merchant", middlewares.AuthorizeJWT(JWT))
	{
		discount.POST("/discount/create", discountH.Create)
		discount.GET("/discount/pagination", discountH.Get)
		discount.GET("/discount/:id", discountH.GetById)
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

	usermerchant := e.Group("api/merchant", middlewares.AuthorizeJWT(JWT))
	{
		usermerchant.POST("/user_merchant/create", usermerchantH.Create)
		usermerchant.GET("/user_merchant/pagination", usermerchantH.Get)
		usermerchant.GET("/user_merchant/:id", usermerchantH.GetById)
		usermerchant.PUT("/user_merchant/update/:id", usermerchantH.Update)
		usermerchant.DELETE("/user_merchant/:id", usermerchantH.Delete)
		usermerchant.DELETE("/user_merchant/bulk-delete", usermerchantH.BulkDelete)
	}

	category := e.Group("api/merchant", middlewares.AuthorizeJWT(JWT))
	{
		category.POST("/category/create", categoryH.Create)
		category.GET("/category/pagination", categoryH.Pagination)
		category.GET("/category/:id", categoryH.GetById)
		category.PUT("/category/update/:id", categoryH.Update)
		category.DELETE("/category/:id", categoryH.Delete)
		category.DELETE("/category/bulk-delete", categoryH.BulkDelete)
	}
	unit := e.Group("api/merchant", middlewares.AuthorizeJWT(JWT))
	{
		unit.POST("/unit/create", unitH.Create)
		unit.GET("/unit/pagination", unitH.Get)
		unit.GET("/unit/:id", unitH.GetById)
		unit.PUT("/unit/:id", unitH.Update)
		unit.DELETE("/unit/:id", unitH.Delete)
		unit.DELETE("/unit/bulk-delete", unitH.BulkDelete)
	}
	deleteAccount := e.Group("api/account", middlewares.AuthorizeJWT(JWT))
	{
		deleteAccount.POST("/request-delete", deleteaccountH.Create)
		// deleteAccount.GET("/unit/pagination", unitH.Get)
		// deleteAccount.PUT("/unit/:id", unitH.Update)
		// deleteAccount.DELETE("/unit/:id", unitH.Delete)
		// deleteAccount.DELETE("/unit/bulk-delete", unitH.BulkDelete)
	}
	authenticator := e.Group("api/merchant", middlewares.AuthorizeJWT(JWT))
	{
		authenticator.POST("/authenticator/request", authenticatorH.Create)
	}
	dashboard := e.Group("api/merchant", middlewares.AuthorizeJWT(JWT))
	{
		dashboard.GET("/dashboard/index", dashboardH.Get)
		dashboard.GET("/dashboard/route", dashboardH.Get)
	}
	product := e.Group("api/merchant", middlewares.AuthorizeJWT(JWT))
	{
		product.POST("/product/create", productH.Create)
		product.GET("/product/:id", productH.GetById)
		product.DELETE("/product/:id", productH.Delete)
		product.DELETE("/product/bulk-delete", productH.BulkDelete)
		product.PUT("/product/bulk-edit", productH.BulkEdit)
		product.PUT("/product/update/:id", productH.Update)
		product.GET("/product/pagination", productH.Get)
		product.GET("/product/merk", getmerkH.Get)
		product.GET("/product/category", getcategoryH.Get)
		product.PUT("/product/upload/:id", productH.UploadImage)
	}

	return e
}
