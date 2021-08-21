// Code generated by Project Forge, see https://projectforge.dev for details.
package controller

import (
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"

	"github.com/kyleu/projectforge/app"
	"github.com/kyleu/projectforge/app/controller/cutil"
	"github.com/kyleu/projectforge/app/sandbox"
	"github.com/kyleu/projectforge/views/vsandbox"
)

func SandboxList(ctx *fasthttp.RequestCtx) {
	act("sandbox.list", ctx, func(as *app.State, ps *cutil.PageState) (string, error) {
		ps.Title = "Sandboxes"
		ps.Data = sandbox.AllSandboxes
		return render(ctx, as, &vsandbox.List{}, ps, "sandbox")
	})
}

func SandboxRun(ctx *fasthttp.RequestCtx) {
	act("sandbox.run", ctx, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := ctxRequiredString(ctx, "key", false)
		if err != nil {
			return "", err
		}
		sb := sandbox.AllSandboxes.Get(key)
		if sb == nil {
			return ersp("no sandbox with key [%s]", key)
		}
		ret, err := sb.Run(ps.Context, as, ps.Logger.With(zap.String("sandbox", key)))
		if err != nil {
			return "", err
		}
		ps.Title = sb.Title
		ps.Data = ret
		if sb.Key == "testbed" {
			return render(ctx, as, &vsandbox.Testbed{}, ps, "sandbox", sb.Key)
		}
		return render(ctx, as, &vsandbox.Run{Key: key, Title: sb.Title, Result: ret}, ps, "sandbox||/sandbox", sb.Key)
	})
}
