package controller

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

//nolint
func AppRoutes() *router.Router {
	w := fasthttp.CompressHandler
	r := router.New()

	r.GET("/", w(Home))
	r.GET("/settings", w(Settings))
	r.GET("/theme", w(ThemeList))
	r.GET("/theme/{key}", w(ThemeEdit))
	r.POST("/theme/{key}", w(ThemeSave))
	r.GET(defaultSearchPath, w(Search))

	r.GET(defaultProfilePath, w(Profile))
	r.POST(defaultProfilePath, w(ProfileSave))
	r.GET("/auth/{key}", w(AuthDetail))
	r.GET("/auth/{key}/callback", w(AuthCallback))
	r.GET("/auth/{key}/logout", w(AuthLogout))

	r.GET("/help", w(Help))
	r.GET("/help/{path:*}", w(Help))

	r.GET("/modules", w(Modules))

	r.GET("/favicon.ico", Favicon)
	r.GET("/robots.txt", RobotsTxt)
	r.GET("/assets/{_:*}", Static)

	r.OPTIONS("/", w(Options))
	r.OPTIONS("/{_:*}", w(Options))
	r.NotFound = NotFound

	return r
}
