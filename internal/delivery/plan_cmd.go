package delivery

import (
	"fmt"

	"github.com/spf13/cobra"
)

// newPlanCommand returns the Cobra command for `xtai plan`.
// This currently prints a placeholder message; a later implementation will wire
// it to the planner use case and LLM provider layer.
func newPlanCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "plan",
		Short: "Generate an execution plan for a natural-language request",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Fprintln(cmd.OutOrStdout(), "xtai plan: planning is not implemented yet")
			return nil
		},
	}

	return cmd
}

