package models

import "gopkg.in/mgo.v2/bson"

type Serial struct {
	ID   int64  `bson:"_id" json:"id"`
	Name string `bson:"Name" json:"name"`
}

func GetAllSerials() (serials []Serial, err error) {
	err = DB().All(COL_SERIAL, bson.M{}, &serials)
	return
}
