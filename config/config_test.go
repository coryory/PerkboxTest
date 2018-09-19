package config

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestConfig(t *testing.T) {

	configFile, err := ioutil.TempFile("", "config_default")
	if err != nil {
		t.Fatal("Unable to create a temp config file.", err)
		return
	}

	// Remove the file once we're done.
	defer func() {
		os.Remove(configFile.Name())
	}()

	SaveLocation = configFile.Name()

	t.Run("ConfigSaveDefault", func(t *testing.T) {
		if err := DefaultConfig.Save(); err != nil {
			t.Error("Error saving default config.", err)
			return
		}
	})

	t.Run("ConfigLoadDefault", func(t *testing.T) {
		conf := new(Config)
		conf.Data = Data{}
		if err := conf.Load(); err != nil {
			t.Error("Error loading default config.", err)
			return
		}
		if !reflect.DeepEqual(*conf, DefaultConfig) {
			t.Errorf("The default config defers from the saved config.\n%#v\n%#v\n", conf, DefaultConfig)
			return
		}
	})

	// TODO test custom configs.

}
