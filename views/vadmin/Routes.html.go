// Code generated by qtc from "Routes.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vadmin/Routes.html:2
package vadmin

//line views/vadmin/Routes.html:2
import (
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"

	"projectforge.dev/projectforge/app"
	"projectforge.dev/projectforge/app/controller/cutil"
	"projectforge.dev/projectforge/views/layout"
)

//line views/vadmin/Routes.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vadmin/Routes.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vadmin/Routes.html:11
type Routes struct {
	layout.Basic
	Routes map[string][]string
}

//line views/vadmin/Routes.html:16
func (p *Routes) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/Routes.html:16
	qw422016.N().S(`
  <div class="card">
    <h3>HTTP Routes</h3>
    <ul class="mt">
`)
//line views/vadmin/Routes.html:21
	x := maps.Keys(p.Routes)
	slices.Sort(x)

//line views/vadmin/Routes.html:24
	for _, k := range x {
//line views/vadmin/Routes.html:24
		qw422016.N().S(`      <li>
        <strong>`)
//line views/vadmin/Routes.html:26
		qw422016.E().S(k)
//line views/vadmin/Routes.html:26
		qw422016.N().S(`</strong>
        <ul>
`)
//line views/vadmin/Routes.html:28
		for _, r := range p.Routes[k] {
//line views/vadmin/Routes.html:28
			qw422016.N().S(`          <li><code>`)
//line views/vadmin/Routes.html:29
			qw422016.E().S(r)
//line views/vadmin/Routes.html:29
			qw422016.N().S(`</code></li>
`)
//line views/vadmin/Routes.html:30
		}
//line views/vadmin/Routes.html:30
		qw422016.N().S(`        </ul>
      </li>
`)
//line views/vadmin/Routes.html:33
	}
//line views/vadmin/Routes.html:33
	qw422016.N().S(`    </ul>
  </div>
`)
//line views/vadmin/Routes.html:36
}

//line views/vadmin/Routes.html:36
func (p *Routes) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/Routes.html:36
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vadmin/Routes.html:36
	p.StreamBody(qw422016, as, ps)
//line views/vadmin/Routes.html:36
	qt422016.ReleaseWriter(qw422016)
//line views/vadmin/Routes.html:36
}

//line views/vadmin/Routes.html:36
func (p *Routes) Body(as *app.State, ps *cutil.PageState) string {
//line views/vadmin/Routes.html:36
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vadmin/Routes.html:36
	p.WriteBody(qb422016, as, ps)
//line views/vadmin/Routes.html:36
	qs422016 := string(qb422016.B)
//line views/vadmin/Routes.html:36
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vadmin/Routes.html:36
	return qs422016
//line views/vadmin/Routes.html:36
}
