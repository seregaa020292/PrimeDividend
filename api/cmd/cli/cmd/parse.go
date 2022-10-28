package cmd

import (
	"fmt"
	"strings"
	"sync"

	"github.com/spf13/cobra"

	currencyRepo "primedividend/api/internal/modules/currency/repository"
	instrumentRepo "primedividend/api/internal/modules/instrument/repository"
	marketRepo "primedividend/api/internal/modules/market/repository"
	providerRepo "primedividend/api/internal/modules/provider/repository"
	registerRepo "primedividend/api/internal/modules/register/repository"
	"primedividend/api/internal/services/parser"
	"primedividend/api/pkg/db/postgres"
	"primedividend/api/pkg/utils/errlog"
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
	errlog.Fatalln(parseCmd.MarkFlagRequired("instrument"))
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

	errlog.Fatalln(parserService.Select())

	if instrument == "all" {
		var wg sync.WaitGroup

		wg.Add(len(instruments))

		for _, i := range instruments {
			go func(instrument string) {
				defer wg.Done()
				errlog.Fatalln(parserService.Execute(instrument))
			}(i)
		}

		wg.Wait()

		return
	}

	errlog.Fatalln(parserService.Execute(instrument))
}
