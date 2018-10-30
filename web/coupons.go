package web

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"

	"github.com/SilverCory/PerkboxTest/data"
)

func (web *Web) CouponList(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		web.WriteError(w, 405, errors.New("method not allowed, get only"))
		return
	}

	var coupons []data.Coupon
	if err := web.data.Engine.Model(&data.Coupon{}).Find(coupons).Error; err != nil {
		web.WriteError(w, 500, errors.New("error fetching coupons: "+err.Error()))
		return
	}

	bytes, err := json.Marshal(coupons)
	if err != nil {
		web.WriteError(w, 500, errors.New("unable to marshal coupons: "+err.Error()))
		return
	}

	w.Write(bytes)
	w.WriteHeader(200)
}

func (web *Web) CouponCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		web.WriteError(w, 405, errors.New("method not allowed, post only"))
		return
	}

	var coupon data.Coupon
	web.getJson(w, r, &coupon)

	if err := coupon.Create(web.data); err != nil {
		web.WriteError(w, 500, err)
	}

	w.WriteHeader(200)
}

// CouponHandler is the parent handler and allows for the URL to be /api/coupon/:id
func (web *Web) CouponHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/coupon/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		web.WriteError(w, 400, errors.New("expected coupon ID in url"))
		return
	}

	switch r.Method {
	case http.MethodPatch:
		web.couponUpdate(w, r, uint(id))
		break

	case http.MethodGet:
		web.couponGet(w, r, uint(id))
		break

	default:
		web.WriteError(w, 405, errors.New("bad method"))
	}

}

// couponUpdate allows for updates to be made  the coupon via a patch request.
func (web *Web) couponUpdate(w http.ResponseWriter, r *http.Request, id uint) {
	var updates map[string]interface{}
	web.getJson(w, r, &updates)

	coupon := new(data.Coupon)
	coupon.ID = uint(id)

	err := web.data.Engine.Model(coupon).Update(updates).Error
	if err != nil {
		web.WriteError(w, 500, errors.New("unable to update: "+err.Error()))
		return
	}

	w.WriteHeader(200)
}

func (web *Web) couponGet(w http.ResponseWriter, r *http.Request, id uint) {
	coupon := new(data.Coupon)
	coupon.ID = uint(id)

	err := web.data.Engine.First(coupon).Error
	if err == gorm.ErrRecordNotFound {
		web.WriteError(w, 404, errors.New("coupon not found"))
		return
	} else if err != nil {
		web.WriteError(w, 500, errors.New("unable to update: "+err.Error()))
		return
	}

	bytes, err := json.Marshal(coupon)
	if err != nil {
		web.WriteError(w, 500, errors.New("unable to marshal coupon: "+err.Error()))
		return
	}

	w.Write(bytes)
	w.WriteHeader(200)
}
