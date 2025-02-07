// Code generated by qtc from "Request.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vadmin/Request.html:2
package vadmin

//line views/vadmin/Request.html:2
import (
	"github.com/valyala/fasthttp"

	"projectforge.dev/projectforge/app"
	"projectforge.dev/projectforge/app/controller/cutil"
	"projectforge.dev/projectforge/views/components"
	"projectforge.dev/projectforge/views/layout"
)

//line views/vadmin/Request.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vadmin/Request.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vadmin/Request.html:11
type Request struct {
	layout.Basic
	RC *fasthttp.RequestCtx
}

//line views/vadmin/Request.html:16
func (p *Request) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/Request.html:16
	qw422016.N().S(`
  <div class="card">
    <div class="right"><a href="#modal-ps"><button type="button">Page State</button></a></div>
    <h3>Request Debug</h3>
    <table>
      <thead>
        <tr>
          <th>Key</th>
          <th>Value</th>
        </tr>
      </thead>
      <tbody>
        <tr>
          <td>ID</td>
          <td>`)
//line views/vadmin/Request.html:30
	qw422016.N().DUL(p.RC.ID())
//line views/vadmin/Request.html:30
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <td>URL</td>
          <td>`)
//line views/vadmin/Request.html:34
	qw422016.E().S(p.RC.URI().String())
//line views/vadmin/Request.html:34
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <td>Protocol</td>
          <td>`)
//line views/vadmin/Request.html:38
	qw422016.E().S(string(p.RC.Request.URI().Scheme()))
//line views/vadmin/Request.html:38
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <td>Host</td>
          <td>`)
//line views/vadmin/Request.html:42
	qw422016.E().S(string(p.RC.Request.URI().Host()))
//line views/vadmin/Request.html:42
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <td>Path</td>
          <td>`)
//line views/vadmin/Request.html:46
	qw422016.E().S(string(p.RC.Request.URI().Path()))
//line views/vadmin/Request.html:46
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <td>Query String</td>
          <td>`)
//line views/vadmin/Request.html:50
	qw422016.E().S(string(p.RC.Request.URI().QueryString()))
//line views/vadmin/Request.html:50
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <td>Body Size</td>
          <td>`)
//line views/vadmin/Request.html:54
	qw422016.N().D(len(p.RC.Request.Body()))
//line views/vadmin/Request.html:54
	qw422016.N().S(`</td>
        </tr>
      </tbody>
    </table>
  </div>
`)
//line views/vadmin/Request.html:59
	if p.RC.Request.Header.Len() > 0 {
//line views/vadmin/Request.html:61
		hd := make(map[string]string, p.RC.Request.Header.Len())
		p.RC.Request.Header.VisitAll(func(k, v []byte) {
			hd[string(k)] = string(v)
		})

//line views/vadmin/Request.html:65
		qw422016.N().S(`  <div class="card">
    <h3>Headers</h3>
    <table>
      <thead>
        <tr>
          <th>Key</th>
          <th>Value</th>
        </tr>
      </thead>
      <tbody>
`)
//line views/vadmin/Request.html:76
		for k, v := range hd {
//line views/vadmin/Request.html:76
			qw422016.N().S(`        <tr>
          <td class="nowrap">`)
//line views/vadmin/Request.html:78
			qw422016.E().S(k)
//line views/vadmin/Request.html:78
			qw422016.N().S(`</td>
          <td>`)
//line views/vadmin/Request.html:79
			qw422016.E().S(v)
//line views/vadmin/Request.html:79
			qw422016.N().S(`</td>
        </tr>
`)
//line views/vadmin/Request.html:81
		}
//line views/vadmin/Request.html:81
		qw422016.N().S(`      </tbody>
    </table>
  </div>
`)
//line views/vadmin/Request.html:85
	}
//line views/vadmin/Request.html:85
	qw422016.N().S(`  `)
//line views/vadmin/Request.html:86
	components.StreamJSONModal(qw422016, "ps", "Page State", ps, 1)
//line views/vadmin/Request.html:86
	qw422016.N().S(`
`)
//line views/vadmin/Request.html:87
}

//line views/vadmin/Request.html:87
func (p *Request) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/Request.html:87
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vadmin/Request.html:87
	p.StreamBody(qw422016, as, ps)
//line views/vadmin/Request.html:87
	qt422016.ReleaseWriter(qw422016)
//line views/vadmin/Request.html:87
}

//line views/vadmin/Request.html:87
func (p *Request) Body(as *app.State, ps *cutil.PageState) string {
//line views/vadmin/Request.html:87
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vadmin/Request.html:87
	p.WriteBody(qb422016, as, ps)
//line views/vadmin/Request.html:87
	qs422016 := string(qb422016.B)
//line views/vadmin/Request.html:87
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vadmin/Request.html:87
	return qs422016
//line views/vadmin/Request.html:87
}
