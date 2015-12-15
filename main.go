// Geo DB based on Mapanet Limited CSV DB
//
// Simple HTTP server
//
// Copyright 2015 Nexway <ceble@nexway.com>

package main

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/spf13/viper"
)

var (
	configPath string
	configName string
	force      bool
	appName    string
	v          *viper.Viper
)

// main set up the commands line and the flags
func main() {
	app := cli.NewApp()
	app.Name = "go-init"
	app.Usage = "Initialize a go project"
	app.Version = "1.1.0"

	// global flags
	//Log Level and Config Path
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "configPath, c",
			EnvVar: "CONFIGPATH",
			Usage:  "config path",
			Value:  "./",
		},
		cli.StringFlag{
			Name:   "configName, n",
			EnvVar: "CONFIGNAME",
			Usage:  "name of the config",
			Value:  "config",
		},
		cli.BoolFlag{
			Name:   "force, f",
			EnvVar: "FORCE",
			Usage:  "force the creatino of the folder",
		},
	}
	app.Action = func(c *cli.Context) {
		if len(c.Args()) == 0 {
			return
		}
		if len(c.Args()) == 1 {
			appName = c.Args()[0]
			directory = "./"

		} else {
			directory = c.Args()[0]
			appName = c.Args()[1]
		}
		Generate(c)
	}
	app.Before = func(ctx *cli.Context) error {
		configPath = ctx.String("configPath")
		configName = ctx.String("ConfigName")
		force = ctx.Bool("force")
		return nil
	}
	app.Run(os.Args)

}

/*************************************
				CONFIG
*************************************/

// LoadConfig load configuration file
func LoadConfig() {
	if v == nil {
		v = viper.New()
		v.SetConfigName(configName)
		v.AddConfigPath(configPath)
		err := v.ReadInConfig() // Find and read the config file
		if err != nil {         // Handle errors reading the config file
			log.Fatal(fmt.Errorf("Fatal error config file: %s \n", err))
		}
	}
}

// Config returns the current viper instance
func Config() *viper.Viper {
	if v == nil {
		LoadConfig()
	}
	return v
}
