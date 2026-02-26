package delivery

import (
	"fmt"

	"github.com/spf13/cobra"
)

// newExecuteCommand returns the Cobra command for `xtai execute`.
// In Phase 1 this only prints a placeholder message; the safe execution
// engine and policy system will be introduced in later phases.
func newExecuteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "execute",
		Short: "Execute a previously generated plan or direct instructions",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Fprintln(cmd.OutOrStdout(), "xtai execute: execution engine is not implemented yet")
			return nil
		},
	}

	return cmd
}

