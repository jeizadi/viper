package main

import "github.com/spf13/pflag"

const (
	// configuration defaults support local development (i.e. "go run ...")
	// Server
	defaultServerAddress = "0.0.0.0"
	defaultServerPort    = "9090"
	defaultServerVersion = "/v1/"

	// Configuration
	defaultConfigDirectory = "deploy/"
	defaultConfigFile      = ""

	// Logging
	defaultLoggingLevel = "debug"
)

var (
	// define flag overrides
	flagServerAddress = pflag.String("server.address", defaultServerAddress, "adress of gRPC server")
	flagServerPort    = pflag.String("server.port", defaultServerPort, "port of gRPC server")
	flagServerVersion = pflag.String("server.version", defaultServerVersion, "endpoint version of server")

	flagConfigDirectory = pflag.String("config.source", defaultConfigDirectory, "directory of the configuration file")
	flagConfigFile      = pflag.String("config.file", defaultConfigFile, "directory of the configuration file")

	//flagLoggingLevel = pflag.String("logging.level", defaultLoggingLevel, "log level of application")
	flagPLoggingLevel = pflag.StringP("logging.level", "d", defaultLoggingLevel, "log level of application")
)