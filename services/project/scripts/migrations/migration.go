package main

import (
	"flag"
	"fmt"
	"os"

	migrations "github.com/go-pg/migrations/v8"
	pg "github.com/go-pg/pg/v10"

	"in-backend/services/project/configs"
)

const usageText = `This program runs command on the db. Supported commands are:
  - init - creates version info table in the database
  - up - runs all available migrations.
  - up [target] - runs available migrations up to the target one.
  - down - reverts last migration.
  - reset - reverts all migrations.
  - version - prints current db version.
  - set_version [version] - sets db version without running migrations.
Usage:
  go run *.go <command> [args]
`

func main() {
	flag.Usage = usage
	flag.Parse()

	cfg, err := configs.LoadConfig(configs.FileName)
	if err != nil {
		println(err.Error())
		os.Exit(-1)
	}

	db := pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Database.Address, cfg.Database.Port),
		User:     cfg.Database.Username,
		Password: cfg.Database.Password,
		Database: cfg.Database.Database,
	})

	mc := migrations.NewCollection()

	err = mc.DiscoverSQLMigrations(".")
	if err != nil {
		exitf(err.Error())
	}

	oldVersion, newVersion, err := mc.Run(db, flag.Args()...)
	if err != nil {
		exitf(err.Error())
	}
	if newVersion != oldVersion {
		fmt.Printf("migrated from version %d to %d\n", oldVersion, newVersion)
	} else {
		fmt.Printf("version is %d\n", oldVersion)
	}
}

func usage() {
	fmt.Print(usageText)
	flag.PrintDefaults()
	os.Exit(2)
}

func errorf(s string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, s+"\n", args...)
}

func exitf(s string, args ...interface{}) {
	errorf(s, args...)
	os.Exit(1)
}
