// Code generated by qtc from "FormVertical.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/components/FormVertical.html:2
package components

//line views/components/FormVertical.html:2
import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"projectforge.dev/projectforge/app/controller/cutil"
	"projectforge.dev/projectforge/views/vutil"
)

//line views/components/FormVertical.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/FormVertical.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/FormVertical.html:12
func StreamFormVerticalInput(qw422016 *qt422016.Writer, key string, id string, title string, value string, indent int, help ...string) {
//line views/components/FormVertical.html:13
	id = cutil.CleanID(key, id)

//line views/components/FormVertical.html:13
	qw422016.N().S(`<div class="mb expanded">`)
//line views/components/FormVertical.html:15
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/FormVertical.html:15
	qw422016.N().S(`<label for="`)
//line views/components/FormVertical.html:16
	qw422016.E().S(id)
//line views/components/FormVertical.html:16
	qw422016.N().S(`"><em class="title">`)
//line views/components/FormVertical.html:16
	qw422016.E().S(title)
//line views/components/FormVertical.html:16
	qw422016.N().S(`</em></label>`)
//line views/components/FormVertical.html:17
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/FormVertical.html:17
	qw422016.N().S(`<div class="mt">`)
//line views/components/FormVertical.html:18
	StreamFormInput(qw422016, key, id, value, help...)
//line views/components/FormVertical.html:18
	qw422016.N().S(`</div>`)
//line views/components/FormVertical.html:19
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/FormVertical.html:19
	qw422016.N().S(`</div>`)
//line views/components/FormVertical.html:21
}

//line views/components/FormVertical.html:21
func WriteFormVerticalInput(qq422016 qtio422016.Writer, key string, id string, title string, value string, indent int, help ...string) {
//line views/components/FormVertical.html:21
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/FormVertical.html:21
	StreamFormVerticalInput(qw422016, key, id, title, value, indent, help...)
//line views/components/FormVertical.html:21
	qt422016.ReleaseWriter(qw422016)
//line views/components/FormVertical.html:21
}

//line views/components/FormVertical.html:21
func FormVerticalInput(key string, id string, title string, value string, indent int, help ...string) string {
//line views/components/FormVertical.html:21
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/FormVertical.html:21
	WriteFormVerticalInput(qb422016, key, id, title, value, indent, help...)
//line views/components/FormVertical.html:21
	qs422016 := string(qb422016.B)
//line views/components/FormVertical.html:21
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/FormVertical.html:21
	return qs422016
//line views/components/FormVertical.html:21
}

//line views/components/FormVertical.html:23
func StreamFormVerticalInputPassword(qw422016 *qt422016.Writer, key string, id string, title string, value string, indent int, help ...string) {
//line views/components/FormVertical.html:24
	id = cutil.CleanID(key, id)

//line views/components/FormVertical.html:24
	qw422016.N().S(`<div class="mb expanded">`)
//line views/components/FormVertical.html:26
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/FormVertical.html:26
	qw422016.N().S(`<label for="`)
//line views/components/FormVertical.html:27
	qw422016.E().S(id)
//line views/components/FormVertical.html:27
	qw422016.N().S(`"><em class="title">`)
//line views/components/FormVertical.html:27
	qw422016.E().S(title)
//line views/components/FormVertical.html:27
	qw422016.N().S(`</em></label>`)
//line views/components/FormVertical.html:28
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/FormVertical.html:28
	qw422016.N().S(`<div class="mt">`)
//line views/components/FormVertical.html:29
	StreamFormInputPassword(qw422016, key, id, value, help...)
//line views/components/FormVertical.html:29
	qw422016.N().S(`</div>`)
//line views/components/FormVertical.html:30
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/FormVertical.html:30
	qw422016.N().S(`</div>`)
//line views/components/FormVertical.html:32
}

//line views/components/FormVertical.html:32
func WriteFormVerticalInputPassword(qq422016 qtio422016.Writer, key string, id string, title string, value string, indent int, help ...string) {
//line views/components/FormVertical.html:32
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/FormVertical.html:32
	StreamFormVerticalInputPassword(qw422016, key, id, title, value, indent, help...)
//line views/components/FormVertical.html:32
	qt422016.ReleaseWriter(qw422016)
//line views/components/FormVertical.html:32
}

//line views/components/FormVertical.html:32
func FormVerticalInputPassword(key string, id string, title string, value string, indent int, help ...string) string {
//line views/components/FormVertical.html:32
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/FormVertical.html:32
	WriteFormVerticalInputPassword(qb422016, key, id, title, value, indent, help...)
//line views/components/FormVertical.html:32
	qs422016 := string(qb422016.B)
//line views/components/FormVertical.html:32
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/FormVertical.html:32
	return qs422016
//line views/components/FormVertical.html:32
}

//line views/components/FormVertical.html:34
func StreamFormVerticalInputNumber(qw422016 *qt422016.Writer, key string, id string, title string, value int, indent int, help ...string) {
//line views/components/FormVertical.html:35
	id = cutil.CleanID(key, id)

//line views/components/FormVertical.html:35
	qw422016.N().S(`<div class="mb expanded">`)
//line views/components/FormVertical.html:37
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/FormVertical.html:37
	qw422016.N().S(`<label for="`)
//line views/components/FormVertical.html:38
	qw422016.E().S(id)
//line views/components/FormVertical.html:38
	qw422016.N().S(`"><em class="title">`)
//line views/components/FormVertical.html:38
	qw422016.E().S(title)
//line views/components/FormVertical.html:38
	qw422016.N().S(`</em></label>`)
//line views/components/FormVertical.html:39
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/FormVertical.html:39
	qw422016.N().S(`<div class="mt">`)
//line views/components/FormVertical.html:40
	StreamFormInputNumber(qw422016, key, id, value, help...)
//line views/components/FormVertical.html:40
	qw422016.N().S(`</div>`)
//line views/components/FormVertical.html:41
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/FormVertical.html:41
	qw422016.N().S(`</div>`)
//line views/components/FormVertical.html:43
}

//line views/components/FormVertical.html:43
func WriteFormVerticalInputNumber(qq422016 qtio422016.Writer, key string, id string, title string, value int, indent int, help ...string) {
//line views/components/FormVertical.html:43
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/FormVertical.html:43
	StreamFormVerticalInputNumber(qw422016, key, id, title, value, indent, help...)
//line views/components/FormVertical.html:43
	qt422016.ReleaseWriter(qw422016)
//line views/components/FormVertical.html:43
}

//line views/components/FormVertical.html:43
func FormVerticalInputNumber(key string, id string, title string, value int, indent int, help ...string) string {
//line views/components/FormVertical.html:43
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/FormVertical.html:43
	WriteFormVerticalInputNumber(qb422016, key, id, title, value, indent, help...)
//line views/components/FormVertical.html:43
	qs422016 := string(qb422016.B)
//line views/components/FormVertical.html:43
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/FormVertical.html:43
	return qs422016
//line views/components/FormVertical.html:43
}

//line views/components/FormVertical.html:45
func StreamFormVerticalInputFloat(qw422016 *qt422016.Writer, key string, id string, title string, value float64, indent int, help ...string) {
//line views/components/FormVertical.html:46
	id = cutil.CleanID(key, id)

//line views/components/FormVertical.html:46
	qw422016.N().S(`<div class="mb expanded">`)
//line views/components/FormVertical.html:48
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/FormVertical.html:48
	qw422016.N().S(`<label for="`)
//line views/components/FormVertical.html:49
	qw422016.E().S(id)
//line views/components/FormVertical.html:49
	qw422016.N().S(`"><em class="title">`)
//line views/components/FormVertical.html:49
	qw422016.E().S(title)
//line views/components/FormVertical.html:49
	qw422016.N().S(`</em></label>`)
//line views/components/FormVertical.html:50
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/FormVertical.html:50
	qw422016.N().S(`<div class="mt">`)
//line views/components/FormVertical.html:51
	StreamFormInputFloat(qw422016, key, id, value, help...)
//line views/components/FormVertical.html:51
	qw422016.N().S(`</div>`)
//line views/components/FormVertical.html:52
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/FormVertical.html:52
	qw422016.N().S(`</div>`)
//line views/components/FormVertical.html:54
}

//line views/components/FormVertical.html:54
func WriteFormVerticalInputFloat(qq422016 qtio422016.Writer, key string, id string, title string, value float64, indent int, help ...string) {
//line views/components/FormVertical.html:54
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/FormVertical.html:54
	StreamFormVerticalInputFloat(qw422016, key, id, title, value, indent, help...)
//line views/components/FormVertical.html:54
	qt422016.ReleaseWriter(qw422016)
//line views/components/FormVertical.html:54
}

//line views/components/FormVertical.html:54
func FormVerticalInputFloat(key string, id string, title string, value float64, indent int, help ...string) string {
//line views/components/FormVertical.html:54
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/FormVertical.html:54
	WriteFormVerticalInputFloat(qb422016, key, id, title, value, indent, help...)
//line views/components/FormVertical.html:54
	qs422016 := string(qb422016.B)
//line views/components/FormVertical.html:54
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/FormVertical.html:54
	return qs422016
//line views/components/FormVertical.html:54
}

//line views/components/FormVertical.html:56
func StreamFormVerticalInputTimestamp(qw422016 *qt422016.Writer, key string, id string, title string, value *time.Time, indent int, help ...string) {
//line views/components/FormVertical.html:57
	id = cutil.CleanID(key, id)

//line views/components/FormVertical.html:57
	qw422016.N().S(`<div class="mb expanded">`)
//line views/components/FormVertical.html:59
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/FormVertical.html:59
	qw422016.N().S(`<label for="`)
//line views/components/FormVertical.html:60
	qw422016.E().S(id)
//line views/components/FormVertical.html:60
	qw422016.N().S(`"><em class="title">`)
//line views/components/FormVertical.html:60
	qw422016.E().S(title)
//line views/components/FormVertical.html:60
	qw422016.N().S(`</em></label>`)
//line views/components/FormVertical.html:61
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/FormVertical.html:61
	qw422016.N().S(`<div class="mt">`)
//line views/components/FormVertical.html:62
	StreamFormInputTimestamp(qw422016, key, id, value, help...)
//line views/components/FormVertical.html:62
	qw422016.N().S(`</div>`)
//line views/components/FormVertical.html:63
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/FormVertical.html:63
	qw422016.N().S(`</div>`)
//line views/components/FormVertical.html:65
}

//line views/components/FormVertical.html:65
func WriteFormVerticalInputTimestamp(qq422016 qtio422016.Writer, key string, id string, title string, value *time.Time, indent int, help ...string) {
//line views/components/FormVertical.html:65
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/FormVertical.html:65
	StreamFormVerticalInputTimestamp(qw422016, key, id, title, value, indent, help...)
//line views/components/FormVertical.html:65
	qt422016.ReleaseWriter(qw422016)
//line views/components/FormVertical.html:65
}

//line views/components/FormVertical.html:65
func FormVerticalInputTimestamp(key string, id string, title string, value *time.Time, indent int, help ...string) string {
//line views/components/FormVertical.html:65
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/FormVertical.html:65
	WriteFormVerticalInputTimestamp(qb422016, key, id, title, value, indent, help...)
//line views/components/FormVertical.html:65
	qs422016 := string(qb422016.B)
//line views/components/FormVertical.html:65
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/FormVertical.html:65
	return qs422016
//line views/components/FormVertical.html:65
}

//line views/components/FormVertical.html:67
func StreamFormVerticalInputTimestampDay(qw422016 *qt422016.Writer, key string, id string, title string, value *time.Time, indent int, help ...string) {
//line views/components/FormVertical.html:68
	id = cutil.CleanID(key, id)

//line views/components/FormVertical.html:68
	qw422016.N().S(`<div class="mb expanded">`)
//line views/components/FormVertical.html:70
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/FormVertical.html:70
	qw422016.N().S(`<label for="`)
//line views/components/FormVertical.html:71
	qw422016.E().S(id)
//line views/components/FormVertical.html:71
	qw422016.N().S(`"><em class="title">`)
//line views/components/FormVertical.html:71
	qw422016.E().S(title)
//line views/components/FormVertical.html:71
	qw422016.N().S(`</em></label>`)
//line views/components/FormVertical.html:72
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/FormVertical.html:72
	qw422016.N().S(`<div class="mt">`)
//line views/components/FormVertical.html:73
	StreamFormInputTimestampDay(qw422016, key, id, value, help...)
//line views/components/FormVertical.html:73
	qw422016.N().S(`</div>`)
//line views/components/FormVertical.html:74
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/FormVertical.html:74
	qw422016.N().S(`</div>`)
//line views/components/FormVertical.html:76
}

//line views/components/FormVertical.html:76
func WriteFormVerticalInputTimestampDay(qq422016 qtio422016.Writer, key string, id string, title string, value *time.Time, indent int, help ...string) {
//line views/components/FormVertical.html:76
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/FormVertical.html:76
	StreamFormVerticalInputTimestampDay(qw422016, key, id, title, value, indent, help...)
//line views/components/FormVertical.html:76
	qt422016.ReleaseWriter(qw422016)
//line views/components/FormVertical.html:76
}

//line views/components/FormVertical.html:76
func FormVerticalInputTimestampDay(key string, id string, title string, value *time.Time, indent int, help ...string) string {
//line views/components/FormVertical.html:76
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/FormVertical.html:76
	WriteFormVerticalInputTimestampDay(qb422016, key, id, title, value, indent, help...)
//line views/components/FormVertical.html:76
	qs422016 := string(qb422016.B)
//line views/components/FormVertical.html:76
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/FormVertical.html:76
	return qs422016
//line views/components/FormVertical.html:76
}

//line views/components/FormVertical.html:78
func StreamFormVerticalInputUUID(qw422016 *qt422016.Writer, key string, id string, title string, value *uuid.UUID, indent int, help ...string) {
//line views/components/FormVertical.html:79
	id = cutil.CleanID(key, id)

//line views/components/FormVertical.html:79
	qw422016.N().S(`<div class="mb expanded">`)
//line views/components/FormVertical.html:81
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/FormVertical.html:81
	qw422016.N().S(`<label for="`)
//line views/components/FormVertical.html:82
	qw422016.E().S(id)
//line views/components/FormVertical.html:82
	qw422016.N().S(`"><em class="title">`)
//line views/components/FormVertical.html:82
	qw422016.E().S(title)
//line views/components/FormVertical.html:82
	qw422016.N().S(`</em></label>`)
//line views/components/FormVertical.html:83
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/FormVertical.html:83
	qw422016.N().S(`<div class="mt">`)
//line views/components/FormVertical.html:84
	StreamFormInputUUID(qw422016, key, id, value, help...)
//line views/components/FormVertical.html:84
	qw422016.N().S(`</div>`)
//line views/components/FormVertical.html:85
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/FormVertical.html:85
	qw422016.N().S(`</div>`)
//line views/components/FormVertical.html:87
}

//line views/components/FormVertical.html:87
func WriteFormVerticalInputUUID(qq422016 qtio422016.Writer, key string, id string, title string, value *uuid.UUID, indent int, help ...string) {
//line views/components/FormVertical.html:87
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/FormVertical.html:87
	StreamFormVerticalInputUUID(qw422016, key, id, title, value, indent, help...)
//line views/components/FormVertical.html:87
	qt422016.ReleaseWriter(qw422016)
//line views/components/FormVertical.html:87
}

//line views/components/FormVertical.html:87
func FormVerticalInputUUID(key string, id string, title string, value *uuid.UUID, indent int, help ...string) string {
//line views/components/FormVertical.html:87
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/FormVertical.html:87
	WriteFormVerticalInputUUID(qb422016, key, id, title, value, indent, help...)
//line views/components/FormVertical.html:87
	qs422016 := string(qb422016.B)
//line views/components/FormVertical.html:87
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/FormVertical.html:87
	return qs422016
//line views/components/FormVertical.html:87
}

//line views/components/FormVertical.html:89
func StreamFormVerticalTextarea(qw422016 *qt422016.Writer, key string, id string, title string, rows int, value string, indent int, help ...string) {
//line views/components/FormVertical.html:90
	id = cutil.CleanID(key, id)

//line views/components/FormVertical.html:90
	qw422016.N().S(`<div class="mb expanded">`)
//line views/components/FormVertical.html:92
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/FormVertical.html:92
	qw422016.N().S(`<label for="`)
//line views/components/FormVertical.html:93
	qw422016.E().S(id)
//line views/components/FormVertical.html:93
	qw422016.N().S(`"><em class="title">`)
//line views/components/FormVertical.html:93
	qw422016.E().S(title)
//line views/components/FormVertical.html:93
	qw422016.N().S(`</em></label>`)
//line views/components/FormVertical.html:94
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/FormVertical.html:94
	qw422016.N().S(`<div class="mt">`)
//line views/components/FormVertical.html:95
	StreamFormTextarea(qw422016, key, id, rows, value, help...)
//line views/components/FormVertical.html:95
	qw422016.N().S(`</div>`)
//line views/components/FormVertical.html:96
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/FormVertical.html:96
	qw422016.N().S(`</div>`)
//line views/components/FormVertical.html:98
}

//line views/components/FormVertical.html:98
func WriteFormVerticalTextarea(qq422016 qtio422016.Writer, key string, id string, title string, rows int, value string, indent int, help ...string) {
//line views/components/FormVertical.html:98
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/FormVertical.html:98
	StreamFormVerticalTextarea(qw422016, key, id, title, rows, value, indent, help...)
//line views/components/FormVertical.html:98
	qt422016.ReleaseWriter(qw422016)
//line views/components/FormVertical.html:98
}

//line views/components/FormVertical.html:98
func FormVerticalTextarea(key string, id string, title string, rows int, value string, indent int, help ...string) string {
//line views/components/FormVertical.html:98
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/FormVertical.html:98
	WriteFormVerticalTextarea(qb422016, key, id, title, rows, value, indent, help...)
//line views/components/FormVertical.html:98
	qs422016 := string(qb422016.B)
//line views/components/FormVertical.html:98
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/FormVertical.html:98
	return qs422016
//line views/components/FormVertical.html:98
}

//line views/components/FormVertical.html:100
func StreamFormVerticalSelect(qw422016 *qt422016.Writer, key string, id string, title string, value string, opts []string, titles []string, indent int, help ...string) {
//line views/components/FormVertical.html:101
	id = cutil.CleanID(key, id)

//line views/components/FormVertical.html:101
	qw422016.N().S(`<div class="mb expanded">`)
//line views/components/FormVertical.html:103
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/FormVertical.html:103
	qw422016.N().S(`<label for="`)
//line views/components/FormVertical.html:104
	qw422016.E().S(id)
//line views/components/FormVertical.html:104
	qw422016.N().S(`"><em class="title">`)
//line views/components/FormVertical.html:104
	qw422016.E().S(title)
//line views/components/FormVertical.html:104
	qw422016.N().S(`</em></label>`)
//line views/components/FormVertical.html:105
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/FormVertical.html:105
	qw422016.N().S(`<div class="mt">`)
//line views/components/FormVertical.html:106
	StreamFormSelect(qw422016, key, id, value, opts, titles, indent)
//line views/components/FormVertical.html:106
	qw422016.N().S(`</div>`)
//line views/components/FormVertical.html:107
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/FormVertical.html:107
	qw422016.N().S(`</div>`)
//line views/components/FormVertical.html:109
}

//line views/components/FormVertical.html:109
func WriteFormVerticalSelect(qq422016 qtio422016.Writer, key string, id string, title string, value string, opts []string, titles []string, indent int, help ...string) {
//line views/components/FormVertical.html:109
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/FormVertical.html:109
	StreamFormVerticalSelect(qw422016, key, id, title, value, opts, titles, indent, help...)
//line views/components/FormVertical.html:109
	qt422016.ReleaseWriter(qw422016)
//line views/components/FormVertical.html:109
}

//line views/components/FormVertical.html:109
func FormVerticalSelect(key string, id string, title string, value string, opts []string, titles []string, indent int, help ...string) string {
//line views/components/FormVertical.html:109
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/FormVertical.html:109
	WriteFormVerticalSelect(qb422016, key, id, title, value, opts, titles, indent, help...)
//line views/components/FormVertical.html:109
	qs422016 := string(qb422016.B)
//line views/components/FormVertical.html:109
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/FormVertical.html:109
	return qs422016
//line views/components/FormVertical.html:109
}

//line views/components/FormVertical.html:111
func StreamFormVerticalDatalist(qw422016 *qt422016.Writer, key string, id string, title string, value string, opts []string, titles []string, indent int, help ...string) {
//line views/components/FormVertical.html:112
	id = cutil.CleanID(key, id)

//line views/components/FormVertical.html:112
	qw422016.N().S(`<div class="mb expanded">`)
//line views/components/FormVertical.html:114
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/FormVertical.html:114
	qw422016.N().S(`<label for="`)
//line views/components/FormVertical.html:115
	qw422016.E().S(id)
//line views/components/FormVertical.html:115
	qw422016.N().S(`"><em class="title">`)
//line views/components/FormVertical.html:115
	qw422016.E().S(title)
//line views/components/FormVertical.html:115
	qw422016.N().S(`</em></label>`)
//line views/components/FormVertical.html:116
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/FormVertical.html:116
	qw422016.N().S(`<div class="mt">`)
//line views/components/FormVertical.html:117
	StreamFormDatalist(qw422016, key, id, value, opts, titles, indent)
//line views/components/FormVertical.html:117
	qw422016.N().S(`</div>`)
//line views/components/FormVertical.html:118
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/FormVertical.html:118
	qw422016.N().S(`</div>`)
//line views/components/FormVertical.html:120
}

//line views/components/FormVertical.html:120
func WriteFormVerticalDatalist(qq422016 qtio422016.Writer, key string, id string, title string, value string, opts []string, titles []string, indent int, help ...string) {
//line views/components/FormVertical.html:120
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/FormVertical.html:120
	StreamFormVerticalDatalist(qw422016, key, id, title, value, opts, titles, indent, help...)
//line views/components/FormVertical.html:120
	qt422016.ReleaseWriter(qw422016)
//line views/components/FormVertical.html:120
}

//line views/components/FormVertical.html:120
func FormVerticalDatalist(key string, id string, title string, value string, opts []string, titles []string, indent int, help ...string) string {
//line views/components/FormVertical.html:120
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/FormVertical.html:120
	WriteFormVerticalDatalist(qb422016, key, id, title, value, opts, titles, indent, help...)
//line views/components/FormVertical.html:120
	qs422016 := string(qb422016.B)
//line views/components/FormVertical.html:120
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/FormVertical.html:120
	return qs422016
//line views/components/FormVertical.html:120
}

//line views/components/FormVertical.html:122
func StreamFormVerticalInputTags(qw422016 *qt422016.Writer, key string, id string, title string, values []string, ps *cutil.PageState, indent int, help ...string) {
//line views/components/FormVertical.html:123
	id = cutil.CleanID(key, id)

//line views/components/FormVertical.html:123
	qw422016.N().S(`<div class="mb expanded">`)
//line views/components/FormVertical.html:125
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/FormVertical.html:125
	qw422016.N().S(`<label for="`)
//line views/components/FormVertical.html:126
	qw422016.E().S(id)
//line views/components/FormVertical.html:126
	qw422016.N().S(`"><em class="title">`)
//line views/components/FormVertical.html:126
	qw422016.E().S(title)
//line views/components/FormVertical.html:126
	qw422016.N().S(`</em></label>`)
//line views/components/FormVertical.html:127
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/FormVertical.html:127
	qw422016.N().S(`<div class="mt">`)
//line views/components/FormVertical.html:128
	StreamFormInputTags(qw422016, key, id, values, ps, help...)
//line views/components/FormVertical.html:128
	qw422016.N().S(`</div>`)
//line views/components/FormVertical.html:129
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/FormVertical.html:129
	qw422016.N().S(`</div>`)
//line views/components/FormVertical.html:131
}

//line views/components/FormVertical.html:131
func WriteFormVerticalInputTags(qq422016 qtio422016.Writer, key string, id string, title string, values []string, ps *cutil.PageState, indent int, help ...string) {
//line views/components/FormVertical.html:131
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/FormVertical.html:131
	StreamFormVerticalInputTags(qw422016, key, id, title, values, ps, indent, help...)
//line views/components/FormVertical.html:131
	qt422016.ReleaseWriter(qw422016)
//line views/components/FormVertical.html:131
}

//line views/components/FormVertical.html:131
func FormVerticalInputTags(key string, id string, title string, values []string, ps *cutil.PageState, indent int, help ...string) string {
//line views/components/FormVertical.html:131
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/FormVertical.html:131
	WriteFormVerticalInputTags(qb422016, key, id, title, values, ps, indent, help...)
//line views/components/FormVertical.html:131
	qs422016 := string(qb422016.B)
//line views/components/FormVertical.html:131
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/FormVertical.html:131
	return qs422016
//line views/components/FormVertical.html:131
}

//line views/components/FormVertical.html:133
func StreamFormVerticalRadio(qw422016 *qt422016.Writer, key string, title string, value string, opts []string, titles []string, indent int, help ...string) {
//line views/components/FormVertical.html:133
	qw422016.N().S(`<div class="mb expanded">`)
//line views/components/FormVertical.html:135
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/FormVertical.html:135
	qw422016.N().S(`<div>`)
//line views/components/FormVertical.html:136
	streamtitleFor(qw422016, help)
//line views/components/FormVertical.html:136
	qw422016.N().S(`>`)
//line views/components/FormVertical.html:136
	qw422016.E().S(title)
//line views/components/FormVertical.html:136
	qw422016.N().S(`</div>`)
//line views/components/FormVertical.html:137
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/FormVertical.html:137
	qw422016.N().S(`<div class="mt">`)
//line views/components/FormVertical.html:139
	StreamFormRadio(qw422016, key, value, opts, titles, indent+2)
//line views/components/FormVertical.html:140
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/FormVertical.html:140
	qw422016.N().S(`</div>`)
//line views/components/FormVertical.html:142
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/FormVertical.html:142
	qw422016.N().S(`</div>`)
//line views/components/FormVertical.html:144
}

//line views/components/FormVertical.html:144
func WriteFormVerticalRadio(qq422016 qtio422016.Writer, key string, title string, value string, opts []string, titles []string, indent int, help ...string) {
//line views/components/FormVertical.html:144
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/FormVertical.html:144
	StreamFormVerticalRadio(qw422016, key, title, value, opts, titles, indent, help...)
//line views/components/FormVertical.html:144
	qt422016.ReleaseWriter(qw422016)
//line views/components/FormVertical.html:144
}

//line views/components/FormVertical.html:144
func FormVerticalRadio(key string, title string, value string, opts []string, titles []string, indent int, help ...string) string {
//line views/components/FormVertical.html:144
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/FormVertical.html:144
	WriteFormVerticalRadio(qb422016, key, title, value, opts, titles, indent, help...)
//line views/components/FormVertical.html:144
	qs422016 := string(qb422016.B)
//line views/components/FormVertical.html:144
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/FormVertical.html:144
	return qs422016
//line views/components/FormVertical.html:144
}

//line views/components/FormVertical.html:146
func StreamFormVerticalBoolean(qw422016 *qt422016.Writer, key string, title string, value bool, indent int, help ...string) {
//line views/components/FormVertical.html:147
	StreamFormVerticalRadio(qw422016, key, title, fmt.Sprint(value), []string{"true", "false"}, []string{"True", "False"}, indent)
//line views/components/FormVertical.html:148
}

//line views/components/FormVertical.html:148
func WriteFormVerticalBoolean(qq422016 qtio422016.Writer, key string, title string, value bool, indent int, help ...string) {
//line views/components/FormVertical.html:148
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/FormVertical.html:148
	StreamFormVerticalBoolean(qw422016, key, title, value, indent, help...)
//line views/components/FormVertical.html:148
	qt422016.ReleaseWriter(qw422016)
//line views/components/FormVertical.html:148
}

//line views/components/FormVertical.html:148
func FormVerticalBoolean(key string, title string, value bool, indent int, help ...string) string {
//line views/components/FormVertical.html:148
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/FormVertical.html:148
	WriteFormVerticalBoolean(qb422016, key, title, value, indent, help...)
//line views/components/FormVertical.html:148
	qs422016 := string(qb422016.B)
//line views/components/FormVertical.html:148
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/FormVertical.html:148
	return qs422016
//line views/components/FormVertical.html:148
}

//line views/components/FormVertical.html:150
func StreamFormVerticalCheckbox(qw422016 *qt422016.Writer, key string, title string, values []string, opts []string, titles []string, indent int, help ...string) {
//line views/components/FormVertical.html:150
	qw422016.N().S(`<div class="mb expanded">`)
//line views/components/FormVertical.html:152
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/FormVertical.html:152
	qw422016.N().S(`<div>`)
//line views/components/FormVertical.html:153
	qw422016.E().S(title)
//line views/components/FormVertical.html:153
	qw422016.N().S(`</div>`)
//line views/components/FormVertical.html:154
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/FormVertical.html:154
	qw422016.N().S(`<div class="mt">`)
//line views/components/FormVertical.html:156
	StreamFormCheckbox(qw422016, key, values, opts, titles, indent+2)
//line views/components/FormVertical.html:157
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/FormVertical.html:157
	qw422016.N().S(`</div>`)
//line views/components/FormVertical.html:159
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/FormVertical.html:159
	qw422016.N().S(`</div>`)
//line views/components/FormVertical.html:161
}

//line views/components/FormVertical.html:161
func WriteFormVerticalCheckbox(qq422016 qtio422016.Writer, key string, title string, values []string, opts []string, titles []string, indent int, help ...string) {
//line views/components/FormVertical.html:161
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/FormVertical.html:161
	StreamFormVerticalCheckbox(qw422016, key, title, values, opts, titles, indent, help...)
//line views/components/FormVertical.html:161
	qt422016.ReleaseWriter(qw422016)
//line views/components/FormVertical.html:161
}

//line views/components/FormVertical.html:161
func FormVerticalCheckbox(key string, title string, values []string, opts []string, titles []string, indent int, help ...string) string {
//line views/components/FormVertical.html:161
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/FormVertical.html:161
	WriteFormVerticalCheckbox(qb422016, key, title, values, opts, titles, indent, help...)
//line views/components/FormVertical.html:161
	qs422016 := string(qb422016.B)
//line views/components/FormVertical.html:161
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/FormVertical.html:161
	return qs422016
//line views/components/FormVertical.html:161
}

//line views/components/FormVertical.html:163
func StreamFormVerticalIconPicker(qw422016 *qt422016.Writer, key string, title string, value string, ps *cutil.PageState, indent int) {
//line views/components/FormVertical.html:163
	qw422016.N().S(`<div class="mb expanded">`)
//line views/components/FormVertical.html:165
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/FormVertical.html:165
	qw422016.N().S(`<em class="title">`)
//line views/components/FormVertical.html:166
	qw422016.E().S(title)
//line views/components/FormVertical.html:166
	qw422016.N().S(`</em>`)
//line views/components/FormVertical.html:167
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/FormVertical.html:167
	qw422016.N().S(`<div class="mt">`)
//line views/components/FormVertical.html:168
	StreamIconPicker(qw422016, key, value, ps, indent)
//line views/components/FormVertical.html:168
	qw422016.N().S(`</div>`)
//line views/components/FormVertical.html:169
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/FormVertical.html:169
	qw422016.N().S(`</div>`)
//line views/components/FormVertical.html:171
}

//line views/components/FormVertical.html:171
func WriteFormVerticalIconPicker(qq422016 qtio422016.Writer, key string, title string, value string, ps *cutil.PageState, indent int) {
//line views/components/FormVertical.html:171
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/FormVertical.html:171
	StreamFormVerticalIconPicker(qw422016, key, title, value, ps, indent)
//line views/components/FormVertical.html:171
	qt422016.ReleaseWriter(qw422016)
//line views/components/FormVertical.html:171
}

//line views/components/FormVertical.html:171
func FormVerticalIconPicker(key string, title string, value string, ps *cutil.PageState, indent int) string {
//line views/components/FormVertical.html:171
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/FormVertical.html:171
	WriteFormVerticalIconPicker(qb422016, key, title, value, ps, indent)
//line views/components/FormVertical.html:171
	qs422016 := string(qb422016.B)
//line views/components/FormVertical.html:171
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/FormVertical.html:171
	return qs422016
//line views/components/FormVertical.html:171
}
