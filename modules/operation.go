package modules

import (
	"RssTgGo/utils"
)

var db utils.DB
var tg TG

func NyaaRss(links []map[string]string) {
	tg.Initialise()
	db.Access(utils.GetDbUrl())
	for _, v := range links {
		if !utils.CheckValidMap(v) {
			continue
		}
		if utils.CheckValid(v["hash"]) {
			return
		}
		db.Insert(v["hash"])
		tg.Send(v, false)
	}
}
