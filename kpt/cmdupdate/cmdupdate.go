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

// Package cmdupdate contains the update command
package cmdupdate

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"kpt.dev/kpt/generated"
	"kpt.dev/kpt/util/update"
)

// NewRunner returns a command runner.
func NewRunner() *Runner {
	r := &Runner{}
	c := &cobra.Command{
		Use:        "update LOCAL_PKG_DIR[@VERSION]",
		Short:      generated.UpdateShort,
		Long:       generated.UpdateLong,
		Example:    generated.UpdateExamples,
		RunE:       r.runE,
		Args:       cobra.ExactArgs(1),
		PreRunE:    r.preRunE,
		SuggestFor: []string{"rebase", "replace"},
	}

	c.Flags().StringVarP(&r.Update.Repo, "repo", "r", "",
		"git repo url for updating contents.  defaults to the repo the package was fetched from.")
	c.Flags().StringVar(&r.strategy, "strategy", string(update.FastForward),
		"update strategy for preserving changes to the local package.")
	c.Flags().BoolVar(&r.Update.DryRun, "dry-run", false,
		"print the git patch rather than merging it.")
	c.Flags().BoolVar(&r.Update.Verbose, "verbose", false,
		"print verbose logging information.")
	r.Command = c
	return r
}

func NewCommand() *cobra.Command {
	return NewRunner().Command
}

// Runner contains the run function.
// TODO, support listing versions
type Runner struct {
	strategy string
	Update   update.Command
	Command  *cobra.Command
}

func (r *Runner) preRunE(c *cobra.Command, args []string) error {
	r.Update.Strategy = update.StrategyType(r.strategy)
	parts := strings.Split(args[0], "@")
	if len(parts) > 2 {
		return fmt.Errorf("at most 1 version permitted")
	}
	r.Update.Path = parts[0]
	if len(parts) > 1 {
		r.Update.Ref = parts[1]
	}

	return nil
}

func (r *Runner) runE(c *cobra.Command, args []string) error {
	if len(r.Update.Ref) > 0 {
		fmt.Fprintf(c.ErrOrStderr(), "updating package %s to %s\n",
			r.Update.Path, r.Update.Ref)
	} else {
		fmt.Fprintf(c.ErrOrStderr(), "updating package %s\n",
			r.Update.Path)
	}
	return r.Update.Run()
}
