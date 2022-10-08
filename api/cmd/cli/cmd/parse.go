package cmd

import (
	"github.com/spf13/cobra"

	currencyRepo "primedivident/internal/modules/currency/repository"
	instrumentRepo "primedivident/internal/modules/instrument/repository"
	marketRepo "primedivident/internal/modules/market/repository"
	providerRepo "primedivident/internal/modules/provider/repository"
	registerRepo "primedivident/internal/modules/register/repository"
	"primedivident/internal/services/parse"
	"primedivident/pkg/db/postgres"
	"primedivident/pkg/utils"
)

var parseCmd = &cobra.Command{
	Use: "parse",
	Run: parseCommand,
}

var instrument string

func init() {
	rootCmd.AddCommand(parseCmd)

	parseCmd.Flags().StringVar(&instrument, "instrument", "etfs", "stocks, bonds, etfs, currencies")
	utils.Fatalln(parseCmd.MarkFlagRequired("instrument"))
}

func parseCommand(cmd *cobra.Command, args []string) {
	db := postgres.NewPostgres(cfg.Postgres)

	parseService := parse.NewParse(
		cfg.Tinkoff,
		instrumentRepo.NewRepository(db),
		currencyRepo.NewRepository(db),
		providerRepo.NewRepository(db),
		marketRepo.NewRepository(db),
		registerRepo.NewRepository(db),
	)

	utils.Fatalln(parseService.Execute(instrument))
}
