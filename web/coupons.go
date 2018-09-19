package web

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/SilverCory/PerkboxTest/data"
)

func (web *Web) CouponHandler(w http.ResponseWriter, r *http.Request) {
	id :=
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

}

func (web *Web) CouponUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		web.WriteError(w, 405, errors.New("method not allowed, patch only"))
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

	var oldCoupon data.Coupon

	if err := coupon.Create(web.data); err != nil {
		web.WriteError(w, 500, err)
	}

}
