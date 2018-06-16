package main

import (
	"cici/andromeda/models"
	"log"

	"gopkg.in/mgo.v2/bson"
)

func update(uid int64) {
	var u = models.NewUser(uid)
	if err := u.Fill(); err != nil {
		log.Fatalf("fill user faild :%d", uid)
	}
	for _, book := range u.Books {
		if book.Status > 101 {
			stamp := models.NewStamp(book.StampID)
			stamp.OwnerID = uid
			models.DB().Update(models.COL_STAMP, bson.M{"_id": stamp.ID}, bson.M{"$set": bson.M{"OwnerID": uid}})
			log.Printf("%d stamp updated", book.StampID)
		}
	}
}

func main() {
	update(3)
	update(4)
}
