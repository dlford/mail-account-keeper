package main

import (
	"fmt"
	"mail-account-keeper/config"
	"sync"
	"time"

	"github.com/go-co-op/gocron"
)

var Version string = "v1.0.0"

func main() {
	var c config.Config
	c.Load(Version)

	var wg sync.WaitGroup

	for _, lc := range c.AccountConfigs {
		s := gocron.NewScheduler(time.Local)
		go func(c_lc config.AccountConfig, c_s *gocron.Scheduler) {
			wg.Add(1)
			defer wg.Done()
			s.Cron(c_lc.Schedule).Do(run, &c_lc, c_s)
			s.StartBlocking()
		}(lc, s)
		run(&lc, s)
	}

	wg.Wait()
}

func run(c *config.AccountConfig, s *gocron.Scheduler) {
	fmt.Printf("Sending mail from account \"%s\"...\n", c.Title)
	// TODO: Implement
	time.Sleep(1 * time.Second)
	fmt.Println("Done!")

	_, next := s.NextRun()
	fmt.Printf("Next send for account \"%s\" scheduled at %s\n", c.Title, next)
}
