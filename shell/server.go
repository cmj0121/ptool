/* the simple HTTP RESTFul API server */
package shell

import (
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

const (
	KEY_FORMAT = "format"

	KEY_FMT_PYTHON = "python"
	KEY_FMT_BASH   = "bash"
	KEY_FMT_NETCAT = "netcat"
	KEY_FMT_WEB    = "web"
)

var (
	RE_REVERSED = regexp.MustCompile(`^/reversed(?:/(\d+\.\d+\.\d+\.\d+):(\d+))?$`)
	RE_SCAN     = regexp.MustCompile(`^/scan(?:/([\w\.-]+))?$`)
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

		data, _ = agent.Generate(shell.Logger)
	case RE_SCAN.MatchString(path):
		match := RE_SCAN.FindStringSubmatch(path)
		agent := &Scan{}

		switch format := r.URL.Query()[KEY_FORMAT]; strings.Join(format, "") {
		case KEY_FMT_WEB:
			agent.Web = true
		}

		if agent.IPHost = match[1]; agent.IPHost == "" {
			agent.IPHost = "127.0.0.1"
		}

		data, _ = agent.Generate(shell.Logger)
	default:
		data, _ = shell.Generate()
	}

	w.Write(data)
}