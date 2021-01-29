/* the simple service provide the PT usage shell script */
package shell

import (
	"fmt"
	"net/http"
	_ "embed"

	"github.com/cmj0121/logger"
	"github.com/cmj0121/argparse"
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
var Welcome string

func Version() (ver string) {
	ver = fmt.Sprintf("%v (v%d.%d.%d)", PROJ_NAME, MAJOR, MINOR, MACRO)
	return
}

type Shell struct {
	argparse.Model

	// the internal logger
	*logger.Logger `-`
	LogLevel       string `name:"log" choices:"warn info debug verbose" help:"log level"`

	Bind string
}

func New() (shell *Shell) {
	shell = &Shell {
		Logger:    logger.New(PROJ_NAME),
		Bind: ":12345",
	}
	return
}

func (shell *Shell) callbackVersion(parser *argparse.ArgParse) (exit bool) {
	fmt.Println(Version())
	return
}

func (shell *Shell) Run() {
	argparse.RegisterCallback(argparse.FN_VERSION, shell.callbackVersion)
	parser := argparse.MustNew(shell)
	switch err := parser.Run(); err {
	case nil:
		shell.Logger.SetLevel(shell.LogLevel)
		shell.Logger.Info("start run on %s", shell.Bind)
		if err := http.ListenAndServe(shell.Bind, shell); err != nil {
			shell.Logger.Crit("http server: %v", err)
		}
	}
}

func (shell *Shell) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	shell.Logger.Info("serve %-22s [%s] %s",r.RemoteAddr, r.Method, r.URL)


	w.Header().Add("Content-Type", "text/plain")
	w.Write([]byte(Welcome))
}

