package view

import (
	"fmt"
	"strings"

	"projectforge.dev/projectforge/app/file"
	"projectforge.dev/projectforge/app/project"
	"projectforge.dev/projectforge/app/project/export/files/helper"
	"projectforge.dev/projectforge/app/project/export/golang"
	"projectforge.dev/projectforge/app/project/export/model"
)

func edit(m *model.Model, p *project.Project, args *model.Args, addHeader bool) (*file.File, error) {
	g := golang.NewGoTemplate([]string{"views", m.PackageWithGroup("v")}, "Edit.html")
	for _, imp := range helper.ImportsForTypes("webedit", m.Columns.Types()...) {
		g.AddImport(imp)
	}
	g.AddImport(helper.ImpApp, helper.ImpComponents, helper.ImpCutil, helper.ImpLayout)
	g.AddImport(helper.AppImport("app/" + m.PackageWithGroup("")))
	veb, err := exportViewEditBody(m, p, args)
	if err != nil {
		return nil, err
	}
	g.AddBlocks(exportViewEditClass(m), veb)
	return g.Render(addHeader)
}

func exportViewEditClass(m *model.Model) *golang.Block {
	ret := golang.NewBlock("Edit", "struct")
	ret.W("{%% code type Edit struct {")
	ret.W("  layout.Basic")
	ret.W("  Model *%s.%s", m.Package, m.Proper())
	ret.W("  IsNew bool")
	ret.W("} %%}")
	return ret
}

func exportViewEditBody(m *model.Model, p *project.Project, args *model.Args) (*golang.Block, error) {
	editURL := "/" + m.Route()
	for _, pk := range m.PKs() {
		editURL += "/{%% " + pk.ToGoString("p.Model.") + " %%}"
	}

	delMsg := fmt.Sprintf("Are you sure you wish to delete %s [{%%%%s p.Model.String() %%%%}]?", m.TitleLower())

	ret := golang.NewBlock("EditBody", "func")
	ret.W("{%% func (p *Edit) Body(as *app.State, ps *cutil.PageState) %%}")
	ret.W("  <div class=\"card\">")
	ret.W("    {%%- if p.IsNew -%%}")
	ret.W("    <div class=\"right\"><a href=\"/%s/random\"><button>Random</button></a></div>", m.Route())
	ret.W("    <h3>{%%= components.SVGRefIcon(`" + m.Icon + "`, ps) %%} New " + m.Title() + "</h3>")
	ret.W("    <form action=\"/%s/new\" class=\"mt\" method=\"post\">", m.Route())
	ret.W("    {%%- else -%%}")
	ret.W("    <div class=\"right\"><a href=\"{%%s p.Model.WebPath() %%}/delete\" onclick=\"return confirm('" + delMsg + "')\"><button>Delete</button></a></div>")
	ret.W("    <h3>{%%= components.SVGRefIcon(`" + m.Icon + "`, ps) %%} Edit " + m.Title() + " [{%%s p.Model.String() %%}]</h3>")
	ret.W("    <form action=\"\" method=\"post\">")
	ret.W("    {%%- endif -%%}")
	ret.W("      <table class=\"mt expanded\">")
	ret.W("        <tbody>")
	editCols := m.Columns.WithoutTag("created").WithoutTag("updated")
	for _, col := range editCols {
		id := ""
		if len(m.RelationsFor(col)) > 0 {
			id = "input-" + col.Camel()
		}
		call, err := col.ToGoEditString("p.Model.", col.Format, id, args.Enums)
		if err != nil {
			return nil, err
		}
		if col.PK {
			ret.W("          {%% if p.IsNew %%}" + call + "{%% endif %%}")
		} else {
			ret.W("          " + call)
		}
	}
	ret.W("          <tr><td colspan=\"2\"><button type=\"submit\">Save Changes</button></td></tr>")
	ret.W("        </tbody>")
	ret.W("      </table>")
	ret.W("    </form>")
	ret.W("  </div>")

	var canAutocomplete bool
	for _, x := range m.Relations {
		if len(x.Src) == 1 {
			canAutocomplete = true
		}
	}
	if canAutocomplete {
		ret.W("  <script>")
		ret.W("    document.addEventListener(\"DOMContentLoaded\", function() {")
		for _, rel := range m.Relations {
			if len(rel.Src) != 1 {
				continue
			}
			relScript, err := exportViewEditRelation(m, rel, p, args)
			if err != nil {
				return nil, err
			}
			if relScript != "" {
				ret.W(relScript)
			}
		}
		ret.W("    });")
		ret.W("  </script>")
	}

	ret.W("{%% endfunc %%}")
	return ret, nil
}

func exportViewEditRelation(m *model.Model, rel *model.Relation, p *project.Project, args *model.Args) (string, error) {
	relModel := args.Models.Get(rel.Table)
	if !relModel.HasTag("search") {
		return "", nil
	}
	if !relModel.CanTraverseRelation() {
		return "", nil
	}
	src := m.Columns.Get(rel.Src[0])
	tgt := relModel.Columns.Get(rel.Tgt[0])

	title := ""
	if titles := relModel.Columns.Searches(); len(titles) > 0 {
		var refs []string
		for _, title := range titles {
			if !title.PK {
				refs = append(refs, fmt.Sprintf("o[%q]", title.Camel()))
			}
		}
		title = fmt.Sprintf(`(o) => %s + " (" + o[%q] + ")"`, strings.Join(refs, " + \" / \" + "), tgt.Camel())
	} else {
		title = fmt.Sprintf("(o) => o[%q]", tgt.Camel())
	}

	val := fmt.Sprintf("(o) => o[%q]", tgt.Camel())

	msg := `      %s.autocomplete(document.getElementById("input-%s"), "/%s?%s.l=10", "q", %s, %s);`
	return fmt.Sprintf(msg, p.CleanKey(), src.Camel(), relModel.Route(), relModel.Camel(), title, val), nil
}
