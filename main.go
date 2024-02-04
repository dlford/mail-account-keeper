package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/go-co-op/gocron"

	"mail-account-keeper/config"
	"mail-account-keeper/mail"
)

var Version string = "v1.2.0"

func main() {
	var c config.Config
	c.Load(Version)

	var wg sync.WaitGroup

	for _, lc := range c.AccountConfigs {
		s := gocron.NewScheduler(time.Local)
		go func(c_lc config.AccountConfig, c_s *gocron.Scheduler, c_al *config.AlertConfig) {
			wg.Add(1)
			defer wg.Done()
			s.Cron(c_lc.Schedule).Do(run, &c_lc, c_s, c_al)
			s.StartBlocking()
		}(lc, s, &c.AlertConfig)
		if c.RunOnStart {
			run(&lc, s, &c.AlertConfig)
		}
	}

	wg.Wait()
}

func run(c *config.AccountConfig, s *gocron.Scheduler, a *config.AlertConfig) {
	fmt.Printf("Sending mail from account \"%s\"...\n", c.Title)

	mail.Send(c, a)

	_, next := s.NextRun()
	fmt.Printf("Next send for account \"%s\" scheduled at %s\n", c.Title, next)
}
