package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

func (c *Config) Load(v string) *Config {
	help := "Usage: mail-account-keeper [--version] | --accounts \"[...]\""

	showVersion := flag.Bool("version", false, "Show version")

	accountsJSONFlag := flag.String("accounts", "", "JSON string of accounts to send mail from")
	alertsJSONFlag := flag.String("alerts", "", "JSON string of accounts to send alerts to")
	accountsJSONEnv := os.Getenv("MAIL_ACCOUNT_KEEPER_ACCOUNTS")
	alertsJSONEnv := os.Getenv("MAIL_ACCOUNT_KEEPER_ALERTS")

	var accountsJSON string
	var alertsJSON string

	flag.Parse()

	if *showVersion {
		fmt.Printf(v)
		os.Exit(0)
	}

	if *accountsJSONFlag != "" {
		accountsJSON = *accountsJSONFlag
	} else if accountsJSONEnv != "" {
		accountsJSON = accountsJSONEnv
	} else {
		fmt.Println(help)
		os.Exit(1)
	}

	if *alertsJSONFlag != "" {
		alertsJSON = *alertsJSONFlag
	} else if alertsJSONEnv != "" {
		alertsJSON = alertsJSONEnv
	}

	var accounts []AccountConfig
	err := json.Unmarshal([]byte(accountsJSON), &accounts)

	if err != nil {
		log.Fatalf("Error parsing accounts JSON: %v", err)
	}

	if alertsJSON != "" {
		var alerts AlertConfig
		err := json.Unmarshal([]byte(alertsJSON), &alerts)
		if err != nil {
			log.Fatalf("Error parsing alerts JSON: %v", err)
		}
		fmt.Printf("Alerts account registered: %s\n", alerts.MailTo)
		c.AlertConfig = alerts
	}

	fmt.Printf("Starting mail-account-keeper %s...\n", v)

	c.AccountConfigs = accounts

	return c
}
