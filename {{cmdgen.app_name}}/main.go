package main

import (
	{% if cmdgen.use_cobra_cmd == "n" %}"flag"
	"fmt"
	"github.com/{{cmdgen.github_username}}/{{cmdgen.app_name}}/version"{% endif %}
	{% if cmdgen.use_cobra_cmd == "y" %}"github.com/{{cmdgen.github_username}}/{{cmdgen.app_name}}/cmd"{% endif %}
)

func main() {

    {% if cmdgen.use_cobra_cmd == "y" %}
    cmd.Execute()
	{% else %}
	versionFlag := flag.Bool("version", false, "Version")
	flag.Parse()

	if *versionFlag {
		fmt.Println("Build Date:", version.BuildDate)
        fmt.Println("Git Commit:", version.GitCommit)
        fmt.Println("Version:", version.Version)
        fmt.Println("Go Version:", version.GoVersion)
        fmt.Println("OS / Arch:", version.OsArch)
		return
	}
	fmt.Println("Hello.")
    {% endif %}
}
