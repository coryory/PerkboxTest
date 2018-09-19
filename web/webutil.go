package web

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func (web *Web) getJson(w http.ResponseWriter, r *http.Request, i interface{}) error {
	if r.Header.Get("Content-Type") != "application/json" {
		web.WriteError(w, 400, errors.New("invalid content type"))
		return errors.New("invalid content type")
	}

	defer r.Body.Close()
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		web.WriteError(w, 400, errors.New("data sent was invalid"))
		return err
	}

	if err := json.Unmarshal(bytes, i); err != nil {
		web.WriteError(w, 400, errors.New("data sent was invalid: "+err.Error()))
		return err
	}

	return nil
}
