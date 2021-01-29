package shell

import (
	"regexp"
	_ "embed"
)

// the project information
const (
	PROJ_NAME = "shell"

	MAJOR = 0
	MINOR = 0
	MACRO = 0
)

// the Shell template
//go:embed assets/welcome
var TMPL_Welcome string
//go:embed assets/reversed-shell
var TMPL_ReversedShell string

var (
	RE_REVERSED_SHELL = regexp.MustCompile(`^/reversed(?:/(\d+\.\d+\.\d+\.\d+):(\d+))?$`)
)
