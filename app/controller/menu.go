package controller

import (
	"context"

	"github.com/kyleu/projectforge/app"
	"github.com/kyleu/projectforge/app/menu"
	"github.com/kyleu/projectforge/app/sandbox"
	"github.com/kyleu/projectforge/app/telemetry"
	"github.com/kyleu/projectforge/app/util"
)

func MenuFor(ctx context.Context, isAuthed bool, isAdmin bool, as *app.State) (menu.Items, error) {
	ctx, span := telemetry.StartSpan(ctx, "menu", "menu:generate")
	defer span.End()

	var ret menu.Items
	// $PF_SECTION_START(routes_start)$
	itemFor := func(t action.Type, i string, r string) *menu.Item {
		return &menu.Item{Key: t.Key, Title: t.Title, Description: t.Description, Icon: i, Route: r}
	}

	if isAuthed {
		ret = append(ret,
			projectMenu(as.Services.Projects.Projects()),
			menu.Separator,
			moduleMenu(as.Services.Modules.Modules()),
			menu.Separator,
		)
	}
	// $PF_SECTION_END(routes_start)$
	// $PF_INJECT_START(codegen)$
	// $PF_INJECT_END(codegen)$
	// $PF_SECTION_START(routes_end)$
	if isAdmin {
		ret = append(ret,
			sandbox.Menu(),
			menu.Separator,
			&menu.Item{Key: "admin", Title: "Settings", Description: "System-wide settings and preferences", Icon: "cog", Route: "/admin"},
			itemFor(action.TypeDoctor, "first-aid", "/doctor"),
		)
	}
	desc := "Get assistance and advice for using " + util.AppName
	ret = append(ret, &menu.Item{Key: "about", Title: "About", Description: desc, Icon: "question", Route: "/about"})
	// $PF_SECTION_END(routes_end)$
	return ret, nil
}
