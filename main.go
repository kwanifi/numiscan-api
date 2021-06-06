package main

import (
	"os"
	"os/signal"
	"time"

	"github.com/kwanifi/numiscan-api/api"
	"github.com/kwanifi/numiscan-api/config"
	"github.com/kwanifi/numiscan-api/dao"
	"github.com/kwanifi/numiscan-api/log"
	"github.com/kwanifi/numiscan-api/services"
	"github.com/kwanifi/numiscan-api/services/modules"
	"github.com/kwanifi/numiscan-api/services/parser/hub3"
	"github.com/kwanifi/numiscan-api/services/scheduler"
)

func main() {
	err := os.Setenv("TZ", "UTC")
	if err != nil {
		log.Fatal("os.Setenv (TZ): %s", err.Error())
	}

	cfg := config.GetConfig()
	d, err := dao.NewDAO(cfg)
	if err != nil {
		log.Fatal("dao.NewDAO: %s", err.Error())
	}

	s, err := services.NewServices(d, cfg)
	if err != nil {
		log.Fatal("services.NewServices: %s", err.Error())
	}

	prs := hub3.NewParser(cfg, d)

	apiServer := api.NewAPI(cfg, s, d)

	sch := scheduler.NewScheduler()

	sch.AddProcessWithInterval(s.UpdateValidatorsMap, time.Minute*10)
	sch.AddProcessWithInterval(s.UpdateProposals, time.Minute*15)
	sch.AddProcessWithInterval(s.UpdateValidators, time.Minute*15)
	sch.EveryDayAt(s.MakeUpdateBalances, 1, 0)
	sch.EveryDayAt(s.MakeStats, 2, 0)

	go s.KeepHistoricalState()

	g := modules.NewGroup(apiServer, sch, prs)
	g.Run()

	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Interrupt, os.Kill)

	<-interrupt
	g.Stop()

	os.Exit(0)
}
