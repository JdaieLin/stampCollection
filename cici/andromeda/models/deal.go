package models

import (
	"errors"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Deal struct {
	ID          int64    `bson:"_id" json:"id"`                      //系统交易ID
	CreateUser  int64    `bson:"CreateUserID" json:"create_user_id"` //创建用户ID
	CreateTime  int64    `bson:"CreateTime" json:"create_time"`      //创建时间
	TargetUsers []int64  `bson:"TargetUserID" json:"target_user_id"` //指定买家列表
	StampIDs    []int64  `bson:"StampIDs" json:"stamp_ids"`          //交易邮票id列表
	Stamps      []*Stamp `bson:"-" json:"stamps"`                    //交易邮票详情
	Price       float64  `bson:"Price" json:"price"`                 //交易价格
	AcceptUser  int64    `bson:"AcceptUser" json:"accept_user"`
	AcceptTime  int64    `bson:"AcceptTime" json:"accpet_time"` //创建时间
	Status      int64    `bson:"Status" json:"status"`          //交易状态,0.等待交易 1.交易结束
}

func NewDeal(id int64) *Deal {
	return &Deal{
		ID: id,
	}
}

func (d *Deal) Fill() (err error) {
	return DB().FindId(COL_DEAL, d.ID, d)
}

func (d *Deal) Accept(user_id int64) (err error) {
	if err = d.Fill(); err != nil {
		return
	}
	d.Status = 1
	d.AcceptUser = user_id
	d.AcceptTime = time.Now().Unix()
	if err = DB().UpdateId(COL_DEAL, d.ID, bson.M{"$set": bson.M{"Status": d.Status, "AcceptUser": d.AcceptUser, "AcceptTime": d.AcceptTime}}); err != nil {
		return
	}
	if err = DB().UpdateAll(COL_STAMP, bson.M{"_id": bson.M{"$in": d.StampIDs}}, bson.M{"$set": bson.M{"OwnerID": user_id}}); err != nil {
		return
	}
	return
}

func ListDeals(user_id int64) (deals []*Deal, err error) {
	if err = DB().All(COL_DEAL, bson.M{"CreateUserID": user_id}, deals); err != nil {
		return
	}
	for _, deal := range deals {
		deal.Stamps, err = GetStamps(deal.StampIDs)
	}
	return
}

func AcceptDeal(deal_id, user_id int64) (err error) {
	return NewDeal(deal_id).Accept(user_id)
}

func CreateDeal(stamp_ids []int64, price float64, user_id int64) (deal *Deal, err error) {
	stamps, err := GetStamps(stamp_ids)
	if err != nil {
		return
	}
	var pass = true
	var floor float64
	for _, stamp := range stamps {
		if stamp.OwnerID != user_id {
			pass = false
			break
		}
		floor += stamp.Floor
	}
	if floor > price {
		pass = false
	}
	if !pass {
		err = errors.New("no permit")
		return
	}
	deal = &Deal{
		CreateUser: user_id,
		CreateTime: time.Now().Unix(),
		StampIDs:   stamp_ids,
		Price:      price,
		Status:     0,
	}
	if deal.ID, err = DB().AutoIncId(COL_DEAL); err != nil {
		return
	}
	if err = DB().Insert(COL_DEAL, deal); err != nil {
		return
	}
	for _, stamp := range stamps {
		stamp.DealIds = append(stamp.DealIds, deal.ID)
		stamp.updateDealID()
	}
	return
}
