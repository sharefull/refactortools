package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/sharefull/refactortools/unused"
	"github.com/sharefull/refactortools/unused/internal"
	"golang.org/x/tools/go/packages"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	unused.Analyzer.Flags = flag.NewFlagSet(unused.Analyzer.Name, flag.ExitOnError)
	unused.Analyzer.Flags.Parse(os.Args[1:])

	if unused.Analyzer.Flags.NArg() < 1 {
		return errors.New("patterns of packages must be specified")
	}

	pkgs, err := packages.Load(unused.Analyzer.Config, unused.Analyzer.Flags.Args()...)
	if err != nil {
		return err
	}

	for _, pkg := range pkgs {

		pass := &internal.Pass{
			Package: pkg,
			Stdin:   os.Stdin,
			Stdout:  os.Stdout,
			Stderr:  os.Stderr,
		}

		if err := unused.Analyzer.Run(pass); err != nil {
			return err
		}
	}

	fmt.Println(unused.Defs)
	fmt.Println(unused.Uses)

	fset := unused.Analyzer.Config.Fset
	for id, obj := range unused.Defs {
		if _, ok := unused.Uses[id]; !ok {
			pos := fset.Position(obj.Pos())
			fmt.Println(pos, id, "unused")
		}
	}

	return nil
}