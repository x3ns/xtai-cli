package delivery

import (
	"fmt"

	"github.com/spf13/cobra"
)

// newReplayCommand returns the Cobra command for `xtai replay`.
// This currently prints a placeholder message; the session store and replay
// mechanics will be added in a later implementation.
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

