// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package config

import (
	"flag"
)

const (
	schema = "input/schema.json"
	output = "output"
)

type Config struct {
	Input   string
	Output  string
	Schema  string
	Include []string
}

type stringsFlags []string

func (s *stringsFlags) String() string {
	return "my string representation"
}

func (s *stringsFlags) Set(value string) error {
	*s = append(*s, value)
	return nil
}

var includeFlags stringsFlags

func New(args []string) (Config, error) {
	var i, o, s string

	fs := flag.NewFlagSet("all", flag.ExitOnError)
	fs.StringVar(&i, "input", "", "Path to intermediate representation")
	fs.StringVar(&o, "output", "", "Directory for generated code files")
	fs.StringVar(&s, "schema", "", "Path or URL to intermediate representation JSON schema")
	fs.Var(&includeFlags, "include", "Specify which data sources, provider and resources to include on the basis of name")

	err := fs.Parse(args)
	if err != nil {
		return Config{}, err
	}

	config := Config{
		Input:  i,
		Schema: schema,
		Output: output,
	}

	if s != "" {
		config.Schema = s
	}

	if o != "" {
		config.Output = o
	}

	return config, nil
}
