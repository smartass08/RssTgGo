package main

import (
	"RssTgGo/modules"
	"RssTgGo/utils"
	"sync"
	"time"
)

var db utils.DB
var wg sync.WaitGroup

func CheckAll() {
	defer wg.Done()
	for _, v := range utils.GetAllRssNyaaLinks() {
		modules.NyaaRss(modules.Nyaa(v))
	}
}

func main() {
	db.Access(utils.GetDbUrl())
	db.GetAllhash()
	for {
		wg.Add(1)
		go CheckAll()
		time.Sleep(time.Minute)
		wg.Wait()
	}

}
