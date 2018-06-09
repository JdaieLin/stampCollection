package controllers

import (
	"cici/andromeda/constant"
	"cici/andromeda/models"
	"encoding/json"
)

type ChestController struct {
	Controller
}

func (c *ChestController) List() {
	var req = new(BaseRequest)
	var err = json.Unmarshal(c.PostBody(), req)
	if err != nil {
		c.WriteError(constant.ERR_REQ_BODY, string(c.PostBody()))
		return
	}
	var u = models.User{
		ID: req.UserId,
	}
	if err = u.ChestList(); err != nil {
		c.WriteError(constant.ERR_REQ_SERVER, err.Error())
		return
	}
	var res = map[string]interface{}{
		"chests": u.Chests.List,
		"ingots": u.Ingots,
	}
	c.WriteSuccessResult(res)
}

func (c *ChestController) Sync() {
	type Request struct {
		BaseRequest
		Chests []models.Chest
	}
	var req = new(Request)
	var err = json.Unmarshal(c.PostBody(), req)
	if err != nil {
		c.WriteError(constant.ERR_REQ_BODY, string(c.PostBody()))
		return
	}
	var u = models.User{
		ID: req.UserId,
	}
	if err = u.ChestSync(req.Chests); err != nil {
		c.WriteError(constant.ERR_REQ_SERVER, err.Error())
		return
	}
	c.WriteSuccess("")
}
