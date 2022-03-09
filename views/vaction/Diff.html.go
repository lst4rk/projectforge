// Code generated by qtc from "Diff.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vaction/Diff.html:1
package vaction

//line views/vaction/Diff.html:1
import (
	"strings"

	"projectforge.dev/projectforge/app"
	"projectforge.dev/projectforge/app/controller/cutil"
	"projectforge.dev/projectforge/app/diff"
)

//line views/vaction/Diff.html:9
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vaction/Diff.html:9
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vaction/Diff.html:9
func streamrenderDiffs(qw422016 *qt422016.Writer, prjKey string, diffs diff.Diffs, as *app.State, ps *cutil.PageState) {
//line views/vaction/Diff.html:9
	qw422016.N().S(`<h4>Diffs</h4><table><thead><tr><th class="shrink">Path</th><th class="shrink">Status</th><th class="shrink">Patch</th></tr></thead><tbody>`)
//line views/vaction/Diff.html:20
	for _, d := range diffs {
//line views/vaction/Diff.html:20
		qw422016.N().S(`<tr><td class="shrink">`)
//line views/vaction/Diff.html:22
		qw422016.E().S(d.Path)
//line views/vaction/Diff.html:22
		qw422016.N().S(`</td><td>`)
//line views/vaction/Diff.html:24
		qw422016.E().S(d.Status.String())
//line views/vaction/Diff.html:25
		if true {
//line views/vaction/Diff.html:25
			qw422016.N().S(`<form action="/run/`)
//line views/vaction/Diff.html:26
			qw422016.E().S(prjKey)
//line views/vaction/Diff.html:26
			qw422016.N().S(`/merge" method="get"><input type="hidden" name="file" value="`)
//line views/vaction/Diff.html:27
			qw422016.E().S(d.Path)
//line views/vaction/Diff.html:27
			qw422016.N().S(`" /><button type="submit" name="to" value="module" title="Incorporate change into module">&lt;-</button><button type="submit" name="to" value="project" title="Overwrite project changes with module version">-&gt;</button></form>`)
//line views/vaction/Diff.html:31
		}
//line views/vaction/Diff.html:31
		qw422016.N().S(`</td><td>`)
//line views/vaction/Diff.html:33
		streamrenderPatch(qw422016, d.Patch, as, ps)
//line views/vaction/Diff.html:33
		qw422016.N().S(`</td></tr>`)
//line views/vaction/Diff.html:35
	}
//line views/vaction/Diff.html:35
	qw422016.N().S(`</tbody></table>`)
//line views/vaction/Diff.html:38
}

//line views/vaction/Diff.html:38
func writerenderDiffs(qq422016 qtio422016.Writer, prjKey string, diffs diff.Diffs, as *app.State, ps *cutil.PageState) {
//line views/vaction/Diff.html:38
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vaction/Diff.html:38
	streamrenderDiffs(qw422016, prjKey, diffs, as, ps)
//line views/vaction/Diff.html:38
	qt422016.ReleaseWriter(qw422016)
//line views/vaction/Diff.html:38
}

//line views/vaction/Diff.html:38
func renderDiffs(prjKey string, diffs diff.Diffs, as *app.State, ps *cutil.PageState) string {
//line views/vaction/Diff.html:38
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vaction/Diff.html:38
	writerenderDiffs(qb422016, prjKey, diffs, as, ps)
//line views/vaction/Diff.html:38
	qs422016 := string(qb422016.B)
//line views/vaction/Diff.html:38
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vaction/Diff.html:38
	return qs422016
//line views/vaction/Diff.html:38
}

//line views/vaction/Diff.html:40
func streamrenderPatch(qw422016 *qt422016.Writer, patch string, as *app.State, ps *cutil.PageState) {
//line views/vaction/Diff.html:41
	lines := strings.Split(patch, "\n")

//line views/vaction/Diff.html:41
	qw422016.N().S(`<pre>`)
//line views/vaction/Diff.html:43
	for _, line := range lines {
//line views/vaction/Diff.html:44
		if len(line) > 0 {
//line views/vaction/Diff.html:45
			switch line[0] {
//line views/vaction/Diff.html:46
			case ' ':
//line views/vaction/Diff.html:46
				qw422016.N().S(`<div title="unchanged" style="color: grey;">`)
//line views/vaction/Diff.html:47
				qw422016.E().S(line[1:])
//line views/vaction/Diff.html:47
				if len(line) == 1 {
//line views/vaction/Diff.html:47
					qw422016.N().S(`&nbsp;`)
//line views/vaction/Diff.html:47
				}
//line views/vaction/Diff.html:47
				qw422016.N().S(`</div>`)
//line views/vaction/Diff.html:48
			case '+':
//line views/vaction/Diff.html:48
				qw422016.N().S(`<div title="added" style="color: green;">`)
//line views/vaction/Diff.html:49
				qw422016.E().S(line[1:])
//line views/vaction/Diff.html:49
				if len(line) == 1 {
//line views/vaction/Diff.html:49
					qw422016.N().S(`&nbsp;`)
//line views/vaction/Diff.html:49
				}
//line views/vaction/Diff.html:49
				qw422016.N().S(`</div>`)
//line views/vaction/Diff.html:50
			case '-':
//line views/vaction/Diff.html:50
				qw422016.N().S(`<div title="removed" style="color: red;">`)
//line views/vaction/Diff.html:51
				qw422016.E().S(line[1:])
//line views/vaction/Diff.html:51
				if len(line) == 1 {
//line views/vaction/Diff.html:51
					qw422016.N().S(`&nbsp;`)
//line views/vaction/Diff.html:51
				}
//line views/vaction/Diff.html:51
				qw422016.N().S(`</div>`)
//line views/vaction/Diff.html:52
			default:
//line views/vaction/Diff.html:53
				qw422016.E().S(line)
//line views/vaction/Diff.html:54
			}
//line views/vaction/Diff.html:55
		}
//line views/vaction/Diff.html:56
	}
//line views/vaction/Diff.html:56
	qw422016.N().S(`</pre>`)
//line views/vaction/Diff.html:58
}

//line views/vaction/Diff.html:58
func writerenderPatch(qq422016 qtio422016.Writer, patch string, as *app.State, ps *cutil.PageState) {
//line views/vaction/Diff.html:58
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vaction/Diff.html:58
	streamrenderPatch(qw422016, patch, as, ps)
//line views/vaction/Diff.html:58
	qt422016.ReleaseWriter(qw422016)
//line views/vaction/Diff.html:58
}

//line views/vaction/Diff.html:58
func renderPatch(patch string, as *app.State, ps *cutil.PageState) string {
//line views/vaction/Diff.html:58
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vaction/Diff.html:58
	writerenderPatch(qb422016, patch, as, ps)
//line views/vaction/Diff.html:58
	qs422016 := string(qb422016.B)
//line views/vaction/Diff.html:58
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vaction/Diff.html:58
	return qs422016
//line views/vaction/Diff.html:58
}
