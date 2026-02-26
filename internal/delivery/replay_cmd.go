package delivery

import (
	"fmt"

	"github.com/spf13/cobra"
)

// newReplayCommand returns the Cobra command for `xtai replay`.
// In Phase 1 this prints a placeholder message; the session store and replay
// mechanics will be added in later phases.
func newReplayCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "replay",
		Short: "Replay a previous xtai session by ID",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Fprintln(cmd.OutOrStdout(), "xtai replay: session replay is not implemented yet")
			return nil
		},
	}

	return cmd
}

