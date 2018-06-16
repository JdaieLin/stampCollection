package models

import "gopkg.in/mgo.v2/bson"

type Atlas struct {
	ID         int64  `bson:"_id" json:"id"`
	Name       string `bson:"Name" json:"name"` //邮票名称
	Image      string `bson:"Image" json:"image"`
	Expain     string `bson:"Expain" json:"expain"`      //邮票介绍
	SerialID   int64  `bson:"SerialID" json:"serial_id"` //套票ID
	SetName    string `bson:"SetName" json:"set_name"`
	Set        string `bson:"Set" json:"set"`
	Order      int64  `bson:"Order" json:"order"` //套票ID
	Date       int64  `bson:"Date" json:"date"`   //发行日期
	Year       int    `bson:"Year" json:"year"`
	RealVolume int64  `bson:"RealVolume" json:"real_volume"`
	Volume     int64  `bson:"Volume" json:"volume"`
	Amount     int64  `bson:"Amount" json:"amount"` //发行数量
	Remain     int64  `bson:"Remain" json:"remain"` //剩余数量
	Floor      int    `bson:"Floor" json:"floor"`
}

func NewAtlas(id int64) Atlas {
	return Atlas{ID: id}
}

func (a *Atlas) All() {
	DB().IncrID(COL_ATLAS, "Amount", -1, a.ID)
}

func (a *Atlas) Recycle() {
	DB().IncrID(COL_ATLAS, "Amount", 1, a.ID)
}

func GetAllAtlas() (atlases []Atlas, err error) {
	err = DB().All(COL_ATLAS, bson.M{}, &atlases)
	return
}
