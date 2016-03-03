package models

import ()

type SystemLoggingConfig struct {
	BaseObj
	Logging string `SNAPROUTE: "KEY", ACCESS:"w",  MULTIPLICITY:"1", DESCRIPTION: "Global logging"`
}

type ComponentLoggingConfig struct {
	BaseObj
	Name  string `SNAPROUTE: "KEY", ACCESS:"w",  MULTIPLICITY:"*", DESCRIPTION: "Module name to set logging level"`
	Level string `DESCRIPTION: "Logging level", DEFAULT: "info"`
}
