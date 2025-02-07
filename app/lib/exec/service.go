// Content managed by Project Forge, see [projectforge.md] for details.
package exec

import (
	"sync"
)

type Service struct {
	Execs Execs
	mu    sync.Mutex
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) NewExec(key string, cmd string, path string, envvars ...string) (*Exec, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	idx := len(s.Execs.GetByKey(key)) + 1
	e := NewExec(key, idx, cmd, path, envvars...)
	s.Execs = append(s.Execs, e)
	s.Execs.Sort()
	return e, nil
}
