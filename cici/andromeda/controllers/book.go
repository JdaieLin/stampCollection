package controllers

import (
	"cici/andromeda/constant"
	"cici/andromeda/models"
	"encoding/json"
)

type SlotController struct {
	Controller
}

func (c *SlotController) Unlock() {
	type Request struct {
		BaseRequest
		SerialID    int64 `json:"serial_id"`
		SerialOrder int   `json:"serial_order"`
	}
	var req = new(Request)
	var err = json.Unmarshal(c.PostBody(), req)
	if err != nil {
		c.WriteError(constant.ERR_REQ_BODY, string(c.PostBody()))
		return
	}
	var u = models.NewUser(req.UserId)
	if err = u.UnlockSlot(req.SerialID, req.SerialOrder); err != nil {
		c.WriteError(constant.ERR_REQ_SERVER, err.Error())
	}
	c.WriteSuccess("")
	return
}

func (c *SlotController) List() {
	var req = new(BaseRequest)
	var err = json.Unmarshal(c.PostBody(), req)
	if err != nil {
		c.WriteError(constant.ERR_REQ_BODY, string(c.PostBody()))
		return
	}
	var u = models.NewUser(req.UserId)
	if slots, err := u.ListSlot(); err != nil {
		c.WriteError(constant.ERR_REQ_SERVER, err.Error())
	} else {
		c.WriteSuccessResult(map[string]interface{}{"slots": slots})
	}
}
