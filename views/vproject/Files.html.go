// Code generated by qtc from "Files.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vproject/Files.html:1
package vproject

//line views/vproject/Files.html:1
import (
	"path/filepath"

	"projectforge.dev/projectforge/app"
	"projectforge.dev/projectforge/app/controller/cutil"
	"projectforge.dev/projectforge/app/project"
	"projectforge.dev/projectforge/views/layout"
	"projectforge.dev/projectforge/views/vfile"
)

//line views/vproject/Files.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vproject/Files.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vproject/Files.html:11
type Files struct {
	layout.Basic
	Project *project.Project
	Path    []string
}

//line views/vproject/Files.html:17
func (p *Files) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vproject/Files.html:17
	qw422016.N().S(`
`)
//line views/vproject/Files.html:19
	prj := p.Project
	fs := as.Services.Projects.GetFilesystem(prj)
	u := "/p/" + prj.Key + "/fs"
	pth := filepath.Join(p.Path...)

//line views/vproject/Files.html:23
	qw422016.N().S(`  `)
//line views/vproject/Files.html:24
	StreamSummary(qw422016, prj, nil, nil, nil, ps)
//line views/vproject/Files.html:24
	qw422016.N().S(`
`)
//line views/vproject/Files.html:25
	if fs.IsDir(pth) {
//line views/vproject/Files.html:26
		files := fs.ListFiles(pth, nil, ps.Logger)

//line views/vproject/Files.html:26
		qw422016.N().S(`  <div class="card">
    <div class="right"><a href="/p/`)
//line views/vproject/Files.html:28
		qw422016.E().S(prj.Key)
//line views/vproject/Files.html:28
		qw422016.N().S(`/stats`)
//line views/vproject/Files.html:28
		if pth != `` {
//line views/vproject/Files.html:28
			qw422016.N().S(`?dir=`)
//line views/vproject/Files.html:28
			qw422016.N().U(pth)
//line views/vproject/Files.html:28
		}
//line views/vproject/Files.html:28
		qw422016.N().S(`"><button>File Stats</button></a></div>
    `)
//line views/vproject/Files.html:29
		vfile.StreamList(qw422016, p.Path, files, fs, u, as, ps)
//line views/vproject/Files.html:29
		qw422016.N().S(`
  </div>
`)
//line views/vproject/Files.html:31
	} else {
//line views/vproject/Files.html:33
		b, err := fs.ReadFile(pth)
		if err != nil {
			panic(err)
		}

//line views/vproject/Files.html:37
		qw422016.N().S(`  <div class="card">
    `)
//line views/vproject/Files.html:39
		vfile.StreamDetail(qw422016, p.Path, b, u, as, ps)
//line views/vproject/Files.html:39
		qw422016.N().S(`
  </div>
`)
//line views/vproject/Files.html:41
	}
//line views/vproject/Files.html:42
}

//line views/vproject/Files.html:42
func (p *Files) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vproject/Files.html:42
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vproject/Files.html:42
	p.StreamBody(qw422016, as, ps)
//line views/vproject/Files.html:42
	qt422016.ReleaseWriter(qw422016)
//line views/vproject/Files.html:42
}

//line views/vproject/Files.html:42
func (p *Files) Body(as *app.State, ps *cutil.PageState) string {
//line views/vproject/Files.html:42
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vproject/Files.html:42
	p.WriteBody(qb422016, as, ps)
//line views/vproject/Files.html:42
	qs422016 := string(qb422016.B)
//line views/vproject/Files.html:42
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vproject/Files.html:42
	return qs422016
//line views/vproject/Files.html:42
}
