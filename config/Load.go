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
	accountsJSONEnv := os.Getenv("MAIL_ACCOUNT_KEEPER_ACCOUNTS")

	var accountsJSON string

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

	var accounts []AccountConfig
	err := json.Unmarshal([]byte(accountsJSON), &accounts)

	if err != nil {
		log.Fatalf("Error parsing accounts JSON: %v", err)
	}

	fmt.Printf("Starting mail-account-keeper %s...\n", v)

	c.AccountConfigs = accounts

	return c
}
