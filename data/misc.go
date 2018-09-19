package data

import (
	"errors"

	"github.com/jinzhu/gorm"
)

var ErrorBadID = errors.New("bad id")

// UserCoupon holds the relationship of a user and a coupon via their ID's.
type UserCoupon struct {
	gorm.Model
	UserId   uint
	CouponId uint
}

func (uc *UserCoupon) Create(h *Handler) error {
	if uc.CouponId == 0 || uc.UserId == 0 {
		return ErrorBadID
	}

	return h.Engine.FirstOrCreate(uc, uc).Error
}

func (u *User) GetCoupons(h *Handler) ([]Coupon, error) {
	var (
		coupons     []Coupon
		userCoupons []UserCoupon
		err         error
	)

	if err = h.Engine.Model(&UserCoupon{}).Where("user_id = ?", u.ID).Find(&userCoupons).Error; err != nil {
		return coupons, err
	}

	for _, v := range userCoupons {
		var coupon Coupon
		err := h.Engine.First(&coupon, "id = ?", v.CouponId)
		if err == nil {
			coupons = append(coupons, coupon)
		}
	}

	return coupons, err

}
