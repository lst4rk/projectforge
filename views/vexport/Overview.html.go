// Code generated by qtc from "Overview.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vexport/Overview.html:1
package vexport

//line views/vexport/Overview.html:1
import (
	"fmt"
	"strings"

	"projectforge.dev/projectforge/app"
	"projectforge.dev/projectforge/app/controller/cutil"
	"projectforge.dev/projectforge/app/project"
	"projectforge.dev/projectforge/app/project/export/model"
	"projectforge.dev/projectforge/app/util"
	"projectforge.dev/projectforge/views/components"
	"projectforge.dev/projectforge/views/layout"
	"projectforge.dev/projectforge/views/vutil"
)

//line views/vexport/Overview.html:15
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vexport/Overview.html:15
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vexport/Overview.html:15
type Overview struct {
	layout.Basic
	Project *project.Project
	Args    *model.Args
}

//line views/vexport/Overview.html:21
func (p *Overview) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vexport/Overview.html:21
	qw422016.N().S(`
  <div class="card">
    <div class="right"><a href="/p/`)
//line views/vexport/Overview.html:23
	qw422016.E().S(p.Project.Key)
//line views/vexport/Overview.html:23
	qw422016.N().S(`/export/config"><button>`)
//line views/vexport/Overview.html:23
	components.StreamSVGRef(qw422016, "edit", 15, 15, "icon", ps)
//line views/vexport/Overview.html:23
	qw422016.N().S(`Edit</button></a></div>
    <h3>`)
//line views/vexport/Overview.html:24
	components.StreamSVGRefIcon(qw422016, `print`, ps)
//line views/vexport/Overview.html:24
	qw422016.N().S(`Export Configuration</h3>
    `)
//line views/vexport/Overview.html:25
	components.StreamJSON(qw422016, p.Args.Config)
//line views/vexport/Overview.html:25
	qw422016.N().S(`
  </div>
  <div class="card">
    <div class="right"><a href="/p/`)
//line views/vexport/Overview.html:28
	qw422016.E().S(p.Project.Key)
//line views/vexport/Overview.html:28
	qw422016.N().S(`/export/groups"><button>`)
//line views/vexport/Overview.html:28
	components.StreamSVGRef(qw422016, "edit", 15, 15, "icon", ps)
//line views/vexport/Overview.html:28
	qw422016.N().S(`Edit</button></a></div>
    <h3>`)
//line views/vexport/Overview.html:29
	components.StreamSVGRefIcon(qw422016, `users`, ps)
//line views/vexport/Overview.html:29
	qw422016.N().S(`Groups</h3>
    <div class="mt">
`)
//line views/vexport/Overview.html:31
	if len(p.Args.Groups) == 0 {
//line views/vexport/Overview.html:31
		qw422016.N().S(`      <em>no groups defined</em>
`)
//line views/vexport/Overview.html:33
	} else {
//line views/vexport/Overview.html:33
		qw422016.N().S(`      `)
//line views/vexport/Overview.html:34
		StreamGroupList(qw422016, p.Args.Groups, 2)
//line views/vexport/Overview.html:34
		qw422016.N().S(`
`)
//line views/vexport/Overview.html:35
	}
//line views/vexport/Overview.html:35
	qw422016.N().S(`    </div>
  </div>
  <div class="card">
    <div class="right">
      <a href="/p/`)
//line views/vexport/Overview.html:40
	qw422016.E().S(p.Project.Key)
//line views/vexport/Overview.html:40
	qw422016.N().S(`/export/models/create/derive"><button>`)
//line views/vexport/Overview.html:40
	components.StreamSVGRef(qw422016, "dna", 15, 15, "icon", ps)
//line views/vexport/Overview.html:40
	qw422016.N().S(`Derive</button></a>
      <a href="/p/`)
//line views/vexport/Overview.html:41
	qw422016.E().S(p.Project.Key)
//line views/vexport/Overview.html:41
	qw422016.N().S(`/export/models/create/new"><button>`)
//line views/vexport/Overview.html:41
	components.StreamSVGRef(qw422016, "file", 15, 15, "icon", ps)
//line views/vexport/Overview.html:41
	qw422016.N().S(`New</button></a>
    </div>
    <h3>`)
//line views/vexport/Overview.html:43
	components.StreamSVGRefIcon(qw422016, `list`, ps)
//line views/vexport/Overview.html:43
	qw422016.N().S(`Models</h3>
    `)
//line views/vexport/Overview.html:44
	StreamModelList(qw422016, p.Args.Models, fmt.Sprintf("/p/%s/export/models", p.Project.Key), as, ps)
//line views/vexport/Overview.html:44
	qw422016.N().S(`
  </div>
`)
//line views/vexport/Overview.html:46
}

//line views/vexport/Overview.html:46
func (p *Overview) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vexport/Overview.html:46
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vexport/Overview.html:46
	p.StreamBody(qw422016, as, ps)
//line views/vexport/Overview.html:46
	qt422016.ReleaseWriter(qw422016)
//line views/vexport/Overview.html:46
}

//line views/vexport/Overview.html:46
func (p *Overview) Body(as *app.State, ps *cutil.PageState) string {
//line views/vexport/Overview.html:46
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vexport/Overview.html:46
	p.WriteBody(qb422016, as, ps)
//line views/vexport/Overview.html:46
	qs422016 := string(qb422016.B)
//line views/vexport/Overview.html:46
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vexport/Overview.html:46
	return qs422016
//line views/vexport/Overview.html:46
}

//line views/vexport/Overview.html:48
func StreamGroupList(qw422016 *qt422016.Writer, groups model.Groups, indent int) {
//line views/vexport/Overview.html:49
	vutil.StreamIndent(qw422016, true, indent)
//line views/vexport/Overview.html:49
	qw422016.N().S(`<ul>`)
//line views/vexport/Overview.html:51
	for _, g := range groups {
//line views/vexport/Overview.html:52
		vutil.StreamIndent(qw422016, true, indent+1)
//line views/vexport/Overview.html:52
		qw422016.N().S(`<li>`)
//line views/vexport/Overview.html:54
		vutil.StreamIndent(qw422016, true, indent+2)
//line views/vexport/Overview.html:55
		qw422016.E().S(g.Key)
//line views/vexport/Overview.html:56
		if len(g.Children) > 0 {
//line views/vexport/Overview.html:57
			StreamGroupList(qw422016, g.Children, indent+3)
//line views/vexport/Overview.html:58
		}
//line views/vexport/Overview.html:59
		vutil.StreamIndent(qw422016, true, indent+1)
//line views/vexport/Overview.html:59
		qw422016.N().S(`</li>`)
//line views/vexport/Overview.html:61
	}
//line views/vexport/Overview.html:62
	vutil.StreamIndent(qw422016, true, indent)
//line views/vexport/Overview.html:62
	qw422016.N().S(`</ul>`)
//line views/vexport/Overview.html:64
}

//line views/vexport/Overview.html:64
func WriteGroupList(qq422016 qtio422016.Writer, groups model.Groups, indent int) {
//line views/vexport/Overview.html:64
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vexport/Overview.html:64
	StreamGroupList(qw422016, groups, indent)
//line views/vexport/Overview.html:64
	qt422016.ReleaseWriter(qw422016)
//line views/vexport/Overview.html:64
}

//line views/vexport/Overview.html:64
func GroupList(groups model.Groups, indent int) string {
//line views/vexport/Overview.html:64
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vexport/Overview.html:64
	WriteGroupList(qb422016, groups, indent)
//line views/vexport/Overview.html:64
	qs422016 := string(qb422016.B)
//line views/vexport/Overview.html:64
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vexport/Overview.html:64
	return qs422016
//line views/vexport/Overview.html:64
}

//line views/vexport/Overview.html:66
func StreamModelList(qw422016 *qt422016.Writer, models model.Models, urlPrefix string, as *app.State, ps *cutil.PageState) {
//line views/vexport/Overview.html:66
	qw422016.N().S(`
  <table class="mt min-200 full-width">
    <tbody>
`)
//line views/vexport/Overview.html:69
	for _, m := range models {
//line views/vexport/Overview.html:71
		u := fmt.Sprintf("%s/%s", urlPrefix, m.Name)
		var prefix string
		if len(m.Group) > 0 {
			prefix += strings.Join(m.Group, "/") + ", "
		}
		if len(m.SeedData) > 0 {
			prefix += fmt.Sprintf("%d seed rows, ", len(m.SeedData))
		}

//line views/vexport/Overview.html:79
		qw422016.N().S(`      <tr>
        <td class="shrink"><a href="`)
//line views/vexport/Overview.html:81
		qw422016.E().S(u)
//line views/vexport/Overview.html:81
		qw422016.N().S(`">`)
//line views/vexport/Overview.html:81
		components.StreamSVGRef(qw422016, m.IconSafe(), 15, 15, ``, ps)
//line views/vexport/Overview.html:81
		qw422016.N().S(`</a> <a href="`)
//line views/vexport/Overview.html:81
		qw422016.E().S(u)
//line views/vexport/Overview.html:81
		qw422016.N().S(`">`)
//line views/vexport/Overview.html:81
		qw422016.E().S(m.Title())
//line views/vexport/Overview.html:81
		qw422016.N().S(`</a></td>
        <td style="text-align: right;"><em>`)
//line views/vexport/Overview.html:82
		qw422016.E().S(prefix)
//line views/vexport/Overview.html:82
		qw422016.N().D(len(m.Columns))
//line views/vexport/Overview.html:82
		qw422016.N().S(` `)
//line views/vexport/Overview.html:82
		qw422016.E().S(util.StringPluralMaybe("field", len(m.Columns)))
//line views/vexport/Overview.html:82
		qw422016.N().S(`</em></td>
      </tr>
`)
//line views/vexport/Overview.html:84
	}
//line views/vexport/Overview.html:84
	qw422016.N().S(`    </tbody>
  </table>
`)
//line views/vexport/Overview.html:87
}

//line views/vexport/Overview.html:87
func WriteModelList(qq422016 qtio422016.Writer, models model.Models, urlPrefix string, as *app.State, ps *cutil.PageState) {
//line views/vexport/Overview.html:87
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vexport/Overview.html:87
	StreamModelList(qw422016, models, urlPrefix, as, ps)
//line views/vexport/Overview.html:87
	qt422016.ReleaseWriter(qw422016)
//line views/vexport/Overview.html:87
}

//line views/vexport/Overview.html:87
func ModelList(models model.Models, urlPrefix string, as *app.State, ps *cutil.PageState) string {
//line views/vexport/Overview.html:87
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vexport/Overview.html:87
	WriteModelList(qb422016, models, urlPrefix, as, ps)
//line views/vexport/Overview.html:87
	qs422016 := string(qb422016.B)
//line views/vexport/Overview.html:87
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vexport/Overview.html:87
	return qs422016
//line views/vexport/Overview.html:87
}
