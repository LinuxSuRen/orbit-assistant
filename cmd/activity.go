package cmd

import "github.com/spf13/cobra"

func NewCommand() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use: "activity",
	}
	cmd.AddCommand(newActivityListCommand(), newActivityUpdateCommand())
	return
}
