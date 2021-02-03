/* the simple service provide the PT usage shell script */
package shell

import (
	"bytes"
	_ "embed"
	"fmt"
	"os"
	"text/template"

	"github.com/cmj0121/argparse"
	"github.com/cmj0121/logger"
)

var (
	//go:embed assets/welcome
	tmplWelcome string
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

	Bind *string `help:"run the server mode and bind on IP:PORT"`

	*ReversedShell `name:"reversed" help:"generate the reversed-shell script"`
	*Scan          `help:"generate the scan script"`
	*Script        `help:"generate script"`
}

func New() (shell *Shell) {
	shell = &Shell{
		Logger: logger.New(PROJ_NAME),
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
	if err := parser.Run(); err == nil {
		shell.Logger.SetLevel(shell.LogLevel)

		switch {
		case shell.Bind != nil:
			// force clean-up the setting
			shell.ReversedShell = nil
			shell.Scan = nil
			shell.Script = nil
			shell.ListenAndServe()
		default:
			if data, err := shell.Generate(); err == nil {
				// show the script template on STDOUT
				os.Stdout.Write(data)
			}
		}
	}
}

func (shell *Shell) Generate() (data []byte, err error) {
	switch {
	case shell.ReversedShell != nil:
		// generate reversed-shell
		data, err = shell.ReversedShell.Generate(shell.Logger)
	case shell.Scan != nil:
		// generate scan shell
		data, err = shell.Scan.Generate(shell.Logger)
	case shell.Script != nil:
		// generate scan shell
		data, err = shell.Script.Generate(shell.Logger)
	default:
		tmpl := template.Must(template.New("welcome").Parse(tmplWelcome))
		buff := bytes.NewBuffer(nil)
		if err = tmpl.Execute(buff, shell); err != nil {
			shell.Logger.Warn("cannot parse template: %v", err)
			return
		}

		data = buff.Bytes()
	}

	return
}
