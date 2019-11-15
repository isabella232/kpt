// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package cmdman contains the man command.
package cmdman

import (
	"github.com/spf13/cobra"
	"kpt.dev/kpt/generated"
	"kpt.dev/kpt/util/man"
)

// NewRunner returns a command runner.
func NewRunner() *Runner {
	r := &Runner{}
	c := &cobra.Command{
		Use:     "man LOCAL_PKG_DIR",
		Args:    cobra.MaximumNArgs(1),
		Short:   generated.ManShort,
		Long:    generated.ManLong,
		Example: generated.ManExamples,
		RunE:    r.runE,
		PreRunE: r.preRunE,
	}

	r.Command = c
	return r
}

func NewCommand() *cobra.Command {
	return NewRunner().Command
}

type Runner struct {
	Man     man.Command
	Command *cobra.Command
}

func (r *Runner) preRunE(c *cobra.Command, args []string) error {
	r.Man.Path = "."
	if len(args) > 0 {
		r.Man.Path = args[0]
	}
	r.Man.StdOut = c.OutOrStdout()
	return nil
}

func (r *Runner) runE(c *cobra.Command, args []string) error {
	return r.Man.Run()
}
