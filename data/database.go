package data

import (
	"github.com/SilverCory/PerkboxTest/config"
	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Handler struct {
	Engine *gorm.DB
}

// NewHandler will create and initialise a new database/storage handler
func NewHandler(conf *config.SQL) (*Handler, error) {
	ret := new(Handler)

	db, err := gorm.Open("mysql", conf.URI)
	if err != nil {
		return ret, err
	}

	ret.Engine = db
	return ret, nil

}

// Migrate and set up keys for the database.
func (h *Handler) Migrate() error {
	// TODO automigrate!
	if err := h.Engine.AutoMigrate().Error; err != nil {
		return err
	}
	return nil
}
