// Code generated by qtc from "Args.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/verror/Args.html:2
package verror

//line views/verror/Args.html:2
import (
	"strconv"

	"projectforge.dev/projectforge/app"
	"projectforge.dev/projectforge/app/controller/cutil"
	"projectforge.dev/projectforge/views/components"
	"projectforge.dev/projectforge/views/layout"
)

//line views/verror/Args.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/verror/Args.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/verror/Args.html:11
type Args struct {
	layout.Basic
	URL        string
	Directions string
	ArgRes     *cutil.ArgResults
	Hidden     map[string]string
}

//line views/verror/Args.html:19
func (p *Args) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/verror/Args.html:19
	qw422016.N().S(`
  <div class="card">
    <h3>`)
//line views/verror/Args.html:21
	if p.Directions == "" {
//line views/verror/Args.html:21
		qw422016.N().S(`Enter Data`)
//line views/verror/Args.html:21
	} else {
//line views/verror/Args.html:21
		qw422016.E().S(p.Directions)
//line views/verror/Args.html:21
	}
//line views/verror/Args.html:21
	qw422016.N().S(`</h3>
    <form action="`)
//line views/verror/Args.html:22
	qw422016.E().S(p.URL)
//line views/verror/Args.html:22
	qw422016.N().S(`" method="get">
`)
//line views/verror/Args.html:23
	for k, v := range p.Hidden {
//line views/verror/Args.html:23
		qw422016.N().S(`      <input type="hidden" name="`)
//line views/verror/Args.html:24
		qw422016.E().S(k)
//line views/verror/Args.html:24
		qw422016.N().S(`" value="`)
//line views/verror/Args.html:24
		qw422016.E().S(v)
//line views/verror/Args.html:24
		qw422016.N().S(`" />
`)
//line views/verror/Args.html:25
	}
//line views/verror/Args.html:25
	qw422016.N().S(`      <table class="mt min-200 expanded">
        <tbody>
`)
//line views/verror/Args.html:28
	for _, arg := range p.ArgRes.Args {
//line views/verror/Args.html:30
		v := p.ArgRes.Values[arg.Key]
		if v == "" {
			v = arg.Default
		}

//line views/verror/Args.html:35
		switch arg.Type {
//line views/verror/Args.html:36
		case "bool":
//line views/verror/Args.html:36
			qw422016.N().S(`          `)
//line views/verror/Args.html:37
			components.StreamTableBoolean(qw422016, arg.Key, arg.Title, v == "true", 5, arg.Description)
//line views/verror/Args.html:37
			qw422016.N().S(`
`)
//line views/verror/Args.html:38
		case "textarea":
//line views/verror/Args.html:38
			qw422016.N().S(`          `)
//line views/verror/Args.html:39
			components.StreamTableTextarea(qw422016, arg.Key, "", arg.Title, 12, v, 5, arg.Description)
//line views/verror/Args.html:39
			qw422016.N().S(`
`)
//line views/verror/Args.html:40
		case "number", "int":
//line views/verror/Args.html:41
			i, _ := strconv.Atoi(v)

//line views/verror/Args.html:41
			qw422016.N().S(`          `)
//line views/verror/Args.html:42
			components.StreamTableInputNumber(qw422016, arg.Key, "", arg.Title, i, 5, arg.Description)
//line views/verror/Args.html:42
			qw422016.N().S(`
`)
//line views/verror/Args.html:43
		default:
//line views/verror/Args.html:43
			qw422016.N().S(`          `)
//line views/verror/Args.html:44
			components.StreamTableInput(qw422016, arg.Key, "", arg.Title, v, 5, arg.Description)
//line views/verror/Args.html:44
			qw422016.N().S(`
`)
//line views/verror/Args.html:45
		}
//line views/verror/Args.html:46
	}
//line views/verror/Args.html:46
	qw422016.N().S(`        </tbody>
      </table>
      <button class="mt" type="submit">Submit</button>
    </form>
  </div>
`)
//line views/verror/Args.html:52
}

//line views/verror/Args.html:52
func (p *Args) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/verror/Args.html:52
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/verror/Args.html:52
	p.StreamBody(qw422016, as, ps)
//line views/verror/Args.html:52
	qt422016.ReleaseWriter(qw422016)
//line views/verror/Args.html:52
}

//line views/verror/Args.html:52
func (p *Args) Body(as *app.State, ps *cutil.PageState) string {
//line views/verror/Args.html:52
	qb422016 := qt422016.AcquireByteBuffer()
//line views/verror/Args.html:52
	p.WriteBody(qb422016, as, ps)
//line views/verror/Args.html:52
	qs422016 := string(qb422016.B)
//line views/verror/Args.html:52
	qt422016.ReleaseByteBuffer(qb422016)
//line views/verror/Args.html:52
	return qs422016
//line views/verror/Args.html:52
}
