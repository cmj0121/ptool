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
var tmplWelcome string
//go:embed assets/reversed-shell
var tmplReversedShell string
//go:embed assets/scan
var tmplScanIPHost string
//go:embed assets/web
var tmplWeb string

var (
	RE_REVERSED_SHELL = regexp.MustCompile(`^/reversed(?:/(\d+\.\d+\.\d+\.\d+):(\d+))?$`)
	RE_SCAN_IP_HOST = regexp.MustCompile(`^/scan/([\w\.-]+)$`)
	RE_WEB_SCAN = regexp.MustCompile(`^/web(?:/([\w\.-]+))$`)
)
