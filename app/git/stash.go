package git

import (
	"context"

	"github.com/pkg/errors"
	"projectforge.dev/projectforge/app/project"
	"projectforge.dev/projectforge/app/util"
)

func (s *Service) gitStashApply(ctx context.Context, prj *project.Project, logger util.Logger) error {
	out, err := gitCmd(ctx, "stashXXX", prj.Path, logger)
	if err != nil {
		if isNoRepo(err) {
			return nil
		}
		return errors.Wrap(err, "unable to apply stash")
	}
	println(out)
	return nil
}

func (s *Service) gitStashPop(ctx context.Context, prj *project.Project, logger util.Logger) error {
	out, err := gitCmd(ctx, "stashXXX", prj.Path, logger)
	if err != nil {
		if isNoRepo(err) {
			return nil
		}
		return errors.Wrap(err, "unable to pop stash")
	}
	println(out)
	return nil
}
