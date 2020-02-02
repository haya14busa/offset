package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/haya14busa/offset"
)

// Populated during build.
var version = "master"

type option struct {
	version bool
	offset  int
}

func setupFlags(opt *option) {
	flag.BoolVar(&opt.version, "version", false, "print version")
	flag.IntVar(&opt.offset, "offset", 0, "byte offset")

	flag.Usage = usage
	flag.Parse()
}

func usage() {
	fmt.Fprintln(os.Stderr, "Usage: offset [FLAGS] [File]")
	fmt.Fprintln(os.Stderr, "\tReturn position from given offset")
	fmt.Fprintln(os.Stderr, "Flags:")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "GitHub: https://github.com/haya14busa/offset")
	os.Exit(2)
}

func main() {
	var opt option
	setupFlags(&opt)
	if opt.version {
		fmt.Fprintln(os.Stdout, version)
		os.Exit(0)
	}
	if err := run(os.Stdin, os.Stdout, flag.Args(), opt); err != nil {
		fmt.Fprintf(os.Stderr, "offset: %v\n", err)
		os.Exit(1)
	}
}

func run(r io.Reader, w io.Writer, args []string, opt option) error {
	if len(args) != 1 {
		return errors.New("the number of given files is not one")
	}
	filename := args[0]
	p, err := offset.FromFilename(filename, opt.offset)
	if err != nil {
		return err
	}
	return json.NewEncoder(w).Encode(p)
}
