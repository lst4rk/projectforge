package controller

import (
	"fmt"

	"projectforge.dev/projectforge/app/project/export/enum"
	"projectforge.dev/projectforge/app/project/export/files/helper"
	"projectforge.dev/projectforge/app/project/export/golang"
	"projectforge.dev/projectforge/app/project/export/model"
)

func controllerList(m *model.Model, grp *model.Column, models model.Models, enums enum.Enums, g *golang.File, prefix string) (*golang.Block, error) {
	ret := blockFor(m, prefix, grp, 0, "list")
	meth := "List"
	grpArgs := ""
	if grp != nil {
		meth = "GetBy" + grp.Title()
		grpArgs = ", " + grp.Camel() + "Arg"
		controllerArgFor(grp, ret, "\"\"", 2)
	}

	suffix := ""
	if m.IsSoftDelete() {
		suffix = ", " + incDel
	}
	if len(m.AllSearches()) > 0 {
		g.AddImport(helper.ImpStrings)
		ret.W("\t\tq := strings.TrimSpace(string(rc.URI().QueryArgs().Peek(\"q\")))")
	}
	ret.W("\t\tprms := ps.Params.Get(%q, nil, ps.Logger).Sanitize(%q)", m.Package, m.Package)
	if grpArgs == "" && len(m.AllSearches()) > 0 {
		ret.W("\t\tvar ret %s.%s", m.Package, m.ProperPlural())
		ret.W("\t\tvar err error")
		ret.W("\t\tif q == \"\" {")
		ret.W("\t\t\tret, err = as.Services.%s.%s(ps.Context, nil, prms%s, ps.Logger)", m.Proper(), meth, suffix)
		ret.W("\t\t} else {")
		ret.W("\t\t\tret, err = as.Services.%s.Search(ps.Context, q, nil, prms%s, ps.Logger)", m.Proper(), suffix)
		ret.W("\t\t}")
	} else {
		ret.W("\t\tret, err := as.Services.%s.%s(ps.Context, nil%s, prms%s, ps.Logger)", m.Proper(), meth, grpArgs, suffix)
	}
	ret.W("\t\tif err != nil {")
	ret.W("\t\t\treturn \"\", err")
	ret.W("\t\t}")
	ret.W("\t\tps.Title = %q", m.TitlePlural())
	ret.W("\t\tps.Data = ret")

	var toStrings string
	for _, rel := range m.Relations {
		relModel := models.Get(rel.Table)
		if !relModel.CanTraverseRelation() {
			continue
		}
		srcCol := m.Columns.Get(rel.Src[0])
		for _, imp := range helper.ImportsForTypes("go", srcCol.Type) {
			g.AddImport(imp)
		}
		tgtCol := relModel.Columns.Get(rel.Tgt[0])
		for _, imp := range helper.ImportsForTypes("go", srcCol.Type) {
			g.AddImport(imp)
		}
		gt, err := model.ToGoType(srcCol.Type, srcCol.Nullable, m.Package, enums)
		if err != nil {
			return nil, err
		}

		ret.W("\t\t%sIDs := make([]%s, 0, len(ret))", relModel.Camel(), gt)
		ret.W("\t\tfor _, x := range ret {")
		ret.W("\t\t\t%sIDs = append(%sIDs, x.%s)", relModel.Camel(), relModel.Camel(), srcCol.Proper())
		ret.W("\t\t}")
		suffix := ""
		if relModel.IsSoftDelete() {
			suffix = ", false"
		}
		c := fmt.Sprintf("%sIDs", relModel.Camel())
		if srcCol.Nullable && !tgtCol.Nullable {
			c = "util.ArrayDefererence(" + c + ")"
		}
		call := "\t\t%s, err := as.Services.%s.GetMultiple(ps.Context, nil%s, ps.Logger, %s...)"
		ret.W(call, relModel.Plural(), relModel.Proper(), suffix, c)
		ret.W("\t\tif err != nil {")
		ret.W("\t\t\treturn \"\", err")
		ret.W("\t\t}")

		toStrings += fmt.Sprintf(", %s: %s", relModel.ProperPlural(), relModel.Plural())
	}
	var searchSuffix string
	if len(m.AllSearches()) > 0 {
		searchSuffix += ", SearchQuery: q"
	}
	ret.W("\t\tpage := &v%s.List{Models: ret%s, Params: ps.Params%s}", m.Package, toStrings, searchSuffix)
	render := "\t\treturn %sRender(rc, as, page, ps, %s%s)"
	ret.W(render, prefix, m.Breadcrumbs(), grp.BC())
	ret.W("\t})")
	ret.W("}")
	return ret, nil
}
