package main

import (
	"cici/andromeda/models"
	"io/ioutil"
	"log"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

func update(set, body string) {
	var err error
	set = strings.TrimSpace(set)
	if err = models.DB().UpdateAll(models.COL_ATLAS, bson.M{"Set": set}, bson.M{"$set": bson.M{"Expain": body}}); err != nil {
		log.Printf("update faild for : %s %s", set, err)
	}
	if err = models.DB().UpdateAll(models.COL_STAMP, bson.M{"atlas.Set": set}, bson.M{"$set": bson.M{"atlas.Expain": body}}); err != nil {
		log.Printf("update faild for : %s %s", set, err)
	}
}

func main() {
	src := "1992_notes_utf8.txt"
	data, err := ioutil.ReadFile(src)
	if err != nil {
		log.Fatal("read faild : ", err)
	}
	var set, body string
	for _, line := range strings.Split(string(data), "\n") {
		if strings.HasPrefix(line, "志号名称") {
			if set != "" {
				update(set, body)
			}
			set = strings.TrimPrefix(line, "志号名称：")
			body = line + "\n"
		} else {
			body += line + "\n"
		}
	}
	if set != "" {
		update(set, body)
	}
}
