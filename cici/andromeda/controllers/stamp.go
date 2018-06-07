package controllers

import (
	"cici/andromeda/constant"
	"cici/andromeda/models"
	"encoding/json"
)

type StampController struct {
	Controller
}

func (c *StampController) Get() {
	type Request struct {
		BaseRequest
		StampIds []int64 `json:"stamp_ids"`
	}
	var req = new(Request)
	var err = json.Unmarshal(c.PostBody(), req)
	if err != nil {
		c.WriteError(constant.ERR_REQ_BODY, string(c.PostBody()))
		return
	}
	if stamps, err := models.GetStamps(req.StampIds); err != nil {
		c.WriteError(constant.ERR_REQ_SERVER, err.Error())
	} else {
		c.WriteSuccessResult(stamps)
	}
}

func (c *StampController) List() {
	type Request struct {
		BaseRequest
		Status        []int  `json:"status"`
		StampKeywrods string `json:"stamp_keywords"`
		Detail        bool   `json:"bool"`
		Page          int    `json:"page"`
		Limit         int    `json:"limit"`
	}
	var req = new(Request)
	var err = json.Unmarshal(c.PostBody(), req)
	if err != nil {
		c.WriteError(constant.ERR_REQ_BODY, string(c.PostBody()))
		return
	}
	type Response struct {
		Stamps []*models.Stamp
		Pages  int
	}
	if stamps, err := models.ListStamps(req.Status, req.UserId, req.Detail); err != nil {
		c.WriteError(constant.ERR_REQ_SERVER, err.Error())
	} else {
		c.WriteSuccessResult(Response{Stamps: stamps})
	}
}

func (c *StampController) Collect() {
	type Request struct {
		BaseRequest
		StampID int64 `json:"stamp_id"`
	}
	var req = new(Request)
	var err = json.Unmarshal(c.PostBody(), req)
	if err != nil {
		c.WriteError(constant.ERR_REQ_BODY, string(c.PostBody()))
		return
	}
	if err = models.CollectStamp(req.UserId, req.StampID); err != nil {
		c.WriteError(constant.ERR_REQ_SERVER, err.Error())
	} else {
		c.WriteSuccess("")
	}
}

func (c *StampController) Buy() {
	type Request struct {
		BaseRequest
		StampID int64 `json:"stamp_id"`
	}
	var req = new(Request)
	var err = json.Unmarshal(c.PostBody(), req)
	if err != nil {
		c.WriteError(constant.ERR_REQ_BODY, string(c.PostBody()))
		return
	}
	if err = models.PurchaseStamp(req.UserId, req.StampID); err != nil {
		c.WriteError(constant.ERR_REQ_SERVER, err.Error())
	} else {
		c.WriteSuccess("")
	}
}
