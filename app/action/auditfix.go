package action

import (
	"context"
	"strings"

	"github.com/pkg/errors"

	"projectforge.dev/projectforge/app/file"
)

func auditRemove(ctx context.Context, fn string, pm *PrjAndMods, ret *Result) error {
	if fn == "" {
		paths := ret.Modules.Paths(false)
		if len(paths) == 0 {
			x := &Result{}
			err := auditRun(ctx, pm, x)
			if err != nil {
				return err
			}
			paths = x.Modules.Paths(false)
		}
		for _, path := range paths {
			err := auditRemove(ctx, path, pm, ret)
			if err != nil {
				return errors.Wrapf(err, "can't fix audit of [%s]", path)
			}
		}
		return nil
	}
	tgt := pm.PSvc.GetFilesystem(pm.Prj)
	ret.AddLog("removed [%s]", fn)
	return tgt.Remove(fn, pm.Logger)
}

func auditHeader(ctx context.Context, fn string, pm *PrjAndMods, ret *Result) error {
	if fn == "" {
		x := &Result{}
		err := auditRun(ctx, pm, x)
		if err != nil {
			return err
		}
		for _, path := range x.Modules.Paths(false) {
			err = auditHeader(ctx, path, pm, ret)
			if err != nil {
				return errors.Wrapf(err, "can't fix audit of [%s]", path)
			}
		}
		return nil
	}
	tgt := pm.PSvc.GetFilesystem(pm.Prj)
	stat, err := tgt.Stat(fn)
	if err != nil {
		return err
	}
	content, err := tgt.ReadFile(fn)
	if err != nil {
		return err
	}
	lines := strings.Split(string(content), "\n")
	var hit bool
	fixed := make([]string, 0, len(lines))
	for _, line := range lines {
		if strings.Contains(line, file.HeaderContent) {
			hit = true
		} else {
			fixed = append(fixed, line)
		}
	}
	if hit {
		final := strings.Join(fixed, "\n")
		err = tgt.WriteFile(fn, []byte(final), stat.Mode(), true)
		if err != nil {
			return errors.Wrapf(err, "can't overwrite file at path [%s]", fn)
		}
		ret.AddLog("removed header from file [%s]", fn)
	} else {
		ret.AddLog("no header to remove from file [%s]", fn)
	}
	return nil
}
