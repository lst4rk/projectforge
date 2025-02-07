// Code generated by qtc from "Files.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vproject/Files.html:1
package vproject

//line views/vproject/Files.html:1
import (
	"path/filepath"
	"strings"

	"projectforge.dev/projectforge/app"
	"projectforge.dev/projectforge/app/controller/cutil"
	"projectforge.dev/projectforge/app/project"
	"projectforge.dev/projectforge/views/layout"
	"projectforge.dev/projectforge/views/vfile"
)

//line views/vproject/Files.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vproject/Files.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vproject/Files.html:12
type Files struct {
	layout.Basic
	Project *project.Project
	Path    []string
}

//line views/vproject/Files.html:18
func (p *Files) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vproject/Files.html:18
	qw422016.N().S(`
`)
//line views/vproject/Files.html:20
	prj := p.Project
	fs := as.Services.Projects.GetFilesystem(prj)
	u := "/p/" + prj.Key + "/fs"
	pth := filepath.Join(p.Path...)

//line views/vproject/Files.html:24
	qw422016.N().S(`  `)
//line views/vproject/Files.html:25
	StreamSummary(qw422016, prj, "/"+strings.Join(p.Path, "/"), nil, nil, nil, nil, ps)
//line views/vproject/Files.html:25
	qw422016.N().S(`
`)
//line views/vproject/Files.html:26
	if fs.IsDir(pth) {
//line views/vproject/Files.html:27
		files := fs.ListFiles(pth, nil, ps.Logger)

//line views/vproject/Files.html:27
		qw422016.N().S(`  <div class="card">
    <div class="right"><a href="/p/`)
//line views/vproject/Files.html:29
		qw422016.E().S(prj.Key)
//line views/vproject/Files.html:29
		qw422016.N().S(`/stats`)
//line views/vproject/Files.html:29
		if pth != `` {
//line views/vproject/Files.html:29
			qw422016.N().S(`?dir=`)
//line views/vproject/Files.html:29
			qw422016.N().U(pth)
//line views/vproject/Files.html:29
		}
//line views/vproject/Files.html:29
		qw422016.N().S(`"><button>File Stats</button></a></div>
    `)
//line views/vproject/Files.html:30
		vfile.StreamList(qw422016, p.Path, files, fs, u, as, ps)
//line views/vproject/Files.html:30
		qw422016.N().S(`
  </div>
`)
//line views/vproject/Files.html:32
	} else {
//line views/vproject/Files.html:34
		b, err := fs.ReadFile(pth)
		if err != nil {
			panic(err)
		}

//line views/vproject/Files.html:38
		qw422016.N().S(`  <div class="card">
    `)
//line views/vproject/Files.html:40
		vfile.StreamDetail(qw422016, p.Path, b, u, as, ps)
//line views/vproject/Files.html:40
		qw422016.N().S(`
  </div>
`)
//line views/vproject/Files.html:42
	}
//line views/vproject/Files.html:43
}

//line views/vproject/Files.html:43
func (p *Files) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vproject/Files.html:43
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vproject/Files.html:43
	p.StreamBody(qw422016, as, ps)
//line views/vproject/Files.html:43
	qt422016.ReleaseWriter(qw422016)
//line views/vproject/Files.html:43
}

//line views/vproject/Files.html:43
func (p *Files) Body(as *app.State, ps *cutil.PageState) string {
//line views/vproject/Files.html:43
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vproject/Files.html:43
	p.WriteBody(qb422016, as, ps)
//line views/vproject/Files.html:43
	qs422016 := string(qb422016.B)
//line views/vproject/Files.html:43
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vproject/Files.html:43
	return qs422016
//line views/vproject/Files.html:43
}
