/* the simple service provide the PT usage shell script */
package main

import (
	"github.com/cmj0121/ptool/shell"
)

func main() {
	agent := shell.New()
	agent.Run()
}
