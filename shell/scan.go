/* the simple service provide the PT usage shell script */
package shell

import (
	"bytes"
	_ "embed"
	"text/template"

	"github.com/cmj0121/argparse"
	"github.com/cmj0121/logger"
)

//go:embed assets/scan
var tmplScan string

type Scan struct {
	argparse.Help

	// the toggle of the scan method, default is nmap
	Web bool `help:"scan the web info"`

	// the target
	IPHost string `default:"127.0.0.1"`
}

func (shell *Scan) Generate(log *logger.Logger) (data []byte, err error) {
	tmpl := template.Must(template.New("scan").Parse(tmplScan))
	buff := bytes.NewBuffer(nil)
	if err = tmpl.Execute(buff, shell); err != nil {
		log.Warn("cannot parse template: %v", err)
		return
	}

	data = buff.Bytes()
	return
}
