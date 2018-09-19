package web

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/SilverCory/PerkboxTest/config"
	"github.com/SilverCory/PerkboxTest/data"
)

var Methods = []string{
	http.MethodGet,
	http.MethodHead,
	http.MethodPost,
	http.MethodPut,
	http.MethodPatch,
	http.MethodDelete,
	http.MethodConnect,
	http.MethodOptions,
	http.MethodTrace,
}

func TestWeb(t *testing.T) {
	dbUri, exists := os.LookupEnv("databaseUri")
	if !exists {
		t.Fatal("No database URI was provided. Please provide a database URI via setting the 'databaseUri' environment variable")
		return
	}

	handler, err := data.NewHandler(&config.SQL{URI: dbUri})
	if err != nil {
		t.Fatal("Unable to connect to the test database!", err)
	}

	web := &Web{
		data: handler,
	}

	t.Run("TestCoupon_Create", func(t *testing.T) {
		t.Run("MethodCheck", func(t *testing.T) {

			for _, v := range Methods {
				t.Run(v, func(t *testing.T) {
					req, err := http.NewRequest(v, "/coupons/create", nil)
					if err != nil {
						t.Fatal(err)
						return
					}

					rec := httptest.NewRecorder()
					web.CouponCreate(rec, req)

					if rec.Code == http.StatusOK {
						t.Fatal("Invalid response from CouponCreate, returned OK http status")
						return
					}
				})
			}

		})
	})
}
