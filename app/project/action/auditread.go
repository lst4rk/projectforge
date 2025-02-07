package action

import (
	"context"
	"path"

	"github.com/pkg/errors"
	"golang.org/x/exp/slices"

	"projectforge.dev/projectforge/app/file"
	"projectforge.dev/projectforge/app/lib/filesystem"
	"projectforge.dev/projectforge/app/util"
)

func getGeneratedFiles(tgt filesystem.FileLoader, ignore []string, logger util.Logger) ([]string, error) {
	filenames, err := tgt.ListFilesRecursive("", ignore, logger)
	if err != nil {
		return nil, err
	}

	var ret []string
	for _, fn := range filenames {
		b, e := tgt.PeekFile(fn, 1024)
		if e != nil {
			return nil, e
		}
		if file.ContainsHeader(string(b)) {
			ret = append(ret, fn)
		}
	}
	return ret, nil
}

func getModuleFiles(ctx context.Context, pm *PrjAndMods) ([]string, error) {
	ret, err := pm.MSvc.GetFilenames(pm.Mods, pm.Logger)
	if err != nil {
		return nil, err
	}

	if pm.Mods.Get("export") != nil {
		args, err := pm.Prj.ModuleArgExport(pm.PSvc, pm.Logger)
		if err != nil {
			return nil, err
		}
		args.Modules = pm.Mods.Keys()
		files, e := pm.ESvc.Files(ctx, pm.Prj, args, true, pm.Logger)
		if e != nil {
			return nil, err
		}
		for _, f := range files {
			ret = append(ret, f.FullPath())
		}
	}
	return ret, nil
}

func getEmptyFolders(tgt filesystem.FileLoader, ignore []string, logger util.Logger, pth ...string) ([]string, error) {
	var ret []string
	pStr := path.Join(pth...)
	fc := len(tgt.ListFiles(pStr, nil, logger))
	ds := tgt.ListDirectories(pStr, ignore, logger)
	if fc == 0 && len(ds) == 0 {
		ret = append(ret, pStr)
	}
	for _, d := range ds {
		p := append(slices.Clone(pth), d)
		childRes, err := getEmptyFolders(tgt, ignore, logger, p...)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to get empty folders for [%s/%s]", path.Join(p...), d)
		}
		ret = append(ret, childRes...)
	}
	return ret, nil
}
