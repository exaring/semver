package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/blang/semver"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %v [OPTIONS] <version>\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "OPTIONS:\n")
		fmt.Fprintf(os.Stderr, "   -major, -minor, -patch   increase version part\n")
		fmt.Fprintf(os.Stderr, "   -build build-name        include additional build name (e.g. alpha)\n")
	}

	major := flag.Bool("major", false, "increase major version")
	minor := flag.Bool("minor", false, "increase minor version")
	patch := flag.Bool("patch", false, "increase patch version")
	build := flag.String("build", "", "set build")
	flag.Parse()

	if len(flag.Args()) == 0 {
		fmt.Fprintf(os.Stderr, "Version missing\n")
		flag.Usage()
		os.Exit(1)
	}

	versionString := flag.Arg(0)
	version, err := semver.New(versionString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Version string %v is not a valid version", versionString)
		os.Exit(1)
	}

	if *major {
		version.Major++
		version.Minor = 0
		version.Patch = 0
		version.Pre = nil
		version.Build = nil
	} else if *minor {
		version.Minor++
		version.Patch = 0
		version.Pre = nil
		version.Build = nil
	} else if *patch {
		version.Patch++
		version.Pre = nil
		version.Build = nil
	}
	if *build != "" {
		version.Build = []string{*build}
	}

	fmt.Println(version.String())
}
