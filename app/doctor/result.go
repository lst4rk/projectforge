package doctor

import (
	"fmt"
	"strings"

	"github.com/kyleu/projectforge/app/util"
)

type Result struct {
	Check    *Check   `json:"-"`
	Key      string   `json:"key"`
	Title    string   `json:"title"`
	Status   string   `json:"status,omitempty"`
	Summary  string   `json:"summary,omitempty"`
	Errors   Errors   `json:"errors,omitempty"`
	Duration int      `json:"duration,omitempty"`
	Logs     []string `json:"logs,omitempty"`
}

func NewResult(check *Check, key string, title string, summary string) *Result {
	return &Result{Check: check, Key: key, Title: title, Status: "OK", Summary: summary}
}

func (p *Result) AddLog(msg string) *Result {
	p.Logs = append(p.Logs, msg)
	return p
}

func (p *Result) WithError(err *Error) *Result {
	p.Status = "error"
	p.Errors = append(p.Errors, err)
	return p
}

func (p *Result) String() string {
	logs := strings.Builder{}
	for _, l := range p.Logs {
		logs.WriteString("\n- ")
		logs.WriteString(l)
	}
	return fmt.Sprintf("%s: %s\n[%s]%s", p.Title, p.Status, p.Summary, logs.String())
}

type Results []*Result

func SimpleOut(path string, cmd string, args []string, outCheck func(r *Result, out string) *Result) func(r *Result) *Result {
	return func(r *Result) *Result {
		fullCmd := strings.Join(append([]string{cmd}, args...), " ")
		exitCode, out, err := util.RunProcessSimple(fullCmd, path)
		if err != nil {
			msg := "[%s] is not present on your computer"
			return r.WithError(NewError("missing", msg, cmd))
		}
		if exitCode != 0 {
			msg := fmt.Sprintf("[%s] returned [%d] as an exit code", cmd, exitCode)
			return r.WithError(NewError("exitcode", msg, "rsvg", fmt.Sprint(exitCode), out))
		}
		return outCheck(r, out)
	}
}
