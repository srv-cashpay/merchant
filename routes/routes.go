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

	h_packages "github.com/srv-cashpay/merchant/handlers/packages"
	r_packages "github.com/srv-cashpay/merchant/repositories/packages"
	s_packages "github.com/srv-cashpay/merchant/services/packages"

	h_authenticator "github.com/srv-cashpay/merchant/handlers/authenticator_request"
	r_authenticator "github.com/srv-cashpay/merchant/repositories/authenticator_request"
	s_authenticator "github.com/srv-cashpay/merchant/services/authenticator_request"

	h_product "github.com/srv-cashpay/merchant/handlers/product"
	r_product "github.com/srv-cashpay/merchant/repositories/product"
	s_product "github.com/srv-cashpay/merchant/services/product"

	h_getmerk "github.com/srv-cashpay/merchant/handlers/product/merk"
	r_getmerk "github.com/srv-cashpay/merchant/repositories/product/merk"
	s_getmerk "github.com/srv-cashpay/merchant/services/product/merk"

	h_getcategory "github.com/srv-cashpay/merchant/handlers/product/category"
	r_getcategory "github.com/srv-cashpay/merchant/repositories/product/category"
	s_getcategory "github.com/srv-cashpay/merchant/services/product/category"

	"github.com/srv-cashpay/middlewares/middlewares"
)

var (
	DB = configs.InitDB()

	JWT = middlewares.NewJWTService()

	authenticatorR = r_authenticator.NewAuthenticatorRepository(DB)
	authenticatorS = s_authenticator.NewAuthenticatorService(authenticatorR, JWT)
	authenticatorH = h_authenticator.NewAuthenticatorHandler(authenticatorS)

	packagesR = r_packages.NewPackagesRepository(DB)
	packagesS = s_packages.NewPackagesService(packagesR, JWT)
	packagesH = h_packages.NewPackagesHandler(packagesS)

	posR = r_pos.NewPosRepository(DB)
	posS = s_pos.NewPosService(posR, JWT)
	posH = h_pos.NewPosHandler(posS)

	merkR = r_merk.NewMerkRepository(DB)
	merkS = s_merk.NewMerkService(merkR, JWT)
	merkH = h_merk.NewMerkHandler(merkS)

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

	getcategoryR = r_getcategory.NewGetCategoryRepository(DB)
	getcategoryS = s_getcategory.NewGetCategoryService(getcategoryR, JWT)
	getcategoryH = h_getcategory.NewCategoryHandler(getcategoryS)
)

func New() *echo.Echo {

	e := echo.New()

	pos := e.Group("api/merchant", middlewares.AuthorizeJWT(JWT))
	{
		pos.POST("/pos/create", posH.Create)
	}
	packages := e.Group("api/merchant", middlewares.AuthorizeJWT(JWT))
	{
		packages.POST("/packages/create", packagesH.Create)
		packages.POST("/midtrans/callback", packagesH.MidtransCallback)

	}
	merk := e.Group("api/merchant", middlewares.AuthorizeJWT(JWT))
	{
		merk.POST("/merk/create", merkH.Create)
		merk.GET("/merk/pagination", merkH.Get)
		merk.PUT("/merk/:id", merkH.Update)
		merk.DELETE("/merk/:id", merkH.Delete)
		merk.DELETE("/merk/bulk-delete", merkH.BulkDelete)
	}
	category := e.Group("api/merchant", middlewares.AuthorizeJWT(JWT))
	{
		category.POST("/category/create", categoryH.Create)
		category.GET("/category/pagination", categoryH.Get)
		category.PUT("/category/:id", categoryH.Update)
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
	}

	return e
}
