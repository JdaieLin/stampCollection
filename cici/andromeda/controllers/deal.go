package controllers

import (
	"cici/andromeda/constant"
	"cici/andromeda/models"
	"encoding/json"
)

type DealController struct {
	Controller
}

func (c *DealController) List() {
	type Request struct {
		BaseRequest
		CreateUserIds []int64 `json:"create_user_id"`
		StampIds      []int64 `json:"stamp_ids"`
		DealIDs       []int64 `json:"deal_ids"`
		SerialID      []int   `json:"serail_id"`
		Limit         int
		Sort          int
	}
	var req = new(Request)
	var err = json.Unmarshal(c.PostBody(), req)
	if err != nil {
		c.WriteError(constant.ERR_REQ_BODY, string(c.PostBody()))
		return
	}
	deals, err := models.ListDeals(req.CreateUserIds, req.StampIds, req.DealIDs)
	if err != nil {
		c.WriteError(constant.ERR_REQ_SERVER, err.Error())
	} else {
		c.WriteSuccessResult(map[string]interface{}{"deals": deals})
	}
}

func (c *DealController) Create() {
	type Request struct {
		BaseRequest
		StampIDs []int64 `json:"stamp_ids"`
		Price    float64
	}
	var req = new(Request)
	var err = json.Unmarshal(c.PostBody(), req)
	if err != nil {
		c.WriteError(constant.ERR_REQ_BODY, string(c.PostBody()))
		return
	}
	deal, err := models.CreateDeal(req.StampIDs, req.Price, req.UserId)
	if err != nil {
		c.WriteError(constant.ERR_REQ_SERVER, err.Error())
	} else {
		c.WriteSuccessResult(deal)
	}
}

func (c *DealController) Accept() {
	type Request struct {
		BaseRequest
		DealID int64 `json:"deal_id"`
	}
	var req = new(Request)
	var err = json.Unmarshal(c.PostBody(), req)
	if err != nil {
		c.WriteError(constant.ERR_REQ_BODY, string(c.PostBody()))
		return
	}
	err = models.AcceptDeal(req.DealID, req.UserId)
	if err != nil {
		c.WriteError(constant.ERR_REQ_SERVER, err.Error())
	} else {
		c.WriteSuccess("")
	}
}
