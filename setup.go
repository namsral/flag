// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package flag

import (
	"fmt"
	"os"
)

// Setup parses flag and setups usage function.
// `positionalArgs` is a string representing positional args, eg: "arg1 arg2 arg3"
// `commands` if provided is a positional argument command description.
func Setup(version, positionalArgs string, commands ...string) {
	Usage = func() {
		fmt.Fprintln(os.Stderr, "Version: ", version)
		if len(commands) != 0 {
			fmt.Fprintf(os.Stderr,
				"USAGE: %s <commands> <parameters> %s\n\n", os.Args[0], positionalArgs)
			fmt.Fprint(os.Stderr, "COMMANDS:", commands[0], "\n\n")
		} else {
			fmt.Fprintf(os.Stderr,
				"USAGE: %s <parameters> %s\n\n", os.Args[0], positionalArgs)
		}
		fmt.Fprintln(os.Stderr, "PARAMETERS:")
		PrintDefaults()
	}
	Parse()
}
