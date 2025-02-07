// Content managed by Project Forge, see [projectforge.md] for details.
package util

import (
	"strings"
)

type Pkg []string

func (p Pkg) Quoted(quote string) string {
	ret := make(Pkg, 0, len(p))
	for _, x := range p {
		ret = append(ret, quote+x+quote)
	}
	return strings.Join(ret, ".")
}

func (p Pkg) StringWith(extra ...string) string {
	return strings.Join(append(p, extra...), "::")
}

func (p Pkg) String() string {
	return p.StringWith()
}

func (p Pkg) ToPath(extra ...string) string {
	return strings.Join(append(p, extra...), "/")
}

func (p Pkg) Equals(other Pkg) bool {
	if len(p) != len(other) {
		return false
	}
	return p.StartsWith(other)
}

func (p Pkg) Trim(src Pkg) Pkg {
	ret := make(Pkg, 0, len(p))
	for idx, v := range p {
		if len(src) >= idx && src[idx] == v {
			continue
		}
		ret = append(ret, v)
	}
	return ret
}

func (p Pkg) Drop(n int) Pkg {
	l := len(p) - n
	if l < 0 {
		l = 0
	}
	ret := make(Pkg, 0, l)
	for idx := 0; idx < l; idx++ {
		ret = append(ret, p[idx])
	}
	return ret
}

func (p Pkg) Last() string {
	return p[len(p)-1]
}

func (p Pkg) Shift() Pkg {
	ret := make(Pkg, 0, len(p)-1)
	for i, s := range p {
		if i == len(p)-1 {
			break
		}
		ret = append(ret, s)
	}
	return ret
}

func (p Pkg) Push(name string) Pkg {
	if strings.Contains(name, "/") {
		return Pkg{"ERROR:contains-slash"}
	}
	return append(p, name)
}

func (p Pkg) With(key string) Pkg {
	return append(append(Pkg{}, p...), key)
}

func (p Pkg) StartsWith(t Pkg) bool {
	if len(p) < len(t) {
		return false
	}
	for i, x := range t {
		if p[i] != x {
			return false
		}
	}
	return true
}

func DefaultIfNil[T any](ptr *T, d T) T {
	if ptr == nil {
		return d
	}
	return *ptr
}
