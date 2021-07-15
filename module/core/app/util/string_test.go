package util_test

import (
	"testing"

	"$PF_PACKAGE$/app/util"

	"github.com/pkg/errors"
)

type ToTitleTest struct {
	TestValue string
	Expected  string
}

func (t *ToTitleTest) Test() error {
	if res := util.ToTitle(t.TestValue); res != t.Expected {
		return errors.Errorf("ToTitle returned [%s], not expected [%s]", res, t.Expected)
	}
	return nil
}

var tests = []*ToTitleTest{
	{TestValue: "SimpleCamelCase", Expected: "Simple Camel Case"},
	{TestValue: "CSVFilesAreCoolButTXTRules", Expected: "CSV Files Are Cool But TXT Rules"},
	{TestValue: "MediaTypes", Expected: "Media Types"},
}

func TestToTitle(t *testing.T) {
	t.Parallel()
	for _, test := range tests {
		err := test.Test()
		if err != nil {
			t.Error(err)
		}
	}
}
