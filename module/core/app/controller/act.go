package controller

import (
	"fmt"
	"time"

	"github.com/valyala/fasthttp"
	"go.uber.org/zap"

	"{{{ .Package }}}/app"
	"{{{ .Package }}}/app/controller/cutil"
	{{{ if.HasModule "marketing" }}}"{{{ .Package }}}/app/site"
	{{{ end }}}"{{{ .Package }}}/app/theme"
	"{{{ .Package }}}/app/util"
	"{{{ .Package }}}/app/web"
	"{{{ .Package }}}/views/verror"
)

const (
	defaultSearchPath  = "/search"
	defaultProfilePath = "/profile"
	defaultIcon        = "app"
)

func act(key string, ctx *fasthttp.RequestCtx, f func(as *app.State, ps *cutil.PageState) (string, error)) {
	as := _currentAppState
	ps := loadPageState(ctx, as.Logger)
	err := initAppRequest(as, ps)
	if err != nil {
		as.Logger.Warnf("%+v", err)
	}
	err = clean(as, ps)
	if err != nil {
		as.Logger.Warnf("%+v", err)
	}
	actComplete(key, as, ps, ctx, f)
}
{{{ if.HasModule "marketing" }}}
func actSite(key string, ctx *fasthttp.RequestCtx, f func(as *app.State, ps *cutil.PageState) (string, error)) {
	as := _currentSiteState
	ps := loadPageState(ctx, as.Logger)
	ps.Menu = site.Menu(ps.Profile, ps.Accounts)
	err := initSiteRequest(as, ps)
	if err != nil {
		as.Logger.Warnf("%+v", err)
	}
	err = clean(as, ps)
	if err != nil {
		as.Logger.Warnf("%+v", err)
	}
	actComplete(key, as, ps, ctx, f)
}
{{{ end }}}
func loadPageState(ctx *fasthttp.RequestCtx, logger *zap.SugaredLogger) *cutil.PageState {
	logger = logger.With(zap.String("path", string(ctx.Request.URI().Path())))

	if store == nil {
		store = initStore([]byte(sessionKey))
	}
	session, err := store.Get(ctx, util.AppKey)
	if err != nil {
		logger.Warnf("error retrieving session: %+v", err)
		session, err = store.New(ctx, util.AppKey)
		if err != nil {
			logger.Warnf("error creating new session: %+v", err)
		}
	}
	flashes := util.StringArrayFromInterfaces(session.Flashes())
	if len(flashes) > 0 {
		err = web.SaveSession(ctx, session, logger)
		if err != nil {
			logger.Warnf("can't save session: %+v", err)
		}
	}

	prof, err := loadProfile(session)
	if err != nil {
		logger.Warnf("can't load profile: %+v", err)
	}

	var a web.Accounts
	authX, ok := session.Values["auth"]
	if ok {
		authS, ok := authX.(string)
		if ok {
			a = web.AccountsFromString(authS)
		}
	}

	return &cutil.PageState{
		Method:   string(ctx.Method()),
		URI:      ctx.Request.URI(),
		Flashes:  flashes,
		Session:  session,
		Profile:  prof,
		Accounts: a,
		Icons:    initialIcons,
		Logger:   logger,
	}
}

func actComplete(key string, as *app.State, ps *cutil.PageState, ctx *fasthttp.RequestCtx, f func(as *app.State, ps *cutil.PageState) (string, error)) {
	status := fasthttp.StatusOK
	cutil.WriteCORS(ctx)
	startNanos := time.Now().UnixNano()
	redir, err := f(as, ps)
	if err != nil {
		status = fasthttp.StatusInternalServerError
		ctx.SetStatusCode(status)

		ps.Logger.Errorf("error running action [%s]: %+v", key, err)

		if len(ps.Breadcrumbs) == 0 {
			ps.Breadcrumbs = cutil.Breadcrumbs{"Error"}
		}
		errDetail := util.GetErrorDetail(err)
		page := &verror.Error{Err: errDetail}

		err := clean(as, ps)
		if err != nil {
			as.Logger.Error(err)
			msg := fmt.Sprintf("error while cleaning request: %+v", err)
			as.Logger.Error(msg)
			_, _ = ctx.Write([]byte(msg))
		}
		redir, err = render(ctx, as, page, ps)
		if err != nil {
			msg := fmt.Sprintf("error while running error handler: %+v", err)
			as.Logger.Error(msg)
			_, _ = ctx.Write([]byte(msg))
		}
	}
	if redir != "" {
		ctx.Response.Header.Set("Location", redir)
		status = fasthttp.StatusFound
		ctx.SetStatusCode(status)
	}
	elapsedMillis := float64((time.Now().UnixNano()-startNanos)/int64(time.Microsecond)) / float64(1000)
	l := ps.Logger.With(zap.String("method", ps.Method), zap.Int("status", status), zap.Float64("elapsed", elapsedMillis))
	l.Debugf("processed request in [%.3fms] (render: %.3fms)", elapsedMillis, ps.RenderElapsed)
}

func clean(as *app.State, ps *cutil.PageState) error {
	if ps.Profile != nil && ps.Profile.Theme == "" {
		ps.Profile.Theme = theme.ThemeDefault.Key
	}
	if ps.RootIcon == "" {
		ps.RootIcon = defaultIcon
	}
	if ps.RootPath == "" {
		ps.RootPath = "/"
	}
	if ps.RootTitle == "" {
		ps.RootTitle = util.AppName
	}
	if ps.SearchPath == "" {
		ps.SearchPath = defaultSearchPath
	}
	if ps.ProfilePath == "" {
		ps.ProfilePath = defaultProfilePath
	}
	if len(ps.Menu) == 0 {
		m, err := MenuFor(as)
		if err != nil {
			return err
		}
		ps.Menu = m
	}
	return nil
}
