// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/go-enjin/golang-org-x-text/message/pipeline"
)

var cmdGenerate = &Command{
	Init:      initGenerate,
	Run:       runGenerate,
	UsageLine: "generate <package> [-out <gofile>]",
	Short:     "generates code to insert translated messages",
}

func initGenerate(cmd *Command) {
	out = cmd.Flag.String("out", "", "output file to write to")
}

func runGenerate(cmd *Command, config *pipeline.Config, args []string) error {
	config.Packages = args
	s, err := pipeline.Extract(config)
	if err != nil {
		return wrap(err, "extraction failed")
	}
	if err := s.Import(); err != nil {
		return wrap(err, "import failed")
	}
	if err := s.Merge(); err != nil {
		return wrap(err, "merge failed")
	}
	return wrap(s.Generate(), "generation failed")
}
