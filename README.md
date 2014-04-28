Flag
===

Flag is a drop in replacement for Go's flag package:

	$ cat > gopher.go
		package main

		import "github/namsral/flag"

		flag.Int("width", 10, "width of tunnel to dig")
		flag.Parse()

	$ go run gopher.go -width 12

But it also parses environment variables:

	$ export GOPHER_WIDTH=12
	$ go run gopher.go

And it parses configuration files:

	$ cat > gopher.conf
		width 12

	$ go run gopher.go -config gopher.conf


It's a port of Go's [flag][] package with the addition of two functions: `ParseEnv` and `ParseFile`.

[flag]: http://golang.org/src/pkg/flagconfiguration


Why
---

Why not use one of the many INI, JSON or YAML parsers?

I think it's best practice to keep your configuration simple and use simple data types like bools, ints, floats and strings to control the behaviour of your application. Consider moving more complex data types to the "data" layer.


Usage
---

It's intended for projects which require a simple configuration made available through command-line flags, configuration files and shell environments. It's similar to the original `flag` package.

Example:

	import "github/namsral/flag"
	
	flag.String("config", "", "help message for config")
	flag.Int("age", 24, "help message for age")
	
	flag.Parse()

Order of precedence:

1. Command line options
2. Environment variables
3. Configuration file
4. Default values

The order can be changed by parsing manually:

		flag.String("config", "", "help message for config")
		flag.Int("age", 24, "help message for age")
		
		flag.CommandLine.ParseEnv(path.Base(os.Args[0]))
		flag.CommandLine.Parse(os.Args[1:])
		flag.CommandLine.ParseFile("/etc/command.conf")


#### Parsing Configuration Files

Create a configuration file:

	$ cat > /etc/command.conf
	# empty newlines and lines beginning with a "#" character are ignored.
	name bob
	
	# keys and values can also be separated by the "=" character
	age=20
	
	# booleans can me empty, set with 0, 1, true, false, etc
	hacker
	
Add a "config" flag:

	flag.String("config", "", "help message for config")
		
Run the command:

	$ go run ./command.go -config /etc/command.conf


#### Parsing Environment Variables

Prefix the environment variable with the name of your command in uppercase:

	$ export COMMAND_AGE=44
	$ go run ./command.go

If you want to customise the prefix, just parse the environment variables manually:

	flag.Int("age", 24, "help message for age")

	flag.CommandLine.ParseEnv("MYCOMMAND")
	
That's it.


License
---


Copyright (c) 2012 The Go Authors. All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are
met:

   * Redistributions of source code must retain the above copyright
notice, this list of conditions and the following disclaimer.
   * Redistributions in binary form must reproduce the above
copyright notice, this list of conditions and the following disclaimer
in the documentation and/or other materials provided with the
distribution.
   * Neither the name of Google Inc. nor the names of its
contributors may be used to endorse or promote products derived from
this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
"AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.