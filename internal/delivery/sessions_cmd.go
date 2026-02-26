package delivery

import (
	"fmt"

	"github.com/spf13/cobra"
)

// newSessionsCommand returns the Cobra command for `xtai sessions`.
// This currently prints a placeholder message; a real implementation will query
// the session store for persisted sessions.
func newSessionsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sessions",
		Short: "List or inspect xtai sessions",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Fprintln(cmd.OutOrStdout(), "xtai sessions: session listing is not implemented yet")
			return nil
		},
	}

	return cmd
}

