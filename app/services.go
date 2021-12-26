// Package app $PF_IGNORE$
package app

import (
	"context"

	"github.com/kyleu/projectforge/app/export"
	"github.com/kyleu/projectforge/app/git"
	"github.com/kyleu/projectforge/app/module"
	"github.com/kyleu/projectforge/app/project"
)

type Services struct {
	Modules  *module.Service
	Projects *project.Service
	Export   *export.Service
	Git      *git.Service
}

func NewServices(_ context.Context, st *State) (*Services, error) {
	return &Services{
		Modules:  module.NewService(st.Files, st.Logger),
		Projects: project.NewService(st.Logger),
		Export:   export.NewService(st.Logger),
		Git:      git.NewService(st.Logger),
	}, nil
}
