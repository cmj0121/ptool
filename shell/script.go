/* the simple service provide the PT usage shell script */
package shell

import (
	"bytes"
	"text/template"

	_ "embed"

	"github.com/cmj0121/argparse"
	"github.com/cmj0121/logger"
)

//go:embed assets/script
var tmplScript string

type Script struct {
	argparse.Help

	PHPRFI bool
}

func (script *Script) Generate(log *logger.Logger) (data []byte, err error) {
	tmpl := template.Must(template.New("script").Parse(tmplScript))
	buff := bytes.NewBuffer(nil)
	if err = tmpl.Execute(buff, script); err != nil {
		log.Warn("cannot parse template: %v", err)
		return
	}

	data = buff.Bytes()
	return
}
