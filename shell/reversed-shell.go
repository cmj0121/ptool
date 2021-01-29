/* the simple service provide the PT usage shell script */
package shell

import (
	"bytes"
	"text/template"

	_ "embed"

	"github.com/cmj0121/argparse"
	"github.com/cmj0121/logger"
)

//go:embed assets/reversed-shell
var tmplReversedShell string

// generate the reversed shell
type ReversedShell struct {
	argparse.Help

	// the toggle of the reversed shell language
	Bash   bool `help:"enabled the bash"`
	Python bool `help:"enabled the python"`
	NetCat bool `help:"enabled the netcat"`

	// the target
	IP   string `default:"127.0.0.1" help:"the local ip address"`
	Port int    `default:"5566" help:"the local bind port"`
}

func (shell *ReversedShell) Generate(log *logger.Logger) (data []byte, err error) {
	if !shell.Bash && !shell.Python && !shell.NetCat {
		log.Debug("set all format on")
		shell.Bash = true
		shell.Python = true
		shell.NetCat = true
	}

	tmpl := template.Must(template.New("reversed-shell").Parse(tmplReversedShell))
	buff := bytes.NewBuffer(nil)
	if err = tmpl.Execute(buff, shell); err != nil {
		log.Warn("cannot parse template: %v", err)
		return
	}

	data = buff.Bytes()
	return
}
