package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/indiependente/go-stock/client/alphavantage"
	"github.com/indiependente/go-stock/config"
	"github.com/indiependente/go-stock/ui"
	"github.com/indiependente/pkg/http/client"
	"github.com/indiependente/pkg/logger"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	supportedAPIs = []string{"https://www.alphavantage.co"}
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
	sym := kingpin.Arg("symbol", "Stock symbol").Default("").String()
	setupFlag := kingpin.Flag("setup", "Configure the tool will require you providing a few information").
		Default("false").Bool()
	kingpin.Parse()

	if *setupFlag {
		filename, err := setup(supportedAPIs)
		if err != nil {
			return fmt.Errorf("could not complete setup: %w", err)
		}
		log.Info("Setup complete. New configuration saved in " + filename)
		return nil
	}

	log.Info("Using config file: " + *configFile)
	if !fileExists(*configFile) {
		log.Warn("Config file not found. Please create it by running go-stock --setup")
		return fmt.Errorf("file %s does not exist", *configFile)
	}
	conf, err := config.ParseFromFile(*configFile)
	if err != nil {
		return fmt.Errorf("could not parse config file: %w", err)
	}

	if *sym == "" {
		log.Warn("You must provide a stock symbol.")
		kingpin.Usage()
		return errors.New("no symbol provided")
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
	err = ui.HandleData(gq.GlobalQuote.Symbol, gq.GlobalQuote.Quote().Price, mts.Series.TimeSeries())
	if err != nil {
		return fmt.Errorf("could not visualise data: %w", err)
	}
	return nil
}
