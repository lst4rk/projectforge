package grpc

import (
	"strings"

	"github.com/kyleu/projectforge/app/export/golang"
	"github.com/kyleu/projectforge/app/export/model"
	"github.com/kyleu/projectforge/app/util"
	"github.com/pkg/errors"
)

type FileArgs struct {
	Class string
	Pkg   string
	CPkg  string
	API   string
	Grp   *model.Column
}

func (a FileArgs) APISuffix() string {
	if a.API == "*" {
		return ""
	}
	return "By" + util.StringToTitle(a.API)
}

func (a FileArgs) KeySuffix() string {
	if a.API == "*" {
		return ""
	}
	return "." + a.API
}

func (a FileArgs) GrpSuffix() string {
	if a.Grp == nil {
		return ""
	}
	return "By" + a.Grp.Proper()
}

func (a FileArgs) AddStaticCheck(ref string, ret *golang.Block, m *model.Model, grp *model.Column, act string) {
	if grp == nil {
		return
	}
	ret.W("\tif %s.%s != %s {", ref, grp.Proper(), grp.Camel())
	msg := "\t\treturn nil, errors.Errorf(\"unauthorized - user acting on behalf of [%%%%s] cannot %s a %s for [%%%%s]\", %s, %s.%s)"
	ret.W(msg, act, m.Camel(), grp.Camel(), ref, grp.Proper())
	ret.W("\t}")
}

func grpcAddSection(b *golang.Block, key string, indent int) {
	ind := ""
	for i := 0; i < indent; i++ {
		ind += "\t"
	}
	b.W(ind+"// $PF_SECTION_START(%s)$", key)
	b.W(ind+"// $PF_SECTION_END(%s)$", key)
}

func idClauseFor(m *model.Model) (string, string) {
	if m.IsSoftDelete() {
		return "\tincludeDeleted, _ := provider.GetBool(p.R, p.TX, \"includeDeleted\")", ", includeDeleted"
	}
	return "", ""
}

func grpcParamsFromRequest(m *model.Model, cPkg string) (*golang.Block, error) {
	ret := golang.NewBlock("grpcParamsFromRequest", "func")
	pks := m.PKs()
	ret.W("func %sParamsFromRequest(r *provider.NuevoRequest) (%s, error) {", m.Camel(), strings.Join(pks.GoTypeKeys(), ", "))
	zeroVals := strings.Join(pks.ZeroVals(), ", ")
	for _, col := range pks {
		err := grpcArgFor(col, ret, zeroVals)
		if err != nil {
			return nil, errors.Wrap(err, "")
		}
	}
	ret.W("\treturn %s, nil", strings.Join(pks.CamelNames(), ", "))
	ret.W("}")
	return ret, nil
}

func grpcArgFor(col *model.Column, b *golang.Block, zeroVals string) error {
	switch col.Type.Key {
	case model.TypeInt.Key:
		b.W("\t%s, err := provider.GetRequestInt(r, %q)", col.Camel(), col.Camel())
		b.W("\tif err != nil {")
		b.W("\t\treturn %s, err", zeroVals)
		b.W("\t}")
	case model.TypeString.Key:
		b.W("\t%s, err := provider.GetRequestString(r, %q)", col.Camel(), col.Camel())
		b.W("\tif err != nil {")
		b.W("\t\treturn %s, err", zeroVals)
		b.W("\t}")
	default:
		return errors.Errorf("unhandled gRPC arg type [%s]", col.Type.String())
	}
	return nil
}
