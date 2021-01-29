/* the simple service provide the PT usage shell script */
package shell

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/cmj0121/logger"
	"github.com/cmj0121/argparse"
)

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

	// always set the header as text/plain
	w.Header().Add("Content-Type", "text/plain")

	// the query PATH
	path := r.URL.EscapedPath()
	shell.Logger.Debug("process the PATH: %#v", path)
	switch {
	case RE_REVERSED_SHELL.MatchString(path):
		ip := "127.0.0.1"
		port := "8888"

		matched := RE_REVERSED_SHELL.FindStringSubmatch(path)
		if matched[1] != "" {
			ip = matched[1]
		}
		if matched[2] != "" {
			port = matched[2]
		}

		tmpl := template.Must(template.New("reversed-shell").Parse(TMPL_ReversedShell))
		tmpl.Execute(w, struct {
			IP string
			PORT string
		} {
			IP: ip,
			PORT: port,
		})
	default:
		w.Write([]byte(TMPL_Welcome))
	}
}

