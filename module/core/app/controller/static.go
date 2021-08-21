package controller

import (
	"path/filepath"
	"strings"

	"github.com/valyala/fasthttp"

	"{{{ .Package }}}/assets"
)

func Favicon(ctx *fasthttp.RequestCtx) {
	data, contentType, err := assets.EmbedAsset("favicon.ico")
	assetResponse(ctx, data, contentType, err)
}

func RobotsTxt(ctx *fasthttp.RequestCtx) {
	data, contentType, err := assets.EmbedAsset("robots.txt")
	assetResponse(ctx, data, contentType, err)
}

func Static(ctx *fasthttp.RequestCtx) {
	path, err := filepath.Abs(strings.TrimPrefix(string(ctx.Request.URI().Path()), "/assets"))
	if err == nil {
		path = strings.TrimPrefix(path, "/")
		data, contentType, e := assets.EmbedAsset(path)
		assetResponse(ctx, data, contentType, e)
	} else {
		ctx.Error(err.Error(), fasthttp.StatusBadRequest)
	}
}

func assetResponse(ctx *fasthttp.RequestCtx, data []byte, contentType string, err error) {
	if err == nil {
		ctx.Response.Header.SetContentType(contentType)
		ctx.SetStatusCode(fasthttp.StatusOK)
		_, _ = ctx.Write(data)
	} else {
		ctx.Error(err.Error(), fasthttp.StatusNotFound)
	}
}
