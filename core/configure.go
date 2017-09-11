package core

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sync"
)

var (
	//Gconf global configure object
	gconf map[string]interface{}
	cl    sync.RWMutex
)

func init() {
	gconf = make(map[string]interface{})
	if err := loadConfigFile(); err != nil {
		log.Fatalln("global config init error,", err)
		return
	}
}

func loadConfigFile() (err error) {
	defer Recover("load config file")
	buf, err := ioutil.ReadFile("./config.json")
	if err != nil {
		return
	}
	tmpConfig := make(map[string]interface{})
	if err = json.Unmarshal(buf, &tmpConfig); err != nil {
		return
	}
	cl.Lock()
	gconf = tmpConfig
	cl.Unlock()
	return
}

func GetConfig(key string) interface{} {
	cl.RLock()
	defer cl.RUnlock()
	if v, has := gconf[key]; has {
		return v
	}
	return nil
}

func GetConfigString(key string) string {
	cl.RLock()
	defer cl.RUnlock()
	if v, has := gconf[key]; has {
		switch v.(type) {
		case string:
			return v.(string)
		default:
			return ""
		}
	}
	return ""
}
