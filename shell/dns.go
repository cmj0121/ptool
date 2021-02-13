/* the simple service provide the PT usage shell script */
package shell

import (
	"bytes"
	"fmt"
	"text/template"

	_ "embed"

	"github.com/cmj0121/argparse"
	"github.com/cmj0121/logger"
)

//go:embed assets/dns
var tmplDNS string

// generate the DNS recon script
type DNS struct {
	argparse.Help

	Server string  `help:"specified DNS server"`
	Domain *string `help:"the target hostname"`
}

func (shell *DNS) Generate(log *logger.Logger) (data []byte, err error) {
	if shell.Domain == nil || *shell.Domain == "" {
		err = fmt.Errorf("should specified the DOMAIN")
		return
	}

	tmpl := template.Must(template.New("dns").Parse(tmplDNS))
	buff := bytes.NewBuffer(nil)
	if err = tmpl.Execute(buff, shell); err != nil {
		log.Warn("cannot parse template: %v", err)
		return
	}

	data = buff.Bytes()
	return
}
