// Code generated by qtc from "Summary.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vmodule/Summary.html:1
package vmodule

//line views/vmodule/Summary.html:1
import (
	"projectforge.dev/projectforge/app/controller/cutil"
	"projectforge.dev/projectforge/app/module"
	"projectforge.dev/projectforge/app/util"
	"projectforge.dev/projectforge/views/components"
	"projectforge.dev/projectforge/views/vsearch"
	"strings"
)

//line views/vmodule/Summary.html:10
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vmodule/Summary.html:10
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vmodule/Summary.html:10
func StreamSummary(qw422016 *qt422016.Writer, mod *module.Module, args util.ValueMap, ps *cutil.PageState, path ...string) {
//line views/vmodule/Summary.html:10
	qw422016.N().S(`
  <div class="card">
`)
//line views/vmodule/Summary.html:13
	q := ""
	if args != nil {
		q = args.GetStringOpt(`q`)
	}

//line views/vmodule/Summary.html:17
	qw422016.N().S(`    `)
//line views/vmodule/Summary.html:18
	vsearch.StreamForm(qw422016, "/m/"+mod.Key+"/search", q, "Search Files", ps)
//line views/vmodule/Summary.html:18
	qw422016.N().S(`
    <h3>`)
//line views/vmodule/Summary.html:19
	components.StreamSVGRefIcon(qw422016, mod.IconSafe(), ps)
//line views/vmodule/Summary.html:19
	qw422016.N().S(` `)
//line views/vmodule/Summary.html:19
	qw422016.E().S(mod.Title())
//line views/vmodule/Summary.html:19
	qw422016.N().S(`</h3>
    <div class="mt">
`)
//line views/vmodule/Summary.html:21
	if len(path) == 0 {
//line views/vmodule/Summary.html:21
		qw422016.N().S(`      <a href="/m/`)
//line views/vmodule/Summary.html:22
		qw422016.E().S(mod.Key)
//line views/vmodule/Summary.html:22
		qw422016.N().S(`/fs"><button>Filesystem</button></a>
`)
//line views/vmodule/Summary.html:23
	} else {
//line views/vmodule/Summary.html:24
		var ctx []string

//line views/vmodule/Summary.html:25
		for _, pth := range path {
//line views/vmodule/Summary.html:26
			ctx = append(ctx, pth)

//line views/vmodule/Summary.html:26
			qw422016.N().S(`      <a href="/m/`)
//line views/vmodule/Summary.html:27
			qw422016.E().S(mod.Key)
//line views/vmodule/Summary.html:27
			qw422016.N().S(`/fs/`)
//line views/vmodule/Summary.html:27
			qw422016.E().S(strings.Join(ctx, `/`))
//line views/vmodule/Summary.html:27
			qw422016.N().S(`"><button>`)
//line views/vmodule/Summary.html:27
			qw422016.E().S(pth)
//line views/vmodule/Summary.html:27
			qw422016.N().S(`}</button></a>
`)
//line views/vmodule/Summary.html:28
		}
//line views/vmodule/Summary.html:29
	}
//line views/vmodule/Summary.html:29
	qw422016.N().S(`      <a href="#modal-module"><button type="button">JSON</button></a>
    </div>
  </div>
  `)
//line views/vmodule/Summary.html:33
	components.StreamJSONModal(qw422016, "module", "Module JSON", mod, 1)
//line views/vmodule/Summary.html:33
	qw422016.N().S(`
`)
//line views/vmodule/Summary.html:34
}

//line views/vmodule/Summary.html:34
func WriteSummary(qq422016 qtio422016.Writer, mod *module.Module, args util.ValueMap, ps *cutil.PageState, path ...string) {
//line views/vmodule/Summary.html:34
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vmodule/Summary.html:34
	StreamSummary(qw422016, mod, args, ps, path...)
//line views/vmodule/Summary.html:34
	qt422016.ReleaseWriter(qw422016)
//line views/vmodule/Summary.html:34
}

//line views/vmodule/Summary.html:34
func Summary(mod *module.Module, args util.ValueMap, ps *cutil.PageState, path ...string) string {
//line views/vmodule/Summary.html:34
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vmodule/Summary.html:34
	WriteSummary(qb422016, mod, args, ps, path...)
//line views/vmodule/Summary.html:34
	qs422016 := string(qb422016.B)
//line views/vmodule/Summary.html:34
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vmodule/Summary.html:34
	return qs422016
//line views/vmodule/Summary.html:34
}
