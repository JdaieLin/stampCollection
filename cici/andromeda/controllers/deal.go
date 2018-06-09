package controllers

import (
	"cici/andromeda/constant"
	"cici/andromeda/models"
)

type DealController struct {
	Controller
}

func (c *DealController) List() {
	type Request struct {
		// TargetUserID []int64
		BaseRequest
		StampIds   []int64 `json:"stamp_id"`
		Rarity     []int
		Brightness []int
		Crack      []int
		Mark       []int
		Stain      []int
		DealID     []int `json:"deal_id"`
		SerialID   []int `json:"serail_id"`
		Limit      int
		Sort       int
	}
	var req *Request
	var err = c.ParseForm(req)
	if err != nil {
		c.WriteError(constant.ERR_REQ_BODY, string(c.PostBody()))
		return
	}
	deals, err := models.ListDeals(req.UserId)
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
	var req *Request
	var err = c.ParseForm(req)
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
	var req *Request
	var err = c.ParseForm(req)
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
