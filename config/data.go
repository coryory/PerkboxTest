package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"path/filepath"
)

// Data abstract containing methods for saving and loading.

// Data an abstract struct used for it's functions to save and load config files.
type Data struct{}

// ConfigSaveLocation the location to save the config to.
var SaveLocation = "config.json"

// DefaultConfigSavedError an error returned if the default config is saved.
var DefaultConfigSavedError = errors.New("the default config has been saved, please edit it")

func (d *Data) save(saveLoc string, inter interface{}) error {
	// Make all the directories
	if err := os.MkdirAll(filepath.Dir(saveLoc), os.ModeDir|0775); err != nil {
		return err
	}

	data, err := json.MarshalIndent(inter, "", "    ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(saveLoc, data, 0660)
}

func (d *Data) load(saveLoc string, inter interface{}) error {

	if _, err := os.Stat(saveLoc); os.IsNotExist(err) {
		return DefaultConfigSavedError
	} else if err != nil {
		return err
	}

	data, err := ioutil.ReadFile(saveLoc)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, inter); err != nil {
		return err
	}

	return nil

}
