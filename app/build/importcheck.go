package build

import (
	"strings"

	"github.com/pkg/errors"

	"projectforge.dev/projectforge/app/util"
)

const sepKey, firstKey, thirdKey, selfKey = "sep", "1st", "3rd", "self"

func check(imports []string, orig []string) ([]string, error) {
	var err error
	var state int
	var lastLine string
	var observed []string
	first, third, self := util.ValueMap{}, util.ValueMap{}, util.ValueMap{}
	observe := func(key string, i string) {
		for _, ob := range observed {
			if ob > i {
				err = errors.Errorf("%s sorting", key)
			}
		}
		observed = append(observed, i)
	}
	clear := func() {
		observed = []string{}
	}
	var lastSep bool

	for idx, imp := range imports {
		i, l := util.StringSplitLast(imp, ':', true)
		switch l {
		case sepKey:
			if state != 0 && lastLine != "" {
				state++
				clear()
			}
			lastSep = true
		case firstKey:
			if state != 1 {
				state = 1
				clear()
			}
			lastSep = false
			if state > 1 {
				err = errors.New("1st party")
			}
			first[i] = orig[idx]
			observe(i, "1st party")
		case thirdKey:
			if state != 2 {
				if !lastSep && len(first) > 0 {
					err = errors.New("missing separator")
				}
				state = 2
				clear()
			}
			lastSep = false
			if state > 2 {
				err = errors.New("3rd party")
			}
			third[i] = orig[idx]
			observe(i, "3rd party")
		case selfKey:
			if state != 3 {
				if !lastSep && (len(third) > 0 || len(first) > 0) {
					err = errors.New("missing separator")
				}
				state = 3
				clear()
			}
			lastSep = false
			if state > 3 {
				err = errors.New("self")
			}
			self[i] = orig[idx]
			observe(i, selfKey)
		default:
			return nil, errors.New("invalid type")
		}
		lastLine = l
	}
	return makeResult(first, third, self), err
}

func makeResult(first util.ValueMap, third util.ValueMap, self util.ValueMap) []string {
	ret := make([]string, 0, len(first)+len(third)+len(self))
	for _, k := range first.Keys() {
		s, _ := first[k].(string)
		ret = append(ret, s)
	}
	if len(first) > 0 && (len(third) > 0 || len(self) > 0) {
		ret = append(ret, "")
	}
	for _, k := range third.Keys() {
		s, _ := third[k].(string)
		ret = append(ret, s)
	}
	if len(third) > 0 && len(self) > 0 {
		ret = append(ret, "")
	}
	for _, k := range self.Keys() {
		s, _ := self[k].(string)
		ret = append(ret, s)
	}
	return ret
}

func importType(i string, self string) string {
	if i == "" {
		return sepKey
	}
	if strings.HasPrefix(i, self) || strings.HasPrefix(i, "{{{ .Package }}}") {
		return selfKey
	}
	if strings.Contains(i, ".") {
		return thirdKey
	}
	return firstKey
}
