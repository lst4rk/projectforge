package controller

import (
	"context"

	"{{{ .Package }}}/app"
	"{{{ .Package }}}/app/menu"
	"{{{ .Package }}}/app/sandbox"
	"{{{ .Package }}}/app/telemetry"
	"{{{ .Package }}}/app/util"
)

func MenuFor(ctx context.Context, isAuthed bool, isAdmin bool, as *app.State) (menu.Items, error) {
	ctx, span := telemetry.StartSpan(ctx, "menu", "menu:generate")
	defer span.End()

	var ret menu.Items
	// $PF_SECTION_START(routes_start)$
	ret = append(ret,
		&menu.Item{Key: "quickstart", Title: "Quickstart", Description: "Check out your fancy app!", Icon: "star", Route: "/quickstart"},
		menu.Separator,
	)
	// $PF_SECTION_END(routes_start)$
	// $PF_INJECT_START(codegen)$
	// $PF_INJECT_END(codegen)$
	// $PF_SECTION_START(routes_end)$
	if isAdmin {
		admin := &menu.Item{Key: "admin", Title: "Settings", Description: "System-wide settings and preferences", Icon: "cog", Route: "/admin"}
		ret = append(ret, sandbox.Menu(), menu.Separator, admin)
	}
	aboutDesc := "Get assistance and advice for using " + util.AppName
	ret = append(ret, &menu.Item{Key: "about", Title: "About", Description: aboutDesc, Icon: "question", Route: "/about"})
	// $PF_SECTION_END(routes_end)$
	return ret, nil
}
