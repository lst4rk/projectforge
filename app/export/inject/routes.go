package inject

import (
	"fmt"
	"strings"

	"github.com/kyleu/projectforge/app/export/model"
	"github.com/kyleu/projectforge/app/file"
)

func Routes(f *file.File, args *model.Args) error {
	out := make([]string, 0, 100)
	for _, m := range args.Models {
		pkNames := make([]string, 0, len(m.PKs()))
		for _, col := range m.PKs() {
			pkNames = append(pkNames, fmt.Sprintf("{%s}", col.Camel()))
		}

		for _, grp := range m.GroupedColumns() {
			pathExtra := fmt.Sprintf("/%s/{%s}", grp.Camel(), grp.Camel())
			callSuffix := fmt.Sprintf("By%s", grp.Proper())

			g := fmt.Sprintf("\tr.GET(\"/%s/%s\", %s%sList)", m.Package, grp.Camel(), m.Proper(), grp.Proper())
			l := fmt.Sprintf("\tr.GET(\"/%s%s\", %sList%s)", m.Package, pathExtra, m.Proper(), callSuffix)
			nf := fmt.Sprintf("\tr.GET(\"/%s%s/new\", %sCreateForm%s)", m.Package, pathExtra, m.Proper(), callSuffix)
			ns := fmt.Sprintf("\tr.POST(\"/%s%s/new\", %sCreate%s)", m.Package, pathExtra, m.Proper(), callSuffix)
			d := fmt.Sprintf("\tr.GET(\"/%s%s/%s\", %sDetail%s)", m.Package, pathExtra, strings.Join(pkNames, "/"), m.Proper(), callSuffix)
			ef := fmt.Sprintf("\tr.GET(\"/%s%s/%s/edit\", %sEditForm%s)", m.Package, pathExtra, strings.Join(pkNames, "/"), m.Proper(), callSuffix)
			es := fmt.Sprintf("\tr.POST(\"/%s%s/%s/edit\", %sEdit%s)", m.Package, pathExtra, strings.Join(pkNames, "/"), m.Proper(), callSuffix)
			out = append(out, g, l, nf, ns, d, ef, es, "")
		}

		l := fmt.Sprintf("\tr.GET(\"/%s\", %sList)", m.Package, m.Proper())
		nf := fmt.Sprintf("\tr.GET(\"/%s/new\", %sCreateForm)", m.Package, m.Proper())
		ns := fmt.Sprintf("\tr.POST(\"/%s/new\", %sCreate)", m.Package, m.Proper())
		d := fmt.Sprintf("\tr.GET(\"/%s/%s\", %sDetail)", m.Package, strings.Join(pkNames, "/"), m.Proper())
		ef := fmt.Sprintf("\tr.GET(\"/%s/%s/edit\", %sEditForm)", m.Package, strings.Join(pkNames, "/"), m.Proper())
		es := fmt.Sprintf("\tr.POST(\"/%s/%s/edit\", %sEdit)", m.Package, strings.Join(pkNames, "/"), m.Proper())
		dl := fmt.Sprintf("\tr.GET(\"/%s/%s/delete\", %sDelete)", m.Package, strings.Join(pkNames, "/"), m.Proper())
		out = append(out, l, nf, ns, d, ef, es, dl)
		if m.IsRevision() {
			rc := m.HistoryColumn()
			msg := "\tr.GET(\"/%s/%s/%s/{%s}\", %s%s)"
			out = append(out, fmt.Sprintf(msg, m.Package, strings.Join(pkNames, "/"), rc.Name, rc.Name, m.Proper(), rc.Proper()))
		}
	}
	content := map[string]string{"codegen": "\n" + strings.Join(out, "\n") + "\n\t// "}
	return file.Inject(f, content)
}
