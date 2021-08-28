package controller

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kyleu/projectforge/app/project"
	"github.com/kyleu/projectforge/app/theme"
	"github.com/kyleu/projectforge/app/util"
	"github.com/kyleu/projectforge/views/vproject"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/projectforge/app/controller/cutil"

	"github.com/kyleu/projectforge/app"
)

func ProjectList(rc *fasthttp.RequestCtx) {
	act("project.root", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		prjs := as.Services.Projects.Projects()
		ps.Title = "Project Listing"
		ps.Data = prjs
		return render(rc, as, &vproject.List{Projects: prjs}, ps, "projects")
	})
}

func ProjectDetail(rc *fasthttp.RequestCtx) {
	act("project.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		prj, err := getProject(rc, as)
		if err != nil {
			return "", err
		}

		ps.Title = fmt.Sprintf("%s (project %s)", prj.Title(), prj.Key)
		ps.Data = prj
		return render(rc, as, &vproject.Detail{Project: prj}, ps, "projects", prj.Key)
	})
}

func ProjectEdit(rc *fasthttp.RequestCtx) {
	act("project.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		prj, err := getProject(rc, as)
		if err != nil {
			return "", err
		}

		ps.Title = fmt.Sprintf("%s (project %s)", prj.Title(), prj.Key)
		ps.Data = prj
		return render(rc, as, &vproject.Edit{Project: prj}, ps, "projects", prj.Key)
	})
}

func ProjectSave(rc *fasthttp.RequestCtx) {
	act("project.save", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		prj, err := getProject(rc, as)
		if err != nil {
			return "", err
		}

		frm, err := cutil.ParseForm(rc)
		if err != nil {
			return "", err
		}
		get := func(k string, def string) string {
			x, _ := frm.GetString(k, true)
			if x == "" {
				return def
			}
			return x
		}
		prj.Name = get("name", prj.Name)
		prj.Version = get("version", prj.Version)
		prj.Package = get("package", prj.Package)
		prj.Args = get("args", prj.Args)
		prj.Port, _ = strconv.Atoi(get("port", fmt.Sprintf("%d", prj.Port)))
		prj.Modules = util.SplitAndTrim(get("modules", strings.Join(prj.Modules, "|")), "|")
		prj.Ignore = util.SplitAndTrim(get("ignore", strings.Join(prj.Ignore, ",")), ",")
		prj.Children = util.SplitAndTrim(get("children", strings.Join(prj.Children, "\n")), "\n")

		prj.Info.Org = get("org", prj.Info.Org)
		prj.Info.AuthorName = get("authorName", prj.Info.AuthorName)
		prj.Info.AuthorEmail = get("authorEmail", prj.Info.AuthorEmail)
		prj.Info.License = get("license", prj.Info.License)
		prj.Info.Bundle = get("bundle", prj.Info.Bundle)
		prj.Info.SigningIdentity = get("signingIdentity", prj.Info.SigningIdentity)
		prj.Info.Homepage = get("homepage", prj.Info.Homepage)
		prj.Info.Sourcecode = get("sourcecode", prj.Info.Sourcecode)
		prj.Info.Summary = get("summary", prj.Info.Summary)
		prj.Info.Description = get("description", prj.Info.Description)

		prj.Build = project.BuildFromMap(frm)
		if prj.Build.Empty() {
			prj.Build = nil
		}
		prj.Theme = theme.ApplyMap(frm)
		if prj.Theme.Equals(theme.ThemeDefault) {
			prj.Theme = nil
		}

		err = as.Services.Projects.Save(prj)
		if err != nil {
			return ersp("unable to save project: %+v", err)
		}

		msg := "Saved changes"
		return flashAndRedir(true, msg, "/p/"+prj.Key, rc, ps)
	})
}

func getProject(rc *fasthttp.RequestCtx, as *app.State) (*project.Project, error) {
	key, err := rcRequiredString(rc, "key", true)
	if err != nil {
		return nil, err
	}

	prj, err := as.Services.Projects.Get(key)
	if err != nil {
		return nil, err
	}
	if prj.Info == nil {
		prj.Info = &project.Info{}
	}
	return prj, nil
}
