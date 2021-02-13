package shell

import (
	"fmt"

	"github.com/cmj0121/logger"
)

// the project information
const (
	PROJ_NAME = "shell"

	MAJOR = 1
	MINOR = 0
	MACRO = 0
)

func Version() (ver string) {
	ver = fmt.Sprintf("%v (v%d.%d.%d)", PROJ_NAME, MAJOR, MINOR, MACRO)
	return
}

// shell-script generator
type Generator interface {
	Generate(log *logger.Logger) (data []byte, err error)
}
