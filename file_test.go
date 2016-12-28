// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package flag_test

import (
	"os"
	"testing"
	"time"

	. "github.com/namsral/flag"
)

// Test parsing a configuration file
func TestParseFile(t *testing.T) {

	f := NewFlagSet(os.Args[0], ContinueOnError)

	boolFlag := f.Bool("bool", false, "bool value")
	bool2Flag := f.Bool("bool2", false, "bool2 value")
	intFlag := f.Int("int", 0, "int value")
	int64Flag := f.Int64("int64", 0, "int64 value")
	uintFlag := f.Uint("uint", 0, "uint value")
	uint64Flag := f.Uint64("uint64", 0, "uint64 value")
	stringFlag := f.String("string", "0", "string value")
	float64Flag := f.Float64("float64", 0, "float64 value")
	durationFlag := f.Duration("duration", 5*time.Second, "time.Duration value")

	err := f.ParseFile("./testdata/test.conf")
	if err != nil {
		t.Fatal("expected no error; got ", err)
	}
	if *boolFlag != true {
		t.Error("bool flag should be true, is ", *boolFlag)
	}
	if *bool2Flag != true {
		t.Error("bool2 flag should be true, is ", *bool2Flag)
	}
	if *intFlag != 22 {
		t.Error("int flag should be 22, is ", *intFlag)
	}
	if *int64Flag != 0x23 {
		t.Error("int64 flag should be 0x23, is ", *int64Flag)
	}
	if *uintFlag != 24 {
		t.Error("uint flag should be 24, is ", *uintFlag)
	}
	if *uint64Flag != 25 {
		t.Error("uint64 flag should be 25, is ", *uint64Flag)
	}
	if *stringFlag != "hello" {
		t.Error("string flag should be `hello`, is ", *stringFlag)
	}
	if *float64Flag != 2718e28 {
		t.Error("float64 flag should be 2718e28, is ", *float64Flag)
	}
	if *durationFlag != 2*time.Minute {
		t.Error("duration flag should be 2m, is ", *durationFlag)
	}
}

func TestParseFileUnknownFlag(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	if err := f.ParseFile("./testdata/bad_test.conf"); err == nil {
		t.Error("parse did not fail for unknown flag; ", err)
	}
}

func TestDefaultConfigFlagname(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)

	f.Bool("bool", false, "bool value")
	f.Bool("bool2", false, "bool2 value")
	f.Int("int", 0, "int value")
	f.Int64("int64", 0, "int64 value")
	f.Uint("uint", 0, "uint value")
	f.Uint64("uint64", 0, "uint64 value")
	stringFlag := f.String("string", "0", "string value")
	f.Float64("float64", 0, "float64 value")
	f.Duration("duration", 5*time.Second, "time.Duration value")

	f.String(DefaultConfigFlagname, "./testdata/test.conf", "config path")

	if err := os.Unsetenv("STRING"); err != nil {
		t.Error(err)
	}

	if err := f.Parse([]string{}); err != nil {
		t.Error("parse failed; ", err)
	}

	if *stringFlag != "hello" {
		t.Error("string flag should be `hello`, is", *stringFlag)
	}
}

func TestDefaultConfigFlagnameMissingFile(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.String(DefaultConfigFlagname, "./testdata/missing", "config path")

	if err := os.Unsetenv("STRING"); err != nil {
		t.Error(err)
	}
	if err := f.Parse([]string{}); err == nil {
		t.Error("expected error of missing config file, got nil")
	}
}
