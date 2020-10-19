package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

const ConfigJsonPath string = "config.json"

type ConfigJson struct {
	BOT_TOKEN    string   `json:"bot_token"`
	DB_URL       string   `json:"db_url"`
	DB_Name      string   `json:"db_name"`
	DB_Col       string   `json:"db_collection"`
	CH_Id        int      `json:"channel_Id"`
	RSS_Nyaa_URL []string `json:"rss_nyaa_links"`
}

var Config *ConfigJson = InitConfig()

func InitConfig() *ConfigJson {
	file, err := ioutil.ReadFile(ConfigJsonPath)
	if err != nil {
		log.Fatal("Config File Bad, exiting!")
	}

	var Config ConfigJson
	err = json.Unmarshal([]byte(file), &Config)
	if err != nil {
		log.Fatal(err)
	}
	return &Config
}

func GetBotToken() string {
	return Config.BOT_TOKEN
}

func GetDbUrl() string {
	return Config.DB_URL
}

func GetDbCollection() string {
	return Config.DB_Col
}

func GetDbName() string {
	return Config.DB_Name
}

func GetChannelId() int {
	return Config.CH_Id
}

func GetAllRssNyaaLinks() []string {
	return Config.RSS_Nyaa_URL
}

func CheckValidMap(data map[string]string) bool {
	if len(data["title"]) != 0 && len(data["size"]) != 0 && len(data["view_link"]) != 0 && len(data["torrent_link"]) != 0 && len(data["hash"]) != 0 {
		return true
	} else {
		return false
	}
}
