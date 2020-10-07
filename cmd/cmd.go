package cmd

import (
	"flag"
	"go-boilerplate/database/migration"
	"os"
)

func LoadCMDConfiguration() {
	isExit := false
	migrate := flag.Bool("migrate", false, "a migrate args")
	flag.Parse()

	if *migrate == true {
		migration.Migrate()
		isExit = true
	}

	if isExit {
		os.Exit(0)
	}
}
