package action

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/kyleu/projectforge/app/filesystem"
	"github.com/kyleu/projectforge/app/project"
)

func onCreate(ctx context.Context, params *Params) *Result {
	ret := newResult(params.Cfg, params.Logger)

	path := params.Cfg.GetStringOpt("path")
	if path == "" {
		path = "."
	}
	if wipe, _ := params.Cfg.GetBool("wipe"); wipe {
		fs := filesystem.NewFileSystem(".", params.Logger)
		if fs.Exists(path) {
			ret.AddLog("removing existing directory [%s]", path)
			_ = fs.RemoveRecursive(path)
		}
	}

	prj := projectFromCfg(project.NewProject(params.ProjectKey, path), params.Cfg)

	if params.CLI {
		err := cliProject(prj, params.MSvc.Keys())
		if err != nil {
			return ret.WithError(err)
		}
	}

	params.ProjectKey = prj.Key

	err := params.PSvc.Save(prj)
	if err != nil {
		return ret.WithError(err)
	}

	_, err = params.PSvc.Refresh(params.RootFiles)
	if err != nil {
		msg := fmt.Sprintf("unable to load newly created project from path [%s]", path)
		return errorResult(errors.Wrap(err, msg), params.Cfg, params.Logger)
	}

	pm, err := getPrjAndMods(ctx, params)
	if err != nil {
		return errorResult(err, params.Cfg, params.Logger)
	}
	retS := onSlam(pm)
	ret = ret.Merge(retS)
	if ret.HasErrors() {
		return ret
	}

	return fullBuild(prj, ret, params.Logger)
}

