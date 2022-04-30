package checks

import (
	"context"

	"go.uber.org/zap"
	"projectforge.dev/projectforge/app/doctor"
)

var air = &doctor.Check{
	Key:     "air",
	Section: "build",
	Title:   "Air",
	Summary: "Used to recompile the project when files change",
	URL:     "https://github.com/cosmtrek/air",
	UsedBy:  "[bin/dev.sh]",
	Fn:      simpleOut(".", "air", []string{"--help"}, noop),
	Solve:   solveAir,
}

func solveAir(ctx context.Context, r *doctor.Result, logger *zap.SugaredLogger) *doctor.Result {
	if r.Errors.Find("missing") != nil || r.Errors.Find("exitcode") != nil {
		r.AddSolution("go get -u github.com/cosmtrek/air")
	}
	return r
}
