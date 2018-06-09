package mongo

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"cici/andromeda/libs/db/mongo/conf"
)

var AutoIncIdCollection string = "amb_config"

type MGO struct {
	*mgo.Database
}

func (db *MGO) AutoIncId(name string) (id int64, err error) {
	result := make(map[string]interface{})
	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"id": 1}},
		Upsert:    true,
		ReturnNew: true,
	}
	_, err = db.C(AutoIncIdCollection).Find(bson.M{"name": name}).Apply(change, result)
	if err != nil {
		return
	}

	rid, ok := result["id"]
	if ok {
		switch rid.(type) {
		case int:
			id = int64(rid.(int))
		case int64:
			id = rid.(int64)
		case float64:
			id = int64(rid.(float64))
		default:
			err = fmt.Errorf("bad id type:%+v", reflect.TypeOf(rid))
		}
	}
	return
}

type MDB struct {
	*mgo.Session
	Conf *conf.MultiDbConfig
}

// Mongodb
func (m *MDB) Open() *MGO {
	return &MGO{m.DB(m.Conf.Database)}
}

// 应该使用_id字段写name(索引)
// 考虑到可能有在使用了的不能直接改
func (m *MDB) Query(query func(*MGO) error) error {
	return query(m.Open())
}

func (m *MDB) AutoIncId(name string) (id int64, err error) {
	err = m.Query(func(db *MGO) error {
		id, err = db.AutoIncId(name)
		return err
	})
	return
}

func NewMDB(conf *conf.MultiDbConfig) (mdb *MDB, err error) {
	var servers []string
	for _, server := range conf.Servers {
		servers = append(servers, server.GetAddr())
	}
	url := strings.Join(servers, ",")
	session, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	// 默认1 minute
	session.SetSocketTimeout(time.Minute * 5)
	mdb = &MDB{
		Session: session,
		Conf:    conf,
	}
	go mdb.autoReconnect()
	return
}

func (m *MDB) autoReconnect() {
	var err error
	for {
		err = m.Ping()
		if err != nil {
			m.Refresh()
		}
		time.Sleep(time.Second * 10)
	}
}
