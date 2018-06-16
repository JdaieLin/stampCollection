package controllers

import (
	"cici/andromeda/constant"
	"cici/andromeda/models"
	"encoding/json"
)

type SweepController struct {
	Controller
}

func (c *SweepController) List() {
	type Request struct {
		BaseRequest
		Max int
	}
	var req = new(Request)
	var err = json.Unmarshal(c.PostBody(), req)
	if err != nil {
		c.WriteError(constant.ERR_REQ_BODY, string(c.PostBody()), err.Error())
		return
	}
	if stamps := models.CacheStampApp.RandN(req.Max); stamps == nil {
		c.WriteError(constant.ERR_REQ_SERVER, string(c.PostBody()))
	} else {
		c.WriteSuccessResult(map[string]interface{}{"stamps": stamps})
	}
	return
}

func (c *SweepController) Sync() {
	type Action struct {
		StampID int `json:"stamp_id"`
		Action  int `json:"action"`
	}
	type Request struct {
		BaseRequest
		Coins  int64
		Stamps []Action
	}
	var req = new(Request)
	var err = json.Unmarshal(c.PostBody(), req)
	if err != nil {
		c.WriteError(constant.ERR_REQ_BODY, string(c.PostBody()))
		return
	}
	if err = models.NewUser(req.UserId).SweepSync(req.Coins); err != nil {
		c.WriteError(constant.ERR_REQ_SERVER, string(c.PostBody()))
	} else {
		c.WriteSuccess("")
	}
}
