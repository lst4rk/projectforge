package gomodel

import (
	"fmt"
	"strings"

	"github.com/kyleu/projectforge/app/export/files/helper"
	"github.com/kyleu/projectforge/app/export/golang"
	"github.com/kyleu/projectforge/app/export/model"
	"github.com/kyleu/projectforge/app/file"
	"github.com/kyleu/projectforge/app/util"
	"github.com/pkg/errors"
)

func DTO(m *model.Model, args *model.Args) (*file.File, error) {
	g := golang.NewFile(m.Package, []string{"app", m.Package}, "dto")
	for _, imp := range helper.ImportsForTypes("dto", m.Columns.Types()...) {
		g.AddImport(imp)
	}
	g.AddImport(helper.ImpStrings)
	if tc, err := modelTableCols(m, g); err == nil {
		g.AddBlocks(tc)
	} else {
		return nil, err
	}
	g.AddBlocks(modelDTO(m), modelDTOToModel(m), modelDTOArray(), modelDTOArrayTransformer(m))
	return g.Render()
}

func modelTableCols(m *model.Model, g *golang.File) (*golang.Block, error) {
	ret := golang.NewBlock("Columns", "procedural")
	ret.W("var (")
	ret.W("\ttable         = %q", m.Name)
	ret.W("\tcolumns       = []string{%s}", strings.Join(util.StringArrayQuoted(m.Columns.Names()), ", "))
	ret.W("\tcolumnsString = strings.Join(columns, \", \")")
	if m.IsRevision() {
		g.AddImport(helper.ImpFmt)
		hc := m.HistoryColumns(true)
		ret.W("")
		ret.W("\tcolumnsCore     = []string{%s}", strings.Join(hc.Const.NamesQuoted(), ", "))
		ret.W("\tcolumns%s = []string{%s}", hc.Col.Proper(), strings.Join(hc.Var.NamesQuoted(), ", "))
		ret.W("")
		ret.W("\ttable%s = table + \"_%s\"", hc.Col.Proper(), hc.Col.Name)
		joinClause := fmt.Sprintf("%%%%q %s join %%%%q %sr on ", m.FirstLetter(), m.FirstLetter())
		var joins []string
		for idx, col := range hc.Const {
			if col.PK || col.HasTag("current_revision") {
				rCol := hc.Var[idx]
				if !(rCol.PK || rCol.HasTag(model.RevisionType)) {
					return nil, errors.Errorf("invalid revision column [%s] at index [%d]", rCol.Name, idx)
				}
				joins = append(joins, fmt.Sprintf("%s.%s = %sr.%s", m.FirstLetter(), col.Name, m.FirstLetter(), rCol.Name))
			}
		}
		joinClause += strings.Join(joins, " and ")
		ret.W("\ttablesJoined  = fmt.Sprintf(%q, table, table%s)", joinClause, hc.Col.Proper())
	}
	ret.W(")")
	return ret, nil
}

func modelDTO(m *model.Model) *golang.Block {
	ret := golang.NewBlock(m.Proper(), "struct")
	ret.W("type dto struct {")
	maxColLength := util.StringArrayMaxLength(m.Columns.CamelNames())
	maxTypeLength := m.Columns.MaxGoDTOKeyLength()
	for _, c := range m.Columns {
		ret.W("\t%s %s `db:%q`", util.StringPad(c.Proper(), maxColLength), util.StringPad(c.ToGoDTOType(), maxTypeLength), c.Name)
	}
	ret.W("}")
	return ret
}

func modelDTOToModel(m *model.Model) *golang.Block {
	ret := golang.NewBlock(m.Proper(), "func")
	ret.W("func (d *dto) To%s() *%s {", m.Proper(), m.Proper())
	ret.W("\tif d == nil {")
	ret.W("\t\treturn nil")
	ret.W("\t}")
	refs := make([]string, 0, len(m.Columns))
	for _, c := range m.Columns {
		switch c.Type.Key {
		case model.TypeMap.Key:
			ret.W("\t%sArg := util.ValueMap{}", c.Camel())
			ret.W("\t_ = util.FromJSON(d.%s, &%sArg)", c.Proper(), c.Camel())
			refs = append(refs, fmt.Sprintf("%s: %sArg", c.Proper(), c.Camel()))
		default:
			refs = append(refs, fmt.Sprintf("%s: d.%s", c.Proper(), c.Proper()))
		}
	}
	ret.W("\treturn &%s{%s}", m.Proper(), strings.Join(refs, ", "))
	ret.W("}")
	return ret
}

func modelDTOArray() *golang.Block {
	ret := golang.NewBlock("DTOArray", "type")
	ret.W("type dtos []*dto")
	return ret
}

func modelDTOArrayTransformer(m *model.Model) *golang.Block {
	ret := golang.NewBlock(fmt.Sprintf("DTOTo%s", m.ProperPlural()), "type")
	ret.W("func (x dtos) To%s() %s {", m.ProperPlural(), m.ProperPlural())
	ret.W("\tret := make(%s, 0, len(x))", m.ProperPlural())
	ret.W("\tfor _, d := range x {")
	ret.W("\t\tret = append(ret, d.To%s())", m.Proper())
	ret.W("\t}")
	ret.W("\treturn ret")
	ret.W("}")
	return ret
}
