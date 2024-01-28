package cmd

import (
	"fmt"
	"log"

	"github.com/dexguitar/chatapp/configs"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "chatapp",
	Short: "chatapp",
	Long:  `this chatapp is undoubtedly a piece of art`,
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	c, err := configs.LoadConfig(".env")
	if err != nil {
		panic(fmt.Sprintf("cannot load config: %s", err.Error()))
		return
	}

	rootCmd.AddCommand(runServerCmd(c))
	rootCmd.AddCommand(migrateCmd(c))

	if err = rootCmd.Execute(); err != nil {
		log.Fatalf("error executing chatapp: '%s'", err.Error())
	}
}
