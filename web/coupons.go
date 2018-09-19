package web

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/SilverCory/PerkboxTest/data"
)

// CouponHandler is the parent handler and allows for the URL to be /api/coupon/:id
func (web *Web) CouponHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/coupon/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		web.WriteError(w, 400, errors.New("expected coupon ID in url"))
		return
	}

	if r.Method == http.MethodPatch {
		web.CouponUpdate(w, r, uint(id))
	}

}

func (web *Web) CouponCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		web.WriteError(w, 405, errors.New("method not allowed, post only"))
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		web.WriteError(w, 400, errors.New("invalid content type"))
		return
	}

	defer r.Body.Close()
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		web.WriteError(w, 400, errors.New("data sent was invalid"))
		return
	}

	var coupon data.Coupon
	if err := json.Unmarshal(bytes, &coupon); err != nil {
		web.WriteError(w, 400, errors.New("data sent was invalid: "+err.Error()))
		return
	}

	if err := coupon.Create(web.data); err != nil {
		web.WriteError(w, 500, err)
	}

	w.WriteHeader(200)
}

func (web *Web) CouponUpdate(w http.ResponseWriter, r *http.Request, id uint) {

	if r.Header.Get("Content-Type") != "application/json" {
		web.WriteError(w, 400, errors.New("invalid content type"))
		return
	}

	defer r.Body.Close()
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		web.WriteError(w, 400, errors.New("data sent was invalid"))
		return
	}

	var updates map[string]interface{}
	if err := json.Unmarshal(bytes, updates); err != nil {
		web.WriteError(w, 400, errors.New("data sent was invalid: "+err.Error()))
		return
	}

	coupon := new(data.Coupon)
	coupon.ID = uint(id)

	err = web.data.Engine.Model(coupon).Update(updates).Error
	if err != nil {
		web.WriteError(w, 500, errors.New("unable to update: "+err.Error()))
		return
	}

	w.WriteHeader(200)
}
