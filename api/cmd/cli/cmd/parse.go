package cmd

import (
	"fmt"
	"strings"
	"sync"

	"github.com/spf13/cobra"

	currencyRepo "primedivident/internal/modules/currency/repository"
	instrumentRepo "primedivident/internal/modules/instrument/repository"
	marketRepo "primedivident/internal/modules/market/repository"
	providerRepo "primedivident/internal/modules/provider/repository"
	registerRepo "primedivident/internal/modules/register/repository"
	"primedivident/internal/services/parser"
	"primedivident/pkg/db/postgres"
	"primedivident/pkg/utils"
)

var parseCmd = &cobra.Command{
	Use: "parse",
	Run: parseCommand,
}

var (
	instrument  string
	instruments = []string{"etfs", "stocks", "bonds", "currencies"}
)

func init() {
	rootCmd.AddCommand(parseCmd)

	parseCmd.Flags().StringVar(
		&instrument,
		"instrument",
		instruments[0],
		fmt.Sprintf("%s, all", strings.Join(instruments, ", ")),
	)
	utils.Fatalln(parseCmd.MarkFlagRequired("instrument"))
}

func parseCommand(cmd *cobra.Command, args []string) {
	db := postgres.NewPostgres(cfg.Postgres)

	parserService := parser.NewParser(
		cfg.Tinkoff,
		instrumentRepo.NewRepository(db),
		currencyRepo.NewRepository(db),
		providerRepo.NewRepository(db),
		marketRepo.NewRepository(db),
		registerRepo.NewRepository(db),
	)

	utils.Fatalln(parserService.Select())

	if instrument == "all" {
		var wg sync.WaitGroup

		wg.Add(len(instruments))

		for _, i := range instruments {
			go func(instrument string) {
				defer wg.Done()
				utils.Fatalln(parserService.Execute(instrument))
			}(i)
		}

		wg.Wait()

		return
	}

	utils.Fatalln(parserService.Execute(instrument))
}
