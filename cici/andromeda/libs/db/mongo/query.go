package mongo

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func (m *MDB) Upsert(collection string, selector interface{}, change interface{}) error {
	return m.Query(func(db *MGO) error {
		_, err := db.C(collection).Upsert(selector, change)
		return err
	})
}

func (m *MDB) UpdateId(collection string, id interface{}, change interface{}) error {
	return m.Query(func(db *MGO) error {
		return db.C(collection).UpdateId(id, change)
	})
}

func (m *MDB) IncrID(collection string, selector string, count int, ids ...int64) error {
	var query = bson.M{"_id": bson.M{"$in": ids}}
	var change = mgo.Change{
		Update:    bson.M{"$inc": bson.M{"id": 1}},
		Upsert:    true,
		ReturnNew: true,
	}
	return m.Query(func(db *MGO) error {
		var result = bson.M{}
		_, err := db.C(collection).Find(query).Apply(change, result)
		return err
	})
}

func (m *MDB) Update(collection string, selector, change interface{}) error {
	return m.Query(func(db *MGO) error {
		return db.C(collection).Update(selector, change)
	})
}
func (m *MDB) UpdateAll(collection string, selector, change interface{}) error {
	return m.Query(func(db *MGO) error {
		_, err := db.C(collection).UpdateAll(selector, change)
		return err
	})
}

func (m *MDB) Insert(collection string, data ...interface{}) error {
	return m.Query(func(db *MGO) error {
		return db.C(collection).Insert(data...)
	})
}

func (m *MDB) All(collection string, query interface{}, result interface{}) error {
	return m.Query(func(db *MGO) error {
		return db.C(collection).Find(query).All(result)
	})
}

// 返回所有复合 query 条件的item， 并且被 projection 限制返回的fields
func (m *MDB) AllSelect(collection string, query interface{}, projection interface{}, result interface{}) error {
	return m.Query(func(db *MGO) error {
		return db.C(collection).Find(query).Select(projection).All(result)
	})
}

func (m *MDB) One(collection string, query interface{}, result interface{}) error {
	return m.Query(func(db *MGO) error {
		return db.C(collection).Find(query).One(result)
	})
}

func (m *MDB) OneSelect(collection string, query interface{}, projection interface{}, result interface{}) error {
	return m.Query(func(db *MGO) error {
		return db.C(collection).Find(query).Select(projection).One(result)
	})
}

//等效于: m.One(collection,bson.M{"_id":id},result)
func (m *MDB) FindId(collection string, id interface{}, result interface{}) error {
	return m.Query(func(db *MGO) error {
		return db.C(collection).Find(bson.M{"_id": id}).One(result)
	})
}

func (m *MDB) RemoveId(collection string, id interface{}) error {
	return m.Query(func(db *MGO) error {
		err := db.C(collection).RemoveId(id)
		return err
	})
}
func (m *MDB) Remove(collection string, selector interface{}) error {
	return m.Query(func(db *MGO) error {
		err := db.C(collection).Remove(selector)
		return err
	})
}
func (m *MDB) RemoveAll(collection string, selector interface{}) error {
	return m.Query(func(db *MGO) error {
		_, err := db.C(collection).RemoveAll(selector)
		return err
	})
}

func (m *MDB) CountId(collection string, id interface{}) (n int) {
	m.Query(func(db *MGO) error {
		var err error
		n, err = db.C(collection).FindId(id).Count()
		return err
	})
	return n
}
func (m *MDB) Count(collection string, query interface{}) (n int) {
	m.Query(func(db *MGO) error {
		var err error
		n, err = db.C(collection).Find(query).Count()
		return err
	})
	return n
}
func (m *MDB) Exist(collection string, query interface{}) bool {
	return m.Count(collection, query) != 0
}
func (m *MDB) ExistId(collection string, id interface{}) bool {
	return m.CountId(collection, id) != 0
}

func (m *MDB) PageSort(collection string, query bson.M, sort []string, offset int, limit int, result interface{}) error {
	return m.Query(func(db *MGO) error {
		return db.C(collection).Find(query).Skip(offset).Limit(limit).Sort(sort...).All(result)
	})
}

func (m *MDB) Page(collection string, query bson.M, offset int, limit int, result interface{}) error {
	return m.Query(func(db *MGO) error {
		return db.C(collection).Find(query).Skip(offset).Limit(limit).All(result)
	})
}

//获取页面数据和“所有”符合条件的记录“总共”的条数
func (m *MDB) PageAndCount(collection string, query bson.M, offset int, limit int, result interface{}) (total int, err error) {
	err = m.Query(func(db *MGO) error {
		total, err = db.C(collection).Find(query).Count()
		if err != nil {
			return err
		}
		return db.C(collection).Find(query).Skip(offset).Limit(limit).All(result)
	})
	return total, err
}

//等同与UpdateId(collection,id,bson.M{"$set":change})
func (m *MDB) SetId(collection string, id interface{}, change interface{}) error {
	return m.UpdateId(collection, id, bson.M{"$set": change})
}
