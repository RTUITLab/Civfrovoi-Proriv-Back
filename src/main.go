package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/RTUITLab/Civfrovoi-Proriv-Back/pkg/config"
	"github.com/RTUITLab/Civfrovoi-Proriv-Back/app"
)

func main() {
	cfg := config.GetConfig()

	a := app.New(cfg)

	log.Infof("App start on %s", a.Port)
	if err := a.Start(); err != nil {
		log.Panic("Failed to start app")
	}
}