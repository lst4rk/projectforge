package action

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/kyleu/projectforge/app/project"
	"github.com/kyleu/projectforge/app/util"
)

func clilog(s string) {
	print(s) // nolint
}

func cliProject(p *project.Project, modKeys []string) error {
	if p.Key == "" || p.Key == "TODO" {
		path, _ := os.Getwd()
		if strings.Contains(path, "/") {
			path = path[strings.LastIndex(path, "/")+1:]
		}
		p.Key = path
	}
	p.Key = promptString("Enter a project key; must only contain alphanumerics", p.Key)

	if p.Name == "" {
		p.Name = p.Key
	}
	p.Name = promptString("Enter a project name; use title case and spaces if needed", p.Name)

	if p.Icon == "" {
		p.Icon = "star"
	}

	if p.Exec == "" {
		p.Exec = p.Key
	}

	p.Version = promptString("Enter a version, such as 0.0.0", p.Version)

	if p.Info.Org == "" {
		p.Info.Org = "todo"
	}
	p.Info.Org = promptString("Enter the github organization that owns this project", p.Info.Org)

	if p.Package == "" {
		p.Package = "github.com/" + p.Info.Org + "/" + p.Key
	}
	p.Package = promptString("Enter your project's package", p.Package)

	if p.Info.Homepage == "" {
		p.Info.Homepage = "https://" + p.Package
	}
	p.Info.Homepage = promptString("Enter this project's home page", p.Info.Homepage)

	if p.Port == 0 {
		p.Port = 20000
	}
	p.Port, _ = strconv.Atoi(promptString("Enter the default port your http server will run on", fmt.Sprint(p.Port)))

	modPromptString := "Enter the modules your project will use as a comma-separated list (or \"all\") from choices:\n  " + strings.Join(modKeys, ", ")
	mods := promptString(modPromptString, strings.Join(p.Modules, ", "))
	p.Modules = util.SplitAndTrim(mods, ",")
	if len(p.Modules) == 1 && p.Modules[0] == "all" {
		p.Modules = modKeys
	}

	if p.Info.License == "" {
		p.Info.License = "Proprietary"
	}
	p.Info.License = promptString("Enter the license used by this project", p.Info.License)

	return nil
}

func promptString(query string, curr string) string {
	clilog(query)
	if curr != "" {
		clilog(" (default: " + curr + ")")
	}
	clilog("\n")
	clilog(" > ")
	text, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		clilog("error: " + err.Error() + "\n")
	}
	text = strings.TrimSuffix(text, "\n")
	if text == "" {
		text = curr
	}
	return text
}
