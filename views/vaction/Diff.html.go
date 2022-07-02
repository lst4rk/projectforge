// Code generated by qtc from "Diff.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vaction/Diff.html:1
package vaction

//line views/vaction/Diff.html:1
import (
	"strings"

	"projectforge.dev/projectforge/app"
	"projectforge.dev/projectforge/app/controller/cutil"
	"projectforge.dev/projectforge/app/file/diff"
	"projectforge.dev/projectforge/app/project/action"
	"projectforge.dev/projectforge/app/util"
)

//line views/vaction/Diff.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vaction/Diff.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vaction/Diff.html:11
func streamrenderDiffs(qw422016 *qt422016.Writer, prjKey string, act action.Type, diffs diff.Diffs, cfg util.ValueMap, as *app.State, ps *cutil.PageState) {
//line views/vaction/Diff.html:11
	qw422016.N().S(`<h4 class="mt">Diffs</h4>`)
//line views/vaction/Diff.html:13
	switch act.Key {
//line views/vaction/Diff.html:14
	case "audit":
//line views/vaction/Diff.html:14
		qw422016.N().S(`<div class="mt"><a href="/run/`)
//line views/vaction/Diff.html:16
		qw422016.E().S(prjKey)
//line views/vaction/Diff.html:16
		qw422016.N().S(`/audit?fix=remove" title="Remove all audited files from project"><button>Purge All</button></a>`)
//line views/vaction/Diff.html:17
		if diffs.HasStatus(diff.StatusDifferent) {
//line views/vaction/Diff.html:17
			qw422016.N().S(` `)
//line views/vaction/Diff.html:17
			qw422016.N().S(`<a href="/run/`)
//line views/vaction/Diff.html:18
			qw422016.E().S(prjKey)
//line views/vaction/Diff.html:18
			qw422016.N().S(`/audit?fix=header" title="Remove header from all audited files"><button>Fix All</button></a>`)
//line views/vaction/Diff.html:19
		}
//line views/vaction/Diff.html:19
		qw422016.N().S(`</div>`)
//line views/vaction/Diff.html:21
	case "build":
//line views/vaction/Diff.html:21
		qw422016.N().S(`<div class="mt"><a href="/run/`)
//line views/vaction/Diff.html:23
		qw422016.E().S(prjKey)
//line views/vaction/Diff.html:23
		qw422016.N().S(`/build?phase=imports&fix=true" title="Fix all imports for incorrect files"><button>Fix All</button></a></div>`)
//line views/vaction/Diff.html:25
	}
//line views/vaction/Diff.html:25
	qw422016.N().S(`<table><thead><tr><th class="shrink">Path</th><th class="shrink">Status</th><th class="shrink">Actions</th><th class="shrink">Patch</th></tr></thead><tbody>`)
//line views/vaction/Diff.html:36
	for _, d := range diffs {
//line views/vaction/Diff.html:36
		qw422016.N().S(`<tr><td class="shrink">`)
//line views/vaction/Diff.html:39
		if strings.Contains(d.Path, "..") || d.Status.Key == diff.StatusNew.Key {
//line views/vaction/Diff.html:40
			qw422016.E().S(d.Path)
//line views/vaction/Diff.html:41
		} else {
//line views/vaction/Diff.html:41
			qw422016.N().S(`<a href="/p/`)
//line views/vaction/Diff.html:42
			qw422016.E().S(prjKey)
//line views/vaction/Diff.html:42
			qw422016.N().S(`/fs/`)
//line views/vaction/Diff.html:42
			qw422016.E().S(d.Path)
//line views/vaction/Diff.html:42
			qw422016.N().S(`" target="_blank">`)
//line views/vaction/Diff.html:42
			qw422016.E().S(d.Path)
//line views/vaction/Diff.html:42
			qw422016.N().S(`</a>`)
//line views/vaction/Diff.html:43
		}
//line views/vaction/Diff.html:43
		qw422016.N().S(`</td><td>`)
//line views/vaction/Diff.html:45
		qw422016.E().S(d.Status.StringFor(act.Key))
//line views/vaction/Diff.html:45
		qw422016.N().S(`</td><td>`)
//line views/vaction/Diff.html:47
		switch act.Key {
//line views/vaction/Diff.html:48
		case "audit":
//line views/vaction/Diff.html:48
			qw422016.N().S(`<form action="/run/`)
//line views/vaction/Diff.html:49
			qw422016.E().S(prjKey)
//line views/vaction/Diff.html:49
			qw422016.N().S(`/audit" method="get"><input type="hidden" name="file" value="`)
//line views/vaction/Diff.html:50
			qw422016.E().S(d.Path)
//line views/vaction/Diff.html:50
			qw422016.N().S(`" /><button type="submit" name="fix" value="remove" title="Remove file from project">Purge</button>`)
//line views/vaction/Diff.html:51
			qw422016.N().S(` `)
//line views/vaction/Diff.html:52
			if d.Status.Key == diff.StatusDifferent.Key {
//line views/vaction/Diff.html:52
				qw422016.N().S(`<button type="submit" name="fix" value="header" title="Remove header from file">Fix</button>`)
//line views/vaction/Diff.html:54
			}
//line views/vaction/Diff.html:54
			qw422016.N().S(`</form>`)
//line views/vaction/Diff.html:56
		case "preview":
//line views/vaction/Diff.html:56
			qw422016.N().S(`<form action="/run/`)
//line views/vaction/Diff.html:57
			qw422016.E().S(prjKey)
//line views/vaction/Diff.html:57
			qw422016.N().S(`/generate" method="get"><input type="hidden" name="file" value="`)
//line views/vaction/Diff.html:58
			qw422016.E().S(d.Path)
//line views/vaction/Diff.html:58
			qw422016.N().S(`" />`)
//line views/vaction/Diff.html:59
			if d.Status.String() != "new" {
//line views/vaction/Diff.html:59
				qw422016.N().S(`<button type="submit" name="to" value="module" title="Incorporate change into module">&lt;-</button>`)
//line views/vaction/Diff.html:60
				qw422016.N().S(` `)
//line views/vaction/Diff.html:61
			}
//line views/vaction/Diff.html:61
			qw422016.N().S(`<button type="submit" name="to" value="project" title="Overwrite project changes with module version">-&gt;</button></form>`)
//line views/vaction/Diff.html:64
		case "build":
//line views/vaction/Diff.html:65
			switch cfg.GetStringOpt("phase") {
//line views/vaction/Diff.html:66
			case "imports":
//line views/vaction/Diff.html:66
				qw422016.N().S(`<form action="/run/`)
//line views/vaction/Diff.html:67
				qw422016.E().S(prjKey)
//line views/vaction/Diff.html:67
				qw422016.N().S(`/build" method="get"><input type="hidden" name="fix" value="true" /><input type="hidden" name="phase" value="imports" /><input type="hidden" name="file" value="`)
//line views/vaction/Diff.html:70
				qw422016.E().S(d.Path)
//line views/vaction/Diff.html:70
				qw422016.N().S(`" /><button type="submit" title="Reorder modules">Fix</button></form>`)
//line views/vaction/Diff.html:73
			case "deployments":
//line views/vaction/Diff.html:74
				if d.Status != diff.StatusIdentical {
//line views/vaction/Diff.html:74
					qw422016.N().S(`<form action="/run/`)
//line views/vaction/Diff.html:75
					qw422016.E().S(prjKey)
//line views/vaction/Diff.html:75
					qw422016.N().S(`/build" method="get"><input type="hidden" name="fix" value="true" /><input type="hidden" name="phase" value="deployments" /><input type="hidden" name="file" value="`)
//line views/vaction/Diff.html:78
					qw422016.E().S(d.Path)
//line views/vaction/Diff.html:78
					qw422016.N().S(`" /><button type="submit" title="Update deployment">Fix</button></form>`)
//line views/vaction/Diff.html:81
				}
//line views/vaction/Diff.html:82
			}
//line views/vaction/Diff.html:83
		}
//line views/vaction/Diff.html:83
		qw422016.N().S(`</td><td>`)
//line views/vaction/Diff.html:85
		streamrenderPatch(qw422016, d.Patch, as, ps)
//line views/vaction/Diff.html:85
		qw422016.N().S(`</td></tr>`)
//line views/vaction/Diff.html:87
	}
//line views/vaction/Diff.html:87
	qw422016.N().S(`</tbody></table>`)
//line views/vaction/Diff.html:90
}

//line views/vaction/Diff.html:90
func writerenderDiffs(qq422016 qtio422016.Writer, prjKey string, act action.Type, diffs diff.Diffs, cfg util.ValueMap, as *app.State, ps *cutil.PageState) {
//line views/vaction/Diff.html:90
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vaction/Diff.html:90
	streamrenderDiffs(qw422016, prjKey, act, diffs, cfg, as, ps)
//line views/vaction/Diff.html:90
	qt422016.ReleaseWriter(qw422016)
//line views/vaction/Diff.html:90
}

//line views/vaction/Diff.html:90
func renderDiffs(prjKey string, act action.Type, diffs diff.Diffs, cfg util.ValueMap, as *app.State, ps *cutil.PageState) string {
//line views/vaction/Diff.html:90
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vaction/Diff.html:90
	writerenderDiffs(qb422016, prjKey, act, diffs, cfg, as, ps)
//line views/vaction/Diff.html:90
	qs422016 := string(qb422016.B)
//line views/vaction/Diff.html:90
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vaction/Diff.html:90
	return qs422016
//line views/vaction/Diff.html:90
}

//line views/vaction/Diff.html:92
func streamrenderPatch(qw422016 *qt422016.Writer, patch string, as *app.State, ps *cutil.PageState) {
//line views/vaction/Diff.html:93
	lines := strings.Split(patch, "\n")

//line views/vaction/Diff.html:93
	qw422016.N().S(`<pre>`)
//line views/vaction/Diff.html:95
	for _, line := range lines {
//line views/vaction/Diff.html:96
		if len(line) > 0 {
//line views/vaction/Diff.html:97
			switch line[0] {
//line views/vaction/Diff.html:98
			case ' ':
//line views/vaction/Diff.html:98
				qw422016.N().S(`<div title="unchanged" style="color: grey;">`)
//line views/vaction/Diff.html:99
				qw422016.E().S(line[1:])
//line views/vaction/Diff.html:99
				if len(line) == 1 {
//line views/vaction/Diff.html:99
					qw422016.N().S(`&nbsp;`)
//line views/vaction/Diff.html:99
				}
//line views/vaction/Diff.html:99
				qw422016.N().S(`</div>`)
//line views/vaction/Diff.html:100
			case '+':
//line views/vaction/Diff.html:100
				qw422016.N().S(`<div title="added" class="success">`)
//line views/vaction/Diff.html:101
				qw422016.E().S(line[1:])
//line views/vaction/Diff.html:101
				if len(line) == 1 {
//line views/vaction/Diff.html:101
					qw422016.N().S(`&nbsp;`)
//line views/vaction/Diff.html:101
				}
//line views/vaction/Diff.html:101
				qw422016.N().S(`</div>`)
//line views/vaction/Diff.html:102
			case '-':
//line views/vaction/Diff.html:102
				qw422016.N().S(`<div title="removed" class="error">`)
//line views/vaction/Diff.html:103
				qw422016.E().S(line[1:])
//line views/vaction/Diff.html:103
				if len(line) == 1 {
//line views/vaction/Diff.html:103
					qw422016.N().S(`&nbsp;`)
//line views/vaction/Diff.html:103
				}
//line views/vaction/Diff.html:103
				qw422016.N().S(`</div>`)
//line views/vaction/Diff.html:104
			default:
//line views/vaction/Diff.html:105
				qw422016.E().S(line)
//line views/vaction/Diff.html:106
			}
//line views/vaction/Diff.html:107
		}
//line views/vaction/Diff.html:108
	}
//line views/vaction/Diff.html:108
	qw422016.N().S(`</pre>`)
//line views/vaction/Diff.html:110
}

//line views/vaction/Diff.html:110
func writerenderPatch(qq422016 qtio422016.Writer, patch string, as *app.State, ps *cutil.PageState) {
//line views/vaction/Diff.html:110
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vaction/Diff.html:110
	streamrenderPatch(qw422016, patch, as, ps)
//line views/vaction/Diff.html:110
	qt422016.ReleaseWriter(qw422016)
//line views/vaction/Diff.html:110
}

//line views/vaction/Diff.html:110
func renderPatch(patch string, as *app.State, ps *cutil.PageState) string {
//line views/vaction/Diff.html:110
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vaction/Diff.html:110
	writerenderPatch(qb422016, patch, as, ps)
//line views/vaction/Diff.html:110
	qs422016 := string(qb422016.B)
//line views/vaction/Diff.html:110
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vaction/Diff.html:110
	return qs422016
//line views/vaction/Diff.html:110
}
