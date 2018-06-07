package controllers

import (
	"cici/andromeda/constant"
	"cici/andromeda/models"

	"encoding/json"
)

type UserController struct {
	Controller
}

func (c *UserController) Register() {
	var u = new(models.User)
	err := json.Unmarshal(c.PostBody(), u)
	if err != nil {
		c.WriteError(constant.ERR_REQ_BODY, string(c.PostBody()))
		return
	}
	if u.Name == "" {
		c.WriteError(constant.ERR_REG_NAME_NULL)
		return
	}

	if u.ExistName() {
		c.WriteError(constant.ERR_REG_NAME_EXIST, u.Name)
		return
	}

	if u.Phone != "" && u.ExistPhone() {
		c.WriteError(constant.ERR_REG_PHONE_EXIST, u.Phone)
		return
	}

	if u.LoginID == "" {
		c.WriteError(constant.ERR_REG_LOGIN_ID_NULL)
		return
	}

	if u.ExistLoginID() {
		c.WriteError(constant.ERR_REG_LOGIN_ID_EXIST, u.LoginID)
		return
	}
	if err = u.Insert(); err != nil {
		c.WriteError(constant.ERR_REQ_SERVER, err.Error)
	} else {
		c.WriteSuccessResult(map[string]interface{}{"id": u.ID})
	}
	return
}

func (c *UserController) Logout() {
	c.WriteSuccess("logout")
}

func (c *UserController) Login() {
	type Request struct {
		BaseRequest
		models.User
	}
	var req *Request
	var err = c.ParseForm(req)
	if err != nil {
		c.WriteError(constant.ERR_REQ_BODY, string(c.PostBody()))
		return
	}
	if req.LoginID == "" {
		c.WriteError(constant.ERR_REG_LOGIN_ID_NULL)
		return
	}
	if req.LoginPassword == "" {
		c.WriteError(constant.ERR_REG_LOGIN_PASSWORD_NULL)
		return
	}
	pass, err := req.Login()
	if err != nil {
		c.WriteError(constant.ERR_LOGIN_LOGIN_ID_INVAID)
		return
	}
	if !pass {
		c.WriteError(constant.ERR_LOGIN_PASSWORD_INVAID)
	} else {
		c.WriteSuccessResult(req.User)
	}
	return
}

func (c *UserController) Update() {
	type Request struct {
		BaseRequest
		models.User
	}
	var req *Request
	var err = c.ParseForm(req)
	if err != nil {
		c.WriteError(constant.ERR_REQ_BODY, string(c.PostBody()))
		return
	}
	var u = req.User
	if u.Name == "" {
		c.WriteError(constant.ERR_REG_NAME_NULL)
		return
	}
	if u.LoginID == "" {
		c.WriteError(constant.ERR_REG_LOGIN_ID_NULL)
		return
	}
	u.ID = req.UserId
	if err = u.Update(); err != nil {
		return
	}
	c.WriteError(constant.ERR_UNSUPPORTED, c.Ctx.Input.URI)
}
