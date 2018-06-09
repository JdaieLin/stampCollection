package models

import (
	"cici/andromeda/conf"
	"cici/andromeda/libs/db/mongo"
)

const (
	COL_USER  = "user"
	COL_DEAL  = "deal"
	COL_STAMP = "stamp"
)

var mdb *mongo.MDB

func init() {
	var err error
	mdb, err = mongo.NewMDB(&conf.App().Mongo)
	if err != nil {
		panic("mongo connect faild:" + err.Error())
	}
}

func DB() *mongo.MDB {
	return mdb
}
