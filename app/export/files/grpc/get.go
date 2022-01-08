package grpc

import (
	"strings"

	"github.com/kyleu/projectforge/app/export/golang"
	"github.com/kyleu/projectforge/app/export/model"
)

func grpcList(m *model.Model, grpcArgs string, grpcRet string, ga *GRPCFileArgs) *golang.Block {
	ret := golang.NewBlock("grpcList", "func")
	ret.W("func %sList%s(%s) %s {", m.Proper(), ga.APISuffix(), grpcArgs, grpcRet)
	idClause, suffix := idClauseFor(m)
	if idClause != "" {
		ret.W(idClause)
	}
	grpcAddSection(ret, "list", 1)
	if ga.Grp == nil {
		ret.W("\tret, err := appState.Services.%s.List(p.Ctx, nil, &filter.Params{}%s)", m.Proper(), suffix)
	} else {
		ret.W("\tret, err := appState.Services.%s.Get%s(p.Ctx, nil, %s, &filter.Params{}%s)", m.Proper(), ga.APISuffix(), ga.Grp.Camel(), suffix)
	}
	ret.W("\tif err != nil {")
	ret.W("\t\treturn nil, err")
	ret.W("\t}")
	ret.W("\tprovider.SetOutput(p.TX, ret)")
	ret.W("\treturn p.TX, nil")
	ret.W("}")
	return ret
}

func grpcSearch(m *model.Model, grpcArgs string, grpcRet string, ga *GRPCFileArgs) *golang.Block {
	ret := golang.NewBlock("grpcSearch", "func")
	ret.W("func %sSearch%s(%s) %s {", m.Proper(), ga.APISuffix(), grpcArgs, grpcRet)
	idClause, suffix := idClauseFor(m)
	if idClause != "" {
		ret.W(idClause)
	}
	ret.W("\tq, _ := provider.GetString(p.R, p.TX, \"q\")")
	ret.W("\tif q == \"\" {")
	ret.W("\t\treturn nil, errors.New(\"must provide [q] in request data\")")
	ret.W("\t}")
	grpcAddSection(ret, "search", 1)
	if ga.Grp == nil {
		ret.W("\tret, err := appState.Services.%s.Search(p.Ctx, q, nil, nil%s)", m.Proper(), suffix)
	} else {
		ret.W("\tret, err := appState.Services.%s.Search%s(p.Ctx, %s, q, nil, nil%s)", m.Proper(), ga.APISuffix(), ga.Grp.Camel(), suffix)
	}
	ret.W("\tif err != nil {")
	ret.W("\t\treturn nil, err")
	ret.W("\t}")
	ret.W("\tprovider.SetOutput(p.TX, ret)")
	ret.W("\treturn p.TX, nil")
	ret.W("}")
	return ret
}

func grpcDetail(m *model.Model, grpcArgs string, grpcRet string, ga *GRPCFileArgs) *golang.Block {
	ret := golang.NewBlock("grpcDetail", "func")
	ret.W("func %sDetail%s(%s) %s {", m.Proper(), ga.APISuffix(), grpcArgs, grpcRet)
	idClause, suffix := idClauseFor(m)
	if idClause != "" {
		ret.W(idClause)
	}
	pks := m.PKs()
	ret.W("\t%s, err := %sParamsFromRequest(p.R)", strings.Join(pks.Names(), ", "), m.Camel())
	ret.W("\tif err != nil {")
	ret.W("\t\treturn nil, err")
	ret.W("\t}")
	grpcAddSection(ret, "detail", 1)
	ret.W("\tret, err := appState.Services.%s.Get(p.Ctx, nil, %s%s)", m.Proper(), strings.Join(pks.Names(), ", "), suffix)
	ret.W("\tif err != nil {")
	ret.W("\t\treturn nil, err")
	ret.W("\t}")
	ga.AddStaticCheck("ret", ret, ga.Grp)
	ret.W("\tprovider.SetOutput(p.TX, ret)")
	ret.W("\treturn p.TX, nil")
	ret.W("}")
	return ret
}
