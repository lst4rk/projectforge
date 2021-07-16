package action

import (
	"github.com/kyleu/projectforge/app/util"
)

type Type struct {
	Key         string
	Title       string
	Icon        string
	Description string
}

var (
	TypeBuild   = Type{Key: "build", Title: "Build", Icon: "star", Description: "Builds the project"}
	TypeCreate  = Type{Key: "create", Title: "Create", Icon: "star", Description: "Creates a new project"}
	TypeDebug   = Type{Key: "debug", Title: "Debug", Icon: "star", Description: "Dumps a ton of information about the project"}
	TypeMerge   = Type{Key: "merge", Title: "Merge", Icon: "star", Description: "Merges changed files as required"}
	TypePreview = Type{Key: "preview", Title: "Preview", Icon: "star", Description: "Show what would happen if you did merge"}
	TypeSlam    = Type{Key: "slam", Title: "Slam", Icon: "star", Description: "Slams all files to the target, ignoring changes"}
	TypeSVG     = Type{Key: "svg", Title: "SVG", Icon: "star", Description: "Builds the project's SVG files"}
	TypeTest    = Type{Key: "test", Title: "Test", Icon: "star", Description: "Runs internal tests, you probably don't want this"}
)

var AllTypes = []Type{TypeBuild, TypeCreate, TypeDebug, TypeMerge, TypePreview, TypeSlam, TypeSVG, TypeTest}
var ProjectTypes = []Type{TypePreview, TypeDebug, TypeMerge, TypeSlam, TypeSVG, TypeBuild}

func TypeFromString(s string) Type {
	for _, t := range AllTypes {
		if t.Key == s {
			return t
		}
	}
	return errorType("No action type available with key [" + s + "]")
}

func errorType(msg string) Type {
	return Type{Key: "error", Title: "Error", Icon: "star", Description: msg}
}

func (t *Type) String() string {
	return t.Key
}

func (t *Type) MarshalJSON() ([]byte, error) {
	return util.ToJSONBytes(t.Key, false), nil
}

func (t *Type) UnmarshalJSON(data []byte) error {
	var s string
	if err := util.FromJSON(data, &s); err != nil {
		return err
	}
	x := TypeFromString(s)
	*t = x
	return nil
}
