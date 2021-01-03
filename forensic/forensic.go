package main

import (
	"fmt"

	"github.com/cmj0121/argparse"
)

const (
	PROJ_NAME = "forensic"

	MAJOR = 0
	MINOR = 0
	MACRO = 0
)

type Forensic struct {
	argparse.Model
}

func Version(parser *argparse.ArgParse) (exit bool) {
	fmt.Printf("%s/v%d.%d.%d\n", PROJ_NAME, MAJOR, MINOR, MACRO)
	return
}

func main() {
	argparse.RegisterCallback(argparse.FN_VERSION, Version)

	forensic := Forensic{}
	parser := argparse.MustNew(&forensic)
	parser.Run()
}
