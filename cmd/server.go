package cmd

import (
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/samazon/server"
	"github.com/spf13/cobra"
)

func init() {
	samazonServerCmd.Flags.Flags().Int("rest-port", 8088, "Change the port the restful service is listening on.")
}

// server runs a server that allows for various clients to connect
var samazonServerCmd = &cobra.Command{
	Use:          "server",
	SilenceUsage: true,
	Short:        "Runs the samazon server listening",
	RunE: func(cmd *cobra.Command, args []string) error {
		restPort, _ := cmd.Flags().GetInt("rest-port")

		rest := server.NewServer(restPort)

		Address := fmt.Sprintf(":%v", restPort)
		if err := http.ListenAndServe(Address, rest.Router); err != nil {
			log.Error().Err(err).Msg("failed to launch server")
		}

		return nil
	},
}
