package models

import ()

type SystemLoggingConfig struct {
	BaseObj
	SRLogger      string `SNAPROUTE: "KEY", ACCESS:"w",  MULTIPLICITY:"1", DESCRIPTION: "Global logging"`
	SystemLogging string `DESCRIPTION: "Global logging", DEFAULT: "on"`
}

type ComponentLoggingConfig struct {
	BaseObj
	Module string `SNAPROUTE: "KEY", ACCESS:"w",  MULTIPLICITY:"*", DESCRIPTION: "Module name to set logging level"`
	Level  string `DESCRIPTION: "Logging level", DEFAULT: "info"`
}
