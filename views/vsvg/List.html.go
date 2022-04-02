// Code generated by qtc from "List.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vsvg/List.html:1
package vsvg

//line views/vsvg/List.html:1
import (
	"strings"

	"projectforge.dev/projectforge/app"
	"projectforge.dev/projectforge/app/action"
	"projectforge.dev/projectforge/app/controller/cutil"
	"projectforge.dev/projectforge/app/project"
	"projectforge.dev/projectforge/views/layout"
	"projectforge.dev/projectforge/views/vproject"
)

//line views/vsvg/List.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vsvg/List.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vsvg/List.html:12
type List struct {
	layout.Basic
	Project  *project.Project
	Keys     []string
	Contents map[string]string
}

//line views/vsvg/List.html:19
func (p *List) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsvg/List.html:19
	qw422016.N().S(`
  `)
//line views/vsvg/List.html:20
	vproject.StreamSummary(qw422016, p.Project, nil, nil, &action.TypeSVG, ps)
//line views/vsvg/List.html:20
	qw422016.N().S(`
  <div class="card">
    <h3>SVG icons</h3>
    <a href="/svg/`)
//line views/vsvg/List.html:23
	qw422016.E().S(p.Project.Key)
//line views/vsvg/List.html:23
	qw422016.N().S(`/x/build"><button>Rebuild</button></a>
  </div>
  <div class="card">
    <h3>Add a new icon</h3>
    <p>
      To add an icon to your application, enter a key from <a href="https://icons8.com/line-awesome" target="_blank">Line Awesome</a>.
      Instead of an icon key, you may also provide a URL to any SVG.
      The SVG will be rewritten with attribution.
    </p>
    <form action="/svg/`)
//line views/vsvg/List.html:32
	qw422016.E().S(p.Project.Key)
//line views/vsvg/List.html:32
	qw422016.N().S(`/x/add" method="get">
      <table>
        <tbody>
          <tr>
            <th class="shrink"><label for="input-src">Source Key</label></th>
            <td><input id="input-src" name="src" value="" /></td>
          </tr>
          <tr>
            <th class="shrink"><label for="input-tgt">Target Key</label></th>
            <td><input id="input-tgt" name="tgt" value="" /></td>
          </tr>
        </tbody>
      </table>
      <div class="mt">
        <button type="submit">Add/Update SVG</button>
      </div>
    </form>
  </div>
  <div class="card">
    <h3>Current Icons</h3>
    <div class="flex-wrap">
`)
//line views/vsvg/List.html:53
	for _, k := range p.Keys {
//line views/vsvg/List.html:53
		qw422016.N().S(`      <div class="icon-container">
        <a href="/svg/`)
//line views/vsvg/List.html:55
		qw422016.E().S(p.Project.Key)
//line views/vsvg/List.html:55
		qw422016.N().S(`/`)
//line views/vsvg/List.html:55
		qw422016.E().S(k)
//line views/vsvg/List.html:55
		qw422016.N().S(`">
          <div>`)
//line views/vsvg/List.html:56
		qw422016.N().S(strings.ReplaceAll(p.Contents[k], "svg-"+k, "svg-"+k+"_adhoc"))
//line views/vsvg/List.html:56
		qw422016.N().S(`</div>
          <div>`)
//line views/vsvg/List.html:57
		qw422016.E().S(k)
//line views/vsvg/List.html:57
		qw422016.N().S(`</div>
        </a>
      </div>
`)
//line views/vsvg/List.html:60
	}
//line views/vsvg/List.html:60
	qw422016.N().S(`    </div>
  </div>
`)
//line views/vsvg/List.html:63
}

//line views/vsvg/List.html:63
func (p *List) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsvg/List.html:63
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vsvg/List.html:63
	p.StreamBody(qw422016, as, ps)
//line views/vsvg/List.html:63
	qt422016.ReleaseWriter(qw422016)
//line views/vsvg/List.html:63
}

//line views/vsvg/List.html:63
func (p *List) Body(as *app.State, ps *cutil.PageState) string {
//line views/vsvg/List.html:63
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vsvg/List.html:63
	p.WriteBody(qb422016, as, ps)
//line views/vsvg/List.html:63
	qs422016 := string(qb422016.B)
//line views/vsvg/List.html:63
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vsvg/List.html:63
	return qs422016
//line views/vsvg/List.html:63
}
