package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "chatapp",
	Short: "chatapp",
	Long:  `this chatapp is undoubtedly a piece of art`,
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	rootCmd.AddCommand(runServerCmd())

	if err := rootCmd.Execute(); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
