package main

import (
	"context"
	"fmt"

	"github.com/indiependente/go-stock/client/alphavantage"
	"github.com/indiependente/go-stock/config"
	"github.com/indiependente/go-stock/ui"
	"github.com/indiependente/pkg/http/client"
	"github.com/indiependente/pkg/logger"
	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	log := logger.GetConsoleLogger("go-stock", logger.INFO)
	err := run(log)
	if err != nil {
		log.Fatal("Stop execution", err)
	}
}

func run(log logger.Logger) error {
	configFile := kingpin.Flag("config", "Configuration file location").Short('c').
		Default("config.yml").String()
	sym := kingpin.Arg("symbol", "Stock symbol").Required().String()
	kingpin.Parse()

	conf, err := config.ParseFromFile(*configFile)
	if err != nil {
		return fmt.Errorf("could not parse config file: %w", err)
	}

	c := client.DefaultHTTPClient(10)
	avClient := alphavantage.Client{
		Doer:   c,
		URL:    conf.URL,
		APIKey: conf.APIKey,
	}
	ctx := context.Background()
	gq, err := avClient.GlobalQuote(ctx, *sym)
	if err != nil {
		return fmt.Errorf("could not get global quote: %w", err)
	}
	mts, err := avClient.MonthlyTimeSeries(ctx, *sym)
	if err != nil {
		return fmt.Errorf("could not get monthly time series: %w", err)
	}
	ui.HandleData(*sym, gq.GlobalQuote.Quote().Price, mts.Series.TimeSeries())
	return nil
}
