package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/spf13/viper"
)

var (
	configPath     string
	v              *viper.Viper
	App            *cli.App
	appName        string
	appUsage       string
	appVersion     string
	appDescription string
)

// Main client installation
func Main() {
	App = cli.NewApp()
	App.EnableBashCompletion = true
	App.Name = appName
	App.Usage = appUsage
	App.Version = appVersion

	// global flags
	//Log Level and Config Path
	App.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "config, c",
			Value:  "./",
			EnvVar: "CONFIG",
			Usage:  "absolute config path",
		},
	}
	App.Commands = []cli.Command{{
		Name:   "start",
		Usage:  appDescription,
		Action: server,
		Flags:  App.Flags,
	}}
}

// Run the application
func Run() {
	App.Before = func(ctx *cli.Context) error {
		configPath = ctx.String("config")
		return nil
	}
	err := App.Run(os.Args)
	if err != nil {
		log.Warn("an eror occur while running the program")
	}
}

// Server launch the API
func server(ctx *cli.Context) {
	//configInit()
	initLogLevel()
	log.Fatal(http.ListenAndServe(Config().GetString("addr"), handler)
}

// configInit initialiaze the loading of the config
func configInit() {
	LoadConfig(configPath, "config")
	Config().SetDefault("log-level", "debug")
	Config().SetDefault("addr", "localhost:8081")
}

// getLogLevel Retrieves log level
func getLogLevel() log.Level {
	lvl, err := log.ParseLevel(Config().GetString("log-level"))
	if err != nil {
		log.WithFields(log.Fields{
			"passed":  lvl,
			"default": "fatal",
		}).Warn("Log level is not valid, fallback to default level")
		return log.FatalLevel
	}
	return lvl
}

// LoadConfig load configuration file
func LoadConfig(pathFile, fileName string) {
	if v == nil {
		v = viper.New()
		v.SetConfigName(fileName)
		v.AddConfigPath(pathFile)
		log.Print("Loading config file")
		err := v.ReadInConfig() // Find and read the config file
		if err != nil {         // Handle errors reading the config file
			log.Fatal(fmt.Errorf("Fatal error config file: %s \n", err))
		}
	}
}

// Config returns the current viper instance
func Config() *viper.Viper {
	if v == nil {
		LoadConfig(configPath, "config")
	}
	return v
}

// initLogLevel Static method to init log level
func initLogLevel() {
	log.SetLevel(getLogLevel())
}

// Get version
func Version() string {
	version, err := ioutil.ReadFile("VERSION")
	if err != nil {
		log.Fatal(err)
	}
	var value = string(version)
	return value
}
