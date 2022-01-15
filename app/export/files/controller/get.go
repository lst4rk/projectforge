package controller

import (
	"fmt"
	"strings"

	"github.com/kyleu/projectforge/app/export/golang"
	"github.com/kyleu/projectforge/app/export/model"
	"github.com/kyleu/projectforge/app/util"
)

const incDel = "cutil.RequestCtxBool(rc, \"includeDeleted\")"

func controllerList(m *model.Model, grp *model.Column) *golang.Block {
	ret := blockFor(m, grp, "list")
	meth := "List"
	grpArgs := ""
	if grp != nil {
		meth = "GetBy" + grp.Title()
		grpArgs = ", " + grp.Camel() + "Arg"
		controllerArgFor(grp, ret, "\"\"", 2)
	}

	ret.W("\t\tps.Title = %q", util.StringToPlural(m.Proper()))
	ret.W("\t\tparams := cutil.ParamSetFromRequest(rc)")
	suffix := ""
	if m.IsSoftDelete() {
		suffix = ", " + incDel
	}
	msg := "\t\tret, err := as.Services.%s.%s(ps.Context, nil%s, params.Get(%q, nil, ps.Logger)%s)"
	ret.W(msg, m.Proper(), meth, grpArgs, m.Package, suffix)
	ret.W("\t\tif err != nil {")
	ret.W("\t\t\treturn \"\", err")
	ret.W("\t\t}")
	ret.W("\t\tps.Data = ret")
	ret.W("\t\treturn render(rc, as, &v%s.List{Models: ret, Params: params}, ps, %q%s)", m.Package, m.Package, grp.BC())
	ret.W("\t})")
	ret.W("}")
	return ret
}

func controllerDetail(m *model.Model, grp *model.Column) *golang.Block {
	ret := blockFor(m, grp, "detail")
	grpHistory := ""
	if grp != nil {
		controllerArgFor(grp, ret, "\"\"", 2)
		grpHistory = fmt.Sprintf(", %q", grp.Camel())
	}
	ret.W("\t\tret, err := %sFromPath(rc, as, ps)", m.Package)
	ret.W("\t\tif err != nil {")
	ret.W("\t\t\treturn \"\", err")
	ret.W("\t\t}")
	checkGrp(ret, grp)
	if m.IsRevision() {
		ret.W("\t\tparams := cutil.ParamSetFromRequest(rc)")
		var prms []string
		for _, pk := range m.PKs() {
			prms = append(prms, "ret."+pk.Proper())
		}
		prmsStr := strings.Join(prms, ", ")
		msg := "\t\trevisions, err := as.Services.%s.GetAllRevisions(ps.Context, nil, %s, params.Get(%q, nil, ps.Logger), false)"
		ret.W(msg, m.Proper(), prmsStr, m.Package)
	}
	ret.W("\t\tps.Title = ret.String()")
	ret.W("\t\tps.Data = ret")
	suffix := ""
	if m.IsRevision() {
		revCol := m.HistoryColumn()
		suffix = fmt.Sprintf(", %s: %s, Params: params", revCol.ProperPlural(), revCol.Plural())
	}
	ret.W("\t\treturn render(rc, as, &v%s.Detail{Model: ret%s}, ps, %q%s, ret.String())", m.Package, suffix, m.Package, grpHistory)
	ret.W("\t})")
	ret.W("}")
	return ret
}

func checkGrp(ret *golang.Block, grp *model.Column, override ...string) {
	if grp == nil {
		return
	}
	x := "ret"
	if len(override) > 0 {
		x = override[0]
	}
	ret.W("\t\tif %s.%s != %sArg {", x, grp.Proper(), grp.Camel())
	ret.W("\t\t\treturn \"\", errors.New(\"unauthorized: incorrect [%s]\")", grp.Camel())
	ret.W("\t\t}")
}

func controllerModelFromPath(m *model.Model) *golang.Block {
	ret := golang.NewBlock(m.Proper()+"FromPath", "func")
	ret.W("func %sFromPath(rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (*%s, error) {", m.Package, m.ClassRef())
	pks := m.PKs()
	for _, col := range pks {
		controllerArgFor(col, ret, "nil", 1)
	}
	args := make([]string, 0, len(pks))
	for _, x := range pks {
		args = append(args, x.Camel()+"Arg")
	}
	suffix := ""
	if m.IsSoftDelete() {
		suffix = ", includeDeleted"
		ret.W("\tincludeDeleted := rc.UserValue(\"includeDeleted\") != nil || " + incDel)
	}
	ret.W("\treturn as.Services.%s.Get(ps.Context, nil, %s%s)", m.Proper(), strings.Join(args, ", "), suffix)
	ret.W("}")

	return ret
}
