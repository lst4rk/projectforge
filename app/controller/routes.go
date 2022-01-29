// Content managed by Project Forge, see [projectforge.md] for details.
package controller

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/projectforge/app/lib/telemetry/httpmetrics"
	"github.com/kyleu/projectforge/app/util"
)

//nolint
func AppRoutes() fasthttp.RequestHandler {
	r := router.New()

	r.GET("/", Home)
	r.GET("/healthcheck", Healthcheck)
	r.GET("/about", About)
	r.GET("/theme", ThemeList)
	r.GET("/theme/{key}", ThemeEdit)
	r.POST("/theme/{key}", ThemeSave)
	r.GET(defaultSearchPath, Search)

	r.GET(defaultProfilePath, Profile)
	r.POST(defaultProfilePath, ProfileSave)
	r.GET("/auth/{key}", AuthDetail)
	r.GET("/auth/callback/{key}", AuthCallback)
	r.GET("/auth/logout/{key}", AuthLogout)

	// $PF_SECTION_START(routes)$
	r.GET("/welcome", Welcome)
	r.POST("/welcome", WelcomeResult)

	r.GET("/doctor", Doctor)
	r.GET("/doctor/all", DoctorRunAll)
	r.GET("/doctor/{key}", DoctorRun)

	r.GET("/p", ProjectList)
	r.GET("/p/new", ProjectForm)
	r.POST("/p/new", ProjectCreate)
	r.GET("/p/{key}", ProjectDetail)
	r.GET("/p/{key}/edit", ProjectEdit)
	r.POST("/p/{key}/edit", ProjectSave)

	r.GET("/p/{key}/fs", ProjectFileRoot)
	r.GET("/p/{key}/fs/{path:*}", ProjectFile)

	r.GET("/svg/{key}", SVGList)
	r.GET("/svg/{key}/x/add", SVGAdd)
	r.GET("/svg/{key}/x/build", SVGBuild)
	r.GET("/svg/{key}/{icon}", SVGDetail)
	r.GET("/svg/{key}/{icon}/setapp", SVGSetApp)
	r.GET("/svg/{key}/{icon}/remove", SVGRemove)

	r.GET("/b/{key}", Build)
	r.GET("/b/{key}/{act}", Build)

	r.GET("/git", GitActionAll)
	r.GET("/git/all/{act}", GitActionAll)

	r.GET("/git/{key}", GitAction)
	r.GET("/git/{key}/{act}", GitAction)

	r.GET("/run/{act}", RunAllActions)
	r.GET("/run/{key}/{act}", RunAction)

	r.GET("/m", ModuleList)
	r.GET("/m/{key}", ModuleDetail)
	r.GET("/m/{key}/fs", ModuleFileRoot)
	r.GET("/m/{key}/fs/{path:*}", ModuleFile)

	r.GET("/test", TestList)
	r.GET("/test/{key}", TestRun)
	// $PF_SECTION_END(routes)$

	r.GET("/admin", Admin)
	r.GET("/admin/sandbox", SandboxList)
	r.GET("/admin/sandbox/{key}", SandboxRun)
	r.GET("/admin/{path:*}", Admin)

	r.GET("/favicon.ico", Favicon)
	r.GET("/robots.txt", RobotsTxt)
	r.GET("/assets/{_:*}", Static)

	r.OPTIONS("/", Options)
	r.OPTIONS("/{_:*}", Options)
	r.NotFound = NotFound

	p := httpmetrics.NewMetrics(util.AppKey)
	return fasthttp.CompressHandlerBrotliLevel(p.WrapHandler(r), fasthttp.CompressBrotliBestSpeed, fasthttp.CompressBestSpeed)
}
