package delivery

import "github.com/spf13/cobra"

// NewRootCommand constructs the root cobra.Command for xtai and wires all
// subcommands that are part of the Phase 1 CLI surface.
func NewRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "xtai",
		Short: "xtai is a natural-language CLI agent for developer workflows",
		Long:  "xtai is a natural-language CLI agent that plans and executes safe, observable workflows across your developer tools.",
	}

	rootCmd.AddCommand(
		newPlanCommand(),
		newExecuteCommand(),
		newReplayCommand(),
		newSessionsCommand(),
	)

	return rootCmd
}

// Execute runs the root command. It is called from the main package.
func Execute() error {
	return NewRootCommand().Execute()
}

