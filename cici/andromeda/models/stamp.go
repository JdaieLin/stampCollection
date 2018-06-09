package models

import (
	"cici/andromeda/libs/util"
	"math/rand"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Stamp struct {
	ID         int64   `bson:"_id" json:"id"`                // 系统邮票ID
	Name       string  `bson:"Name" json:"name"`             //邮票名称
	Expain     string  `bson:"Expain" json:"expain"`         //邮票介绍
	Rarity     int     `bson:"Rarity" json:"rarity"`         // 稀有度 0.L 1.N 2.R 3.SR 4.SSR
	Score      float64 `bson:"Score" json:"score"`           // 评分
	Brightness int     `bson:"Brightness" json:"brightness"` //辉度
	Crack      int     `bson:"Crack" json:"crack"`           //破损
	Mark       int     `bson:"mark" json:"mark"`             //邮戳
	Stain      int     `bson:"Stain" json:"stain"`           //污点
	SerialID   int64   `bson:"SerialID" json:"serial_id"`    //套票ID
	SerialNum  int64   `bson:"SerialNum" json:"serial_num"`  //套票ID
	OwnerID    int64   `bson:"OwnerID" json:"owner_id"`      //所属用户ID，系统ID为100
	Floor      float64 `bson:"Floor" json:"floor"`           //底价
	Image      string  `bson:"Image" json:"image"`
	DealIds    []int64 `bson:"DealIds" json:"deal_ids"`       //交易记录
	CreateTime int64   `bson:"CreateTime" json:"create_time"` //首次入库时间
	LastModify int64   `bson:"LastModify" json:"last_modify"` // 上次操作时间（收藏/购买/交易）
}

func NewStamp(id int64) *Stamp {
	return &Stamp{ID: id}
}

func (s *Stamp) Fill() error {
	return DB().FindId(COL_STAMP, s.ID, s)
}

func (s *Stamp) updateDealID() (err error) {
	return DB().Update(COL_STAMP, bson.M{"_id": s.ID}, bson.M{"$set": bson.M{"DealIds": s.DealIds}})
}

func (s *Stamp) updateOwner() (err error) {
	return DB().Update(COL_STAMP, bson.M{"_id": s.ID}, bson.M{"$set": bson.M{"OwnerID": s.OwnerID}})
}

func (s *Stamp) Purchase(user_id int64) (err error) {
	if err = s.Fill(); err != nil {
		return
	}
	s.OwnerID = user_id
	if err = s.updateOwner(); err != nil {
		return
	}
	CacheStampApp.Purchase(s.ID, user_id)
	return
}

func PurchaseStamp(user_id, stamp_id int64) (err error) {
	return NewStamp(stamp_id).Purchase(user_id)
}

func CollectStamp(user_id, stamp_id int64) (err error) {
	return NewUser(user_id).CollectStamp(stamp_id)
}

type CacheStamp struct {
	Slice []*Stamp
	Map   map[int64]*Stamp
	Rand  *rand.Rand
}

func NewCacheStamp() *CacheStamp {
	return &CacheStamp{
		Map:  make(map[int64]*Stamp),
		Rand: rand.New(rand.NewSource(time.Now().Unix())),
	}
}

func (c *CacheStamp) Init() (err error) {
	if err = DB().All(COL_STAMP, bson.M{"OwnerID": 100}, &c.Slice); err != nil {
		return
	}
	for _, stamp := range c.Slice {
		c.Map[stamp.ID] = stamp
	}
	return
}

func (c *CacheStamp) Purchase(stamp_id, buyer_user_id int64) {
	if stamp, ok := c.Map[stamp_id]; ok {
		stamp.OwnerID = buyer_user_id
	}
}

func (c *CacheStamp) RandN(n int) (stamps []*Stamp) {
	for i := 0; i < n; i++ {
		if stamp := c.RandOne(); stamp == nil {
			break
		} else {
			stamps = append(stamps, stamp)
		}
	}
	return
}

func (c *CacheStamp) RandOne() (stamp *Stamp) {
	for i := 0; i < 100 && (stamp == nil || stamp.OwnerID != 100); i++ {
		stamp = c.Slice[c.Rand.Intn(len(c.Slice))]
	}
	return
}

func ListStamps(status []int, user_id int64, detail bool) (stamps []*Stamp, err error) {
	var u = NewUser(user_id)
	if err = u.Fill(); err != nil {
		return
	}
	for _, book := range u.Books {
		if util.MatchArrayInt(book.Status, status) {
			var stamp = NewStamp(book.StampID)
			if err = stamp.Fill(); err != nil {
				continue
			}
			stamps = append(stamps, stamp)
		}
	}
	return
}

func GetStamps(ids []int64) (stamps []*Stamp, err error) {
	if err = DB().All(COL_STAMP, bson.M{"_id": bson.M{"$in": ids}}, &stamps); err != nil {
		return
	}
	return
}

var CacheStampApp *CacheStamp

func init() {
	CacheStampApp = NewCacheStamp()
	if err := CacheStampApp.Init(); err != nil {
		panic("初始化邮票缓存失败")
	}
}
