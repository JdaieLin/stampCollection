package models

import (
	"errors"
	"math/rand"
	"time"

	"cici/andromeda/constant"

	"gopkg.in/mgo.v2/bson"
)

type UserChain struct {
	Address string `bson:"Address" json:"address"` //区块链用户地址
	Key     string `bson:"Key" json:"key"`         //区块链用户密钥
}

type StampBook struct {
	StampID    int64 `bson:"StampID" json:"stamp_id"`
	Status     int   `bson:"Status" json:"status"`
	CreateTime int64 `bson:"CreateTime" json:"create_time"`
}

type Chest struct {
	Hour   int   `bson:"Hour" json:"hour"`
	Ingot  int64 `bson:"Ingot" json:"ingot"`
	Opened bool  `bson:"Opened" json:"opened"`
}

type Chests struct {
	List [24]Chest `bson:"List" json:"list"`
	Date int       `bson:"Date" json:"date"`
}

type Slot struct {
	SerialID    int64 `bson:"SerialID" json:"serial_id"`
	SerialOrder int   `bson:"SerialOrder" json:"serial_order"`
	CreateTime  int64 `bson:"CreateTime" json:"create_time"`
}

type User struct {
	ID            int64        `bson:"_id" json:"id"`                       //系统用户ID
	Chain         UserChain    `bson:"Chain" json:"chain"`                  //区块链用户信息
	Name          string       `bson:"Name" json:"name"`                    //显示用户昵称
	LoginID       string       `bson:"LoginID" json:"login_id"`             //登陆用户名
	LoginPassword string       `bson:"LoginPassword" json:"login_password"` //登陆密码
	Mail          string       `bson:"Mail" json:"mail"`                    //用户邮箱
	Phone         string       `bson:"Phone" json:"phone"`                  //用户手机号
	IDCard        string       `bson:"IDCard" json:"id_card"`               //用户身份证
	Coins         int64        `bson:"Coins" json:"coins"`                  //用户铜板
	Ingots        int64        `bson:"Ingots" json:"ingots"`                //用户元宝
	Books         []*StampBook `bson:"Books" json:"books"`                  //邮票列表
	Chests        Chests       `bson:"Chests" json:"chests"`
	Slots         []Slot       `bson:"Slots" json:"slots"`
}

func NewUser(id int64) *User {
	return &User{
		ID: id,
	}
}

func (u *User) ExistName() bool {
	return DB().Exist(COL_USER, bson.M{"Name": u.Name})
}

func (u *User) ExistLoginID() bool {
	return DB().Exist(COL_USER, bson.M{"LoginID": u.LoginID})
}

func (u *User) ExistPhone() bool {
	return DB().Exist(COL_USER, bson.M{"Phone": u.Phone})
}

func (u *User) Insert() (err error) {
	if u.ID == 0 {
		if u.ID, err = DB().AutoIncId(COL_USER); err != nil {
			return
		}
	}
	return DB().Insert(COL_USER, u)
}

func (u *User) Update() (err error) {
	return DB().UpdateId(COL_USER, u.ID, bson.M{"$set": u})
}

func (u *User) Login() (pass bool, err error) {
	password := u.LoginPassword
	if err = DB().One(COL_USER, bson.M{"LoginID": u.LoginID}, u); err != nil {
		return
	}
	pass = u.LoginPassword == password
	return
}

func (u *User) Fill() (err error) {
	if err = DB().FindId(COL_USER, u.ID, u); err != nil {
		err = errors.New("db find faild:" + err.Error())
		return
	}
	return
}

func (u *User) updateChest() (err error) {
	return DB().Update(COL_USER, bson.M{"_id": u.ID}, bson.M{"$set": bson.M{"Chests": u.Chests, "Ingots": u.Ingots}})
}

func (u *User) ChestList() (err error) {
	if err = DB().FindId(COL_USER, u.ID, u); err != nil {
		return
	}
	year, month, day := time.Now().Date()
	if u.Chests.Date != year*10000+int(month)*100+day {
		u.Chests.Date = year*10000 + int(month)*100 + day
		for i := 0; i < 24; i++ {
			u.Chests.List[i].Hour = i
			u.Chests.List[i].Ingot = rand.Int63n(10)
		}
		if err = u.updateChest(); err != nil {
			return
		}
	}
	return
}

func (u *User) SweepSync(coins int64) (err error) {
	if err = u.Fill(); err != nil {
		return
	}
	return DB().Update(COL_USER, bson.M{"_id": u.ID}, bson.M{"$set": bson.M{"Coins": u.Coins + coins}})
}

func (u *User) ChestSync(chests []Chest) (err error) {
	if err = u.ChestList(); err != nil {
		return
	}
	for _, chest := range chests {
		if chest.Hour > 0 && chest.Hour < 24 && chest.Opened {
			if u.Chests.List[chest.Hour].Opened == false {
				u.Ingots += u.Chests.List[chest.Hour].Ingot
			}
			u.Chests.List[chest.Hour].Opened = true
		}
	}
	if err = u.updateChest(); err != nil {
		return
	}
	return
}

func (u *User) OnsaleStamps(stamp_ids []int64) (err error) {
	if err = u.Fill(); err != nil {
		return
	}
	for _, stamp_id := range stamp_ids {
		for _, book := range u.Books {
			if book.StampID == stamp_id {
				book.Status = constant.STATUS_ONSALE
				break
			}
		}
	}
	if err = DB().Update(COL_USER, bson.M{"_id": u.ID}, bson.M{"$set": bson.M{"Books": u.Books}}); err != nil {
		return
	}
	return
}

func (u *User) SellStamp(stamp_ids ...int64) (err error) {
	if err = u.Fill(); err != nil {
		return
	}
	for _, stamp_id := range stamp_ids {
		for i, book := range u.Books {
			if book.StampID == stamp_id {
				u.Books = append(u.Books[:i], u.Books[i+1:]...)
				break
			}
		}
	}
	if err = DB().Update(COL_USER, bson.M{"_id": u.ID}, bson.M{"$set": bson.M{"Books": u.Books}}); err != nil {
		return
	}
	return
}

func (u *User) BuyStamp(stamp_ids ...int64) (err error) {
	if err = u.Fill(); err != nil {
		return
	}
	for _, stamp_id := range stamp_ids {
		var found bool
		for _, book := range u.Books {
			if book.StampID == stamp_id {
				book.Status = constant.STATUS_DEAL_OWNER
				break
			}
		}
		if !found {
			u.Books = append(u.Books, &StampBook{StampID: stamp_id, Status: constant.STATUS_DEAL_OWNER, CreateTime: time.Now().Unix()})
		}
	}
	if err = DB().Update(COL_USER, bson.M{"_id": u.ID}, bson.M{"$set": bson.M{"Books": u.Books}}); err != nil {
		return
	}
	return
}

func (u *User) ListSlot() (slots []Slot, err error) {
	if err = u.Fill(); err != nil {
		return
	}
	slots = u.Slots
	return
}

func (u *User) UnlockSlot(serial_id int64, serial_order int) (err error) {
	if err = u.Fill(); err != nil {
		return
	}
	var found bool
	for _, slot := range u.Slots {
		if slot.SerialID == serial_id && slot.SerialOrder == serial_order {
			found = true
			break
		}
	}
	if !found {
		if u.Coins > 1000 {
			u.Slots = append(u.Slots, Slot{SerialID: serial_id, SerialOrder: serial_order, CreateTime: time.Now().Unix()})
			u.Coins -= 1000
			if err = DB().Update(COL_USER, bson.M{"_id": u.ID}, bson.M{"$set": bson.M{"Slots": u.Slots, "Coins": u.Coins}}); err != nil {
				return
			}
		} else {
			err = errors.New("铜钱余额不足")
		}
	}
	return
}

func (u *User) CollectStamp(stamp_id int64) (err error) {
	if err = u.Fill(); err != nil {
		return
	}
	var found bool
	for _, book := range u.Books {
		if book.StampID == stamp_id {
			found = true
			break
		}
	}
	if !found {
		if u.Ingots > 200 {
			u.Books = append(u.Books, &StampBook{StampID: stamp_id, Status: constant.STATUS_COLLECT, CreateTime: time.Now().Unix()})
			u.Ingots -= 200
			if err = DB().Update(COL_USER, bson.M{"_id": u.ID}, bson.M{"$set": bson.M{"Books": u.Books, "Ingots": u.Ingots}}); err != nil {
				return
			}
		} else {
			err = errors.New("元宝余额不足")
		}
	}
	return
}

func (u *User) Remove() (err error) {
	return DB().RemoveId(COL_USER, u.ID)
}
