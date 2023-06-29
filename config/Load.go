package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func (c *Config) Load(v string) *Config {
	help := "Usage: mail-account-keeper [-v|--version] | [-c|--config path/to/config.yml]"

	var path string
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "-v", "--version":
			fmt.Println(v)
			os.Exit(0)
			break
		case "-c", "--config":
			if len(os.Args) < 3 {
				fmt.Println(help)
				os.Exit(1)
			}
			path = os.Args[2]
			break
		default:
			fmt.Println(help)
			os.Exit(1)
		}
	}
	if path == "" {
		path = "/etc/blocklister.yml"
	}

	fmt.Printf("Starting Blocklister %s...\n", v)

	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Cannot read config file: %s\n#%v", path, err)
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Cannot unmarshal config file: %v", err)
	}

	fmt.Printf("Loaded configuration file: %s\n", path)
	return c
}
