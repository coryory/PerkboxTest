package config

import "os"

// Config the main configuration.
type Config struct {
	Data  `json:"-"`
	MySQL SQL `json:"sql"`
	Web   Web `json:"web"`
}

// SQL contains all the apropriate information for the SQL connection.
type SQL struct {
	URI string `json:"uri"`
}

// Web configures gin and other web elements
type Web struct {
	ListenAddress string `json:"listen_address"`
}

// DefaultConfig the default configuration to save.
var DefaultConfig = Config{
	Data: Data{},
	MySQL: SQL{
		URI: "username:password@tcp(127.0.0.1:3306)/perkbox?charset=utf8&parseTime=True&loc=Local",
	},
	Web: Web{
		ListenAddress: ":8080",
	},
}

// Save saves the config.
func (c *Config) Save() error {
	saveLoc, envThere := os.LookupEnv("CONFIG_LOC")
	if !envThere {
		saveLoc = SaveLocation
	}

	return c.save(saveLoc, c)
}

// Load loads the config.
func (c *Config) Load() error {

	saveLoc, envThere := os.LookupEnv("CONFIG_LOC")
	if !envThere {
		saveLoc = SaveLocation
	}

	if err := c.load(saveLoc, c); err == DefaultConfigSavedError {
		if err := DefaultConfig.Save(); err != nil {
			return err
		}
		return DefaultConfigSavedError
	} else if err != nil {
		return err
	}

	return nil

}
