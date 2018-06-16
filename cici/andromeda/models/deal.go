package models

import (
	"cici/andromeda/constant"
	"errors"
	"fmt"
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
	d.Status = constant.STATUS_CLOSED
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

func ListDeals(user_ids, stamp_ids, deal_ids []int64) (deals []*Deal, err error) {
	var query = bson.M{}
	if len(user_ids) != 0 {
		query["CreateUserID"] = bson.M{
			"$in": user_ids,
		}
	}
	if len(stamp_ids) != 0 {
		query["StampIDs"] = bson.M{
			"$in": stamp_ids,
		}
	}
	if len(deal_ids) != 0 {
		query["_id"] = bson.M{
			"$in": deal_ids,
		}
	}
	if err = DB().All(COL_DEAL, query, &deals); err != nil {
		return
	}
	for _, deal := range deals {
		deal.Stamps, err = GetStamps(deal.StampIDs)
	}
	return
}

func AcceptDeal(deal_id, user_id int64) (err error) {
	var deal = NewDeal(deal_id)
	if err = deal.Accept(user_id); err != nil {
		return
	}
	if err = NewUser(deal.CreateUser).SellStamp(deal.StampIDs...); err != nil {
		return
	}
	if err = NewUser(user_id).BuyStamp(deal.StampIDs...); err != nil {
		return
	}
	return
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
		if len(stamp.DealIds) > 0 {
			latest_deal := NewDeal(stamp.DealIds[len(stamp.DealIds)-1])
			if err = latest_deal.Fill(); err != nil {
				err = fmt.Errorf("deal info unvailable ：%d", latest_deal.ID)
				return
			}
			if latest_deal.Status == constant.STATUS_OPEN {
				err = fmt.Errorf("can't open new deal for stamp:%d", latest_deal.ID)
				return
			}
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
		Status:     constant.STATUS_OPEN,
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
	if err = NewUser(user_id).OnsaleStamps(stamp_ids); err != nil {
		return
	}
	return
}
