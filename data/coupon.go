package data

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Coupon containing all relevant data for the best deals, creator and expiry.
type Coupon struct {
	gorm.Model
	Name    string    `json:"name"`
	Brand   string    `json:"brand"`
	Value   float64   `json:"value"`
	Expiry  time.Time `json:"expiry"`
	Code    string    `json:"code" sql:"index" gorm:"not null;unique_index"`
	Creator User      `json:"-"`
}

func (c *Coupon) Create(h *Handler) error {
	err := h.Engine.Create(c).Error
	if err != nil {
		return err
	}

	if c.Creator.CreatedAt.IsZero() || c.ID == 0 {
		return nil
	}

	// TODO user<->coupon foreign key.
	return nil
}

func (c *Coupon) Load(h *Handler) error {
	return h.Engine.First(c, c).Error
}
