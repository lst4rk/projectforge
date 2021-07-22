package action

import (
	"github.com/kyleu/projectforge/app/diff"
	"github.com/kyleu/projectforge/app/module"
	"github.com/kyleu/projectforge/app/project"
	"github.com/kyleu/projectforge/app/util"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func onSlam(prj *project.Project, mods module.Modules, cfg util.ValueMap, mSvc *module.Service, pSvc *project.Service, logger *zap.SugaredLogger) *Result {
	ret := newResult(cfg, logger)
	start := util.TimerStart()
	srcFiles, diffs, err := diffs(prj, mods, true, mSvc, pSvc, logger)
	if err != nil {
		return ret.WithError(err)
	}

	for _, f := range diffs {
		switch f.Status {
		case diff.StatusIdentical, diff.StatusMissing, diff.StatusSkipped:
			// noop
		case diff.StatusDifferent, diff.StatusNew:
			src := srcFiles.Get(f.Path)
			tgtFS := pSvc.GetFilesystem(prj)
			err := tgtFS.WriteFile(f.Path, []byte(src.Content), src.Mode, true)
			if err != nil {
				return ret.WithError(errors.Wrapf(err, "unable to write updated content to [%s]", f.Path))
			}
		default:
			return ret.WithError(errors.Errorf("unhandled diff status [%s]", f.Status))
		}
	}

	mr := &module.Result{Keys: mods.Keys(), Status: "OK", Diffs: diffs, Duration: util.TimerEnd(start)}
	ret.Modules = append(ret.Modules, mr)
	return ret
}
