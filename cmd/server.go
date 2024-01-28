package cmd

import (
	"net/http"

	"github.com/dexguitar/chatapp/configs"
	"github.com/spf13/cobra"
)

var runServerCmd = func(c *configs.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "runserver",
		Short: "Runs a server",
		Long:  `Runs a server on specified host and port (first and second argument)`,
		RunE:  runServer(c),
	}
}

func runServer(c *configs.Config) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		mux := http.NewServeMux()
		mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
			cnt, err := writer.Write([]byte("hello there!"))
			if cnt == 0 || err != nil {
				panic(err)
			}
		})

		if err := http.ListenAndServe(c.Host+c.Port, mux); err != nil {
			panic(err)
		}

		return nil
	}
}
