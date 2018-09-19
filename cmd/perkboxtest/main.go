package main

import (
	"github.com/SilverCory/PerkboxTest/config"
	"github.com/SilverCory/PerkboxTest/data"
	"github.com/SilverCory/PerkboxTest/web"
)

func main() {
	conf := new(config.Config)
	e(conf.Load())

	handler, err := data.NewHandler(&conf.SQL)
	e(err)

	e(handler.Migrate())

	w := web.NewWeb(conf.Web, handler)
	e(w.Start())
}

func e(err error) {
	if err != nil {
		panic(err)
	}
}
