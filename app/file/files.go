package file

import (
	"strings"
)

type Files []*File

func (f Files) Apply(m *Changeset) (Files, error) {
	ret := make(Files, 0, len(f))
	for _, file := range f {
		file.Apply(m)
	}
	return ret, nil
}

func (f Files) String() string {
	var sb strings.Builder
	for _, file := range f {
		sb.WriteString(" - ")
		sb.WriteString(file.FullPath())
		sb.WriteString("\n")
	}
	return sb.String()
}

func (f Files) Get(path string) *File {
	for _, file := range f {
		if file.FullPath() == path {
			return file
		}
	}
	return nil
}

func (f Files) Exists(path string) bool {
	return f.Get(path) != nil
}
