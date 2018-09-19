package web

import (
	"fmt"
	"net/http"

	"github.com/SilverCory/PerkboxTest/config"
	"github.com/SilverCory/PerkboxTest/data"
)

type Web struct {
	data *data.Handler
	conf config.Web
}

func NewWeb(conf config.Web, data *data.Handler) *Web {
	ret := &Web{
		data: data,
		conf: conf,
	}

	// TODO basic userauth middleware

	http.HandleFunc("/api/coupon/create", ret.CouponCreate)
	http.HandleFunc("/api/coupon/", ret.CouponHandler)
	return ret
}

func (web *Web) Start() error {
	return http.ListenAndServe(web.conf.ListenAddress, nil)
}

func (web *Web) WriteError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	w.Write([]byte(fmt.Sprintf(`{"error": %q}`, err)))
}
