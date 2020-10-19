package modules

import (
	"context"
	"github.com/mmcdole/gofeed"
	"log"
	"net/http"
	"strings"
	"time"
)

func Nyaa(link string) []map[string]string {
	var maps []map[string]string
	fp := gofeed.NewParser()
	fp.Client = &http.Client{Timeout: time.Second * 20}
	feed, err := fp.ParseURLWithContext(link, context.Background())
	if err != nil {
		log.Println(err)
		if strings.Contains(err.Error(), "connection reset by peer") {
			time.Sleep(time.Second * 2)
			Nyaa(link)
		}
	}
	for _, v := range feed.Items {
		m := make(map[string]string)
		m["title"] = v.Title
		m["view_link"] = v.GUID
		m["torrent_link"] = v.Link
		m["hash"] = v.Extensions["nyaa"]["infoHash"][0].Value
		m["size"] = v.Extensions["nyaa"]["size"][0].Value
		maps = append(maps, m)
	}
	return maps
}
