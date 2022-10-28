package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"primedividend/api/internal/config"
)

var (
	rootCmd = &cobra.Command{}
	cfg     config.Config
)

func init() {
	cobra.OnInitialize(func() {
		cfg = config.GetConfig()
	})
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
